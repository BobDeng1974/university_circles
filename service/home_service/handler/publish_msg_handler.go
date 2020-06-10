package handler

import (
	"context"
	"fmt"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"regexp"
	"time"
	"university_circles/service/home_service/logic"
	"university_circles/service/home_service/models"
	pb "university_circles/service/home_service/pb/home"
	"university_circles/service/home_service/utils/common"
	"university_circles/service/home_service/utils/logger"
)

type HomeHandler struct {
}

func (h *HomeHandler) SavePublishMsg(ctx context.Context, req *pb.PublishMsg, resp *pb.Response) (err error) {
	fmt.Println("vContent", req.Content)
	hl := &logic.HomeLogic{}

	msgId := common.GenRandomRecommendMsgId(32)
	createTime := time.Now().Format("2006-01-02 15:04:05")
	orderId, _ := time.Parse("2006-01-02 15:04:05", createTime)

	if err = hl.SaveHomeToES(msgId, orderId, createTime, req); err != nil {
		resp.Success = -1
		resp.Msg = "save home publish msg to es failed"
		return nil
	}

	// 存储动态到db
	go hl.AsyncSaveHomeToDB(msgId, req)

	resp.Success = 0
	resp.Msg = "success"
	return nil
}

func (h *HomeHandler) DeletePublishMsg(ctx context.Context, req *pb.DeleteMsgRequest, resp *pb.Response) (err error) {
	hl := &logic.HomeLogic{}
	err = hl.Delete(req.Uid, req.Id)
	if err != nil {
		resp.Success = -1
		resp.Msg = "delete home publish msg failed"
		logger.Logger.Warn("delete home publish msg failed", zap.Error(err))
		return
	}
	return nil
}

func (h *HomeHandler) GetHomeMsgList(ctx context.Context, req *pb.AllMsgListRequest, resp *pb.HomeMsgListResponse) (err error) {
	hl := &logic.HomeLogic{}
	resp.PublishMsgResponseList, err = hl.GetHomeMsgList(req.MsgType, req.OrderId, req.LoadMore, req.Uid, req.UniversityId)
	fmt.Println("GetHomeMsgList", req, err)
	if err != nil {
		logger.Logger.Warn("get home publish msg failed", zap.Error(err))
		return
	}
	fmt.Println(resp.PublishMsgResponseList)
	return nil
}

func (h *HomeHandler) GetUserMsgList(ctx context.Context, req *pb.UserMsgListRequest, resp *pb.OtherMsgListResponse) (err error) {
	hl := &logic.HomeLogic{}
	resp.PublishMsgResponseList, err = hl.GetUserMsgList(req.Uid, req.OrderId, req.UniversityId, req.MsgType)
	fmt.Println(" GetUserMsgList", resp.PublishMsgResponseList, req, err)
	if err != nil {
		logger.Logger.Warn("get user publish msg failed", zap.Error(err))
		return
	}

	fmt.Println(resp.PublishMsgResponseList)
	return nil
}

func (h *HomeHandler) GetMsgDetail(ctx context.Context, req *pb.OneMsgRequest, resp *pb.OtherMsgListResponse) (err error) {
	hl := &logic.HomeLogic{}
	resp.PublishMsgResponseList, err = hl.GetMsgDetail(req.Uid, req.Id, req.UniversityId)
	if err != nil {
		logger.Logger.Warn("get home publish msg detail failed", zap.Error(err))
		return
	}

	fmt.Println("GetMsgDetail11111", resp, err)
	return
}

func (h *HomeHandler) GetDetail(ctx context.Context, req *pb.OneMsgRequest, resp *pb.PublishMsgResponse) (err error) {
	hl := &logic.HomeLogic{}
	resp, err = hl.GetDetail(req.Uid, req.Id, req.UniversityId)
	if err != nil {
		logger.Logger.Warn("get home publish msg detail failed", zap.Error(err))
		return
	}

	fmt.Println("GetMsgDetail11111", resp, err)
	return
}

func (h *HomeHandler) SaveMsgComment(ctx context.Context, req *pb.PublishMsgComment, resp *pb.Response) (err error) {
	hl := &logic.HomeLogic{}

	CommId := uuid.NewV4().String()
	createTime := time.Now().Format("2006-01-02 15:04:05")
	orderId, _ := time.Parse("2006-01-02 15:04:05", createTime)

	// 匹配文本中的url
	urlInTextRegexp := regexp.MustCompile(`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)
	urls := urlInTextRegexp.FindAllString(req.Content, 30)

	var urlsInText []*pb.UrlsInText
	for _, url := range urls {
		urlInText := &pb.UrlsInText{
			OriginalUrl: url,
			Url:         url,
		}
		urlsInText = append(urlsInText, urlInText)
	}

	userInfo, err := hl.GetUserInfoFromEs(req.User.Id)

	req.Status = "NORMAL"
	req.Type = "COMMENT"
	req.Id = CommId
	req.CreatedAt = createTime
	req.OrderId = orderId.Unix()
	req.UrlsInText = urlsInText
	req.User = userInfo
	if err = hl.SaveMsgCommentToEs(req); err != nil {
		logger.Logger.Warn("save home publish msg comment to es failed", zap.Error(err))
		return
	}

	go func() {
		replyToCommentId := ""
		if req.ReplyToComment != nil {
			replyToCommentId = req.ReplyToComment.Id
		}
		comment := &home.HomeComment{
			CommID:           CommId,
			UID:              req.User.Id,
			Type:             req.Type,
			TargetId:         req.TargetId,
			Level:            int8(req.Level),
			Content:          req.Content,
			ReplyToCommentId: replyToCommentId,
			Status:           req.Status,
		}
		if err = hl.SaveMsgCommentToDB(comment); err != nil {
			logger.Logger.Warn("save home publish msg comment to db failed", zap.Error(err))
			return
		}

		if len(req.Pictures) > 0 {
			for _, picture := range req.Pictures {
				commentPic := &home.HomeCommentPicture{
					OwnerID:      CommId,
					ThumbnailUrl: picture.ThumbnailUrl,
					MiddlePicUrl: picture.MiddlePicUrl,
					PicUrl:       picture.PicUrl,
					Format:       picture.Format,
					Width:        int(picture.Width),
					Height:       int(picture.Height),
				}
				if err = hl.SaveMsgCommentPicture(commentPic); err != nil {
					logger.Logger.Warn("save home publish msg comment picture to db failed", zap.Error(err))
				}
			}
		}

		if len(req.UrlsInText) > 0 {
			for _, url := range req.UrlsInText {
				commentUrlInText := &home.HomeCommentUrlsInText{
					OwnerID:     CommId,
					Title:       url.Title,
					OriginalUrl: url.OriginalUrl,
					URL:         url.Url,
				}
				if err = hl.SaveMsgCommentUrlInText(commentUrlInText); err != nil {
					logger.Logger.Warn("save home publish msg comment url in text to db failed", zap.Error(err))
				}
			}
		}
		return
	}()

	go hl.AddCommentCountToCache(req.TargetId)

	resp.Success = 0
	resp.Msg = "success"
	return nil
}

func (h *HomeHandler) GetMsgCommentList(ctx context.Context, req *pb.MsgCommentListRequest, resp *pb.MsgCommentListResponse) (err error) {
	hl := &logic.HomeLogic{}
	resp.PublishMsgComments, err = hl.GetCommentList(req.MsgId, req.OrderId)
	if err != nil {
		logger.Logger.Warn("get home publish msg failed", zap.Error(err))
		return
	}
	fmt.Println(resp.PublishMsgComments)
	return nil
}

func (h *HomeHandler) DeletePublishMsgComment(ctx context.Context, req *pb.DeleteMsgCommentRequest, resp *pb.Response) (err error) {
	hl := &logic.HomeLogic{}
	if err = hl.DeleteComment(req.Uid, req.Id); err != nil {
		logger.Logger.Warn("get home publish msg failed", zap.Error(err))
		return
	}
	go hl.SubCommentCountToCache(req.Id)

	return nil
}

func (h *HomeHandler) SaveUserOperateMsgCount(ctx context.Context, req *pb.UserOperateCountRequest, resp *pb.Response) (err error) {
	hl := &logic.HomeLogic{}

	// 1.分享   2.评论   3.点赞  4.取消点赞
	if req.Type == 1 {
		err = hl.AddShareCountToCache(req.MsgId)
	} else if req.Type == 2 {
		err = hl.AddCommentCountToCache(req.MsgId)
	} else if req.Type == 3 {
		err = hl.AddLikeCountToCache(req.MsgId, req.Uid)
		err = hl.AddUserLikeToCache(req.MsgId, req.Uid)
	} else if req.Type == 4 {
		err = hl.SubLikeCountToCache(req.MsgId)
		err = hl.SubUserLikeToCache(req.MsgId, req.Uid)
	}

	if err != nil {
		resp.Success = -1
		resp.Msg = "user operate failed"
		logger.Logger.Warn("delete home publish msg failed", zap.Error(err))
		return
	}

	resp.Success = 0
	resp.Msg = "user operate success"
	return nil
}

func (h *HomeHandler) GetUserOperateMsgCount(ctx context.Context, req *pb.UserOperateCountRequest, resp *pb.UserOperateCountResponse) (err error) {
	hl := &logic.HomeLogic{}

	// 1.分享   2.评论   3.点赞
	if req.Type == 1 {
		resp.Count, err = hl.GetShareCountFromCache(req.MsgId)
	} else if req.Type == 2 {
		resp.Count, err = hl.GetCommentCountFromCache(req.MsgId)
	} else if req.Type == 3 {
		resp.Count, err = hl.GetLikeCountFromCache(req.MsgId)
	}

	if err != nil {
		logger.Logger.Warn("delete home publish msg failed", zap.Error(err))
		return
	}

	return nil
}

func (h *HomeHandler) GetUserOperateMsgRecodeList(ctx context.Context, req *pb.UserOperateRecodeListRequest, resp *pb.OtherMsgListResponse) (err error) {
	return nil
}
