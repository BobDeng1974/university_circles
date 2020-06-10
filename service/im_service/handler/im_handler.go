package handler

import (
	"context"
	"github.com/volatiletech/null"
	"go.uber.org/zap"
	"time"
	"university_circles/service/im_service/logic"
	im "university_circles/service/im_service/models"
	"university_circles/service/im_service/utils/errcode"
	"university_circles/service/im_service/utils/logger"
	"unsafe"

	//"go.uber.org/zap"
	//"university_circles/service/im_service/logic"
	pb "university_circles/service/im_service/pb/im"
	//"university_circles/service/im_service/utils/logger"
)

type ImHandler struct {
}

func (i *ImHandler) SendPeerMsg(ctx context.Context, req *pb.SendPeerMsgReq, resp *pb.Response) (err error) {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser == nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser.Mid != req.Uid {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if req.ChatType == "system" {
		isFriend, err := il.GetFriendById(req.Uid, req.ReceiverUid, req.Appid)
		if err != nil {
			resp.Msg = errcode.ErrAddFriendFailed.Msg
			resp.Success = errcode.ErrAddFriendFailed.Code
			return nil
		}

		if isFriend != nil {
			logger.Logger.Warn("is friend , do not need to send msg", zap.Any("req", req))
			resp.Msg = errcode.ErrIsFriendFailed.Msg
			resp.Success = errcode.ErrIsFriendFailed.Code
			return nil
		}
	} else {
		isFriend, err := il.GetFriendById(req.Uid, req.ReceiverUid, req.Appid)
		if err != nil {
			resp.Msg = errcode.ErrAddFriendFailed.Msg
			resp.Success = errcode.ErrAddFriendFailed.Code
			return nil
		}

		if isFriend == nil {
			logger.Logger.Warn("not friend , do not send msg", zap.Any("req", req))
			resp.Msg = "send msg failed"
			resp.Success = -1
			return nil
		}
	}


	if err = il.SendPeerMsg(req.Appid, req.Uid, req.ReceiverUid, req.ChatType, req.Method, req.MsgType, req.Content, req.Token, req.Status, req.Username); err != nil {
		resp.Msg = "send msg failed"
		resp.Success = -1
		return
	}

	return
}

func (i *ImHandler) AddFriends(ctx context.Context, req *pb.AddFriendReq, resp *pb.Response) (err error) {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser == nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser.Mid != req.Uid {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	isFriend, err := il.GetFriendById(req.Uid, req.FriendUid, req.Appid)
	if err != nil {
		resp.Msg = errcode.ErrAddFriendFailed.Msg
		resp.Success = errcode.ErrAddFriendFailed.Code
		return
	}

	if isFriend != nil {
		resp.Msg = errcode.ErrIsFriendFailed.Msg
		resp.Success = errcode.ErrIsFriendFailed.Code
		return
	}

	// 双方添加好友
	now := time.Now().Unix()
	friend := &im.Friend{
		Appid:     req.Appid,
		UID:       req.Uid,
		FriendUID: req.FriendUid,
		Timestamp: *(*int)(unsafe.Pointer(&now)),
	}

	if err = il.AddFriend(friend); err != nil {
		resp.Msg = errcode.ErrAddFriendFailed.Msg
		resp.Success = errcode.ErrAddFriendFailed.Code
		return
	}

	userFriend := &im.Friend{
		Appid:     req.Appid,
		UID:       req.FriendUid,
		FriendUID: req.Uid,
		Timestamp: *(*int)(unsafe.Pointer(&now)),
	}

	if err = il.AddFriend(userFriend); err != nil {
		resp.Msg = errcode.ErrAddFriendFailed.Msg
		resp.Success = errcode.ErrAddFriendFailed.Code
		return
	}

	resp.Msg = "success"
	resp.Success = 0

	return nil
}

func (i *ImHandler) DelFriends(ctx context.Context, req *pb.DelFriendReq, resp *pb.Response) error {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser == nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser.Mid != req.Uid {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	// 删除双方好友
	friend, err := il.GetFriendById(req.Uid, req.FriendUid, req.Appid)
	if err != nil {
		resp.Msg = errcode.ErrDelFriendFailed.Msg
		resp.Success = errcode.ErrDelFriendFailed.Code
		return err
	}

	if friend == nil {
		resp.Msg = errcode.ErrDelFriendFailed.Msg
		resp.Success = errcode.ErrDelFriendFailed.Code
		return err
	}

	if err = il.DelFriend(friend); err != nil {
		resp.Msg = errcode.ErrDelFriendFailed.Msg
		resp.Success = errcode.ErrDelFriendFailed.Code
		return err
	}

	userFriend, err := il.GetFriendById(req.FriendUid, req.Uid, req.Appid)
	if err != nil {
		resp.Msg = errcode.ErrDelFriendFailed.Msg
		resp.Success = errcode.ErrDelFriendFailed.Code
		return err
	}

	if userFriend == nil {
		resp.Msg = errcode.ErrDelFriendFailed.Msg
		resp.Success = errcode.ErrDelFriendFailed.Code
		return err
	}

	if err = il.DelFriend(userFriend); err != nil {
		resp.Msg = errcode.ErrDelFriendFailed.Msg
		resp.Success = errcode.ErrDelFriendFailed.Code
		return err
	}

	resp.Msg = "success"
	resp.Success = 0
	return nil
}

func (i *ImHandler) UpdateFriendRemark(ctx context.Context, req *pb.UpdateFriendRemarkReq, resp *pb.Response) (err error) {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser == nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser.Mid != req.Uid {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	friend, err := il.GetFriendById(req.Uid, req.FriendUid, req.Appid)
	if err != nil {
		resp.Msg = errcode.ErrUpdateFriendRemarkFailed.Msg
		resp.Success = errcode.ErrUpdateFriendRemarkFailed.Code
		return err
	}

	if friend == nil {
		resp.Msg = errcode.ErrUpdateFriendRemarkFailed.Msg
		resp.Success = errcode.ErrUpdateFriendRemarkFailed.Code
		return err
	}

	friend.Remark = null.String{String: req.Remark, Valid: true}

	err = il.UpdateFriendRemark(friend)
	if err != nil {
		resp.Msg = errcode.ErrUpdateFriendRemarkFailed.Msg
		resp.Success = errcode.ErrUpdateFriendRemarkFailed.Code
		return
	}
	return nil
}

func (i *ImHandler) AddFriendBlackList(ctx context.Context, req *pb.AddFriendBlackListReq, resp *pb.Response) (err error) {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser == nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser.Mid != req.Uid {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	isBlackList, err := il.GetBlackListById(req.Uid, req.FriendUid, req.Appid)
	if err != nil {
		resp.Msg = errcode.ErrAddFriendBlackListFailed.Msg
		resp.Success = errcode.ErrAddFriendBlackListFailed.Code
		return
	}

	if isBlackList != nil {
		resp.Msg = errcode.ErrIsBlackListFailed.Msg
		resp.Success = errcode.ErrIsBlackListFailed.Code
		return
	}

	now := time.Now().Unix()
	friend := &im.Blacklist{
		Appid:     req.Appid,
		UID:       req.Uid,
		FriendUID: req.FriendUid,
		Timestamp: *(*int)(unsafe.Pointer(&now)),
	}

	if err = il.AddFriendBlackList(friend); err != nil {
		resp.Msg = errcode.ErrAddFriendBlackListFailed.Msg
		resp.Success = errcode.ErrAddFriendBlackListFailed.Code
		return
	}

	resp.Msg = "success"
	resp.Success = 0

	return nil
}

func (i *ImHandler) DelFriendBlackList(ctx context.Context, req *pb.DelFriendBlackListReq, resp *pb.Response) error {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser == nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser.Mid != req.Uid {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	blacklist, err := il.GetBlackListById(req.Uid, req.FriendUid, req.Appid)
	if err != nil {
		resp.Msg = errcode.ErrDelFriendBlackListFailed.Msg
		resp.Success = errcode.ErrDelFriendBlackListFailed.Code
		return err
	}

	if blacklist == nil {
		resp.Msg = errcode.ErrDelFriendBlackListFailed.Msg
		resp.Success = errcode.ErrDelFriendBlackListFailed.Code
		return err
	}

	if err = il.DelFriendBlackList(blacklist); err != nil {
		resp.Msg = errcode.ErrDelFriendBlackListFailed.Msg
		resp.Success = errcode.ErrDelFriendBlackListFailed.Code
		return err
	}

	resp.Msg = "success"
	resp.Success = 0

	return nil
}

func (i *ImHandler) GetAllFriends(ctx context.Context, req *pb.FriendListReq, resp *pb.FriendListResp) (err error) {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		return errcode.ErrAuthFailed
	}

	if esUser == nil {
		return errcode.ErrAuthFailed
	}

	if esUser.Mid != req.Uid {
		return errcode.ErrAuthFailed
	}

	friends, err := il.GetAllFriend(req.Uid, req.Appid)
	if err != nil {
		return err
	}

	for _, f := range friends {
		var user *pb.User
		student, err := il.FindOneStudentByMid(f.FriendUID)
		if err != nil {
			return err
		}

		if student != nil {
			user = &pb.User{
				Uid:      f.FriendUID,
				Nickname: student.ScreenName,
				Mute:     false,
				ImageUrl: student.Avatar + "?x-oss-process=image/resize,m_fixed,h_800,w_800",
				Remark:   f.Remark.String,
			}
		} else {
			teacher, err := il.FindOneTeacherByMid(f.FriendUID)
			if err != nil {
				return err
			}

			if teacher != nil {
				user = &pb.User{
					Uid:      f.FriendUID,
					Nickname: teacher.ScreenName,
					Mute:     false,
					ImageUrl: teacher.Avatar + "?x-oss-process=image/resize,m_fixed,h_800,w_800",
					Remark:   f.Remark.String,
				}
			}
		}

		if user != nil {
			resp.FriendUids = append(resp.FriendUids, user)
		}
	}

	return nil
}

func (i *ImHandler) GetBlackList(ctx context.Context, req *pb.BlackListReq, resp *pb.BlackListResp) error {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		return errcode.ErrAuthFailed
	}

	if esUser == nil {
		return errcode.ErrAuthFailed
	}

	if esUser.Mid != req.Uid {
		return errcode.ErrAuthFailed
	}

	blacklist, err := il.GetBlackList(req.Uid, req.Appid)
	if err != nil {
		return err
	}

	for _, f := range blacklist {
		var user *pb.User
		student, err := il.FindOneStudentByMid(f.FriendUID)
		if err != nil {
			return err
		}

		if student != nil {
			user = &pb.User{
				Uid:      f.FriendUID,
				Nickname: student.ScreenName,
				Mute:     false,
				ImageUrl: student.Avatar + "?x-oss-process=image/resize,m_fixed,h_800,w_800",
			}
		} else {
			teacher, err := il.FindOneTeacherByMid(f.FriendUID)
			if err != nil {
				return err
			}

			if teacher != nil {
				user = &pb.User{
					Uid:      f.FriendUID,
					Nickname: teacher.ScreenName,
					Mute:     false,
					ImageUrl: teacher.Avatar + "?x-oss-process=image/resize,m_fixed,h_800,w_800",
				}
			} else {
				return nil
			}
		}

		resp.BlacklistUids = append(resp.BlacklistUids, user)
	}
	return nil
}

func (i *ImHandler) CreateGroup(ctx context.Context, req *pb.Group, resp *pb.CreateGroupResp) (err error) {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser == nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser.Mid != req.Master {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	group := &im.Group{
		Appid:  req.Appid,
		Name:  req.Name,
		Super:  *(*int8)(unsafe.Pointer(&req.Super)),
		Master: req.Master,
		Notice: null.String{String: req.Notice},
		ImageURL: req.ImageUrl,
	}

	if err = il.CreateGroup(group); err != nil {
		resp.Msg = errcode.ErrCreteGroupFailed.Msg
		resp.Success = errcode.ErrCreteGroupFailed.Code
		return
	}

	resp.Msg = "success"
	resp.Success = 0
	resp.GroupId = group.ID

	return nil
}

func (i *ImHandler) DelGroup(ctx context.Context, req *pb.DelGroupReq, resp *pb.Response) error {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser == nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser.Mid != req.Uid {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	group, err := il.GetGroupById(req.GroupId, req.Appid)
	if err != nil {
		resp.Msg = errcode.ErrDelGroupFailed.Msg
		resp.Success = errcode.ErrDelGroupFailed.Code
		return err
	}

	if group == nil {
		resp.Msg = errcode.ErrDelGroupFailed.Msg
		resp.Success = errcode.ErrDelGroupFailed.Code
		return err
	}

	if group.Master != req.Uid {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return err
	}

	if err = il.DelGroup(group); err != nil {
		resp.Msg = errcode.ErrDelGroupFailed.Msg
		resp.Success = errcode.ErrDelGroupFailed.Code
		return err
	}

	resp.Msg = "success"
	resp.Success = 0

	return nil
}

func (i *ImHandler) GetGroupMemberList(ctx context.Context, req *pb.GroupMemberListReq, resp *pb.GroupMemberListResp) error {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		return errcode.ErrAuthFailed
	}

	if esUser == nil {
		return errcode.ErrAuthFailed
	}

	if esUser.Mid != req.Uid {
		return errcode.ErrAuthFailed
	}

	// 判断请求用户是否为群组成员
	isGroupMembers, err := il.GetGroupMember(req.Uid, req.GroupId)
	if err != nil {
		return err
	}

	if isGroupMembers == nil {
		return errcode.ErrAuthFailed
	}

	groupMembers, err := il.GetGroupMemberList(req.GroupId, req.Appid)
	if err != nil {
		return err
	}

	for _, g := range groupMembers {
		var user *pb.User
		friend, err := il.GetFriendById(req.Uid, g.UID, req.Appid)
		if err != nil {
			return err
		}

		remark := ""
		if friend != nil {
			remark = friend.Remark.String
		}


		student, err := il.FindOneStudentByMid(g.UID)
		if err != nil {
			return err
		}

		if student != nil {
			user = &pb.User{
				Uid:      g.UID,
				Nickname: student.ScreenName,
				Mute:     g.Mute,
				ImageUrl: student.Avatar + "?x-oss-process=image/resize,m_fixed,h_800,w_800",
				Remark:   remark,
			}
		} else {
			teacher, err := il.FindOneTeacherByMid(g.UID)
			if err != nil {
				return err
			}

			if teacher != nil {
				user = &pb.User{
					Uid:      g.UID,
					Nickname: teacher.ScreenName,
					Mute:     g.Mute,
					ImageUrl: teacher.Avatar + "?x-oss-process=image/resize,m_fixed,h_800,w_800",
					Remark:   remark,
				}
			} else {
				return nil
			}
		}

		resp.Num, err = il.GetGroupMemberNum(req.GroupId, req.Appid)
		if err != nil {
			return err
		}

		resp.Users = append(resp.Users, user)
	}
	return nil
}

func (i *ImHandler) GetGroupList(ctx context.Context, req *pb.GroupListReq, resp *pb.GroupListResp) error {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		return errcode.ErrAuthFailed
	}

	if esUser == nil {
		return errcode.ErrAuthFailed
	}

	if esUser.Mid != req.Uid {
		return errcode.ErrAuthFailed
	}

	groups, err := il.GetGroupListByUid(req.Uid, req.Appid)
	if err != nil {
		return err
	}

	for _, g := range groups {
		num, _ := il.GetGroupMemberNum(g.ID, req.Appid)

		group := &pb.Group{
			Id:     g.ID,
			Appid:  req.Appid,
			Master: g.Master,
			Super:  int64(g.Super),
			Name:   g.Name,
			Notice: g.Notice.String,
			ImageUrl: g.ImageURL,
			Num: num,
		}
		resp.Groups = append(resp.Groups, group)
	}
	return nil
}

func (i *ImHandler) JoinGroup(ctx context.Context, req *pb.JoinGroupReq, resp *pb.Response) (err error) {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser == nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser.Mid != req.Uid {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	// 判断请求用户是否为群组成员
	isGroupMembers, err := il.GetGroupMember(req.Uid, req.GroupId)
	if err != nil {
		return err
	}

	// 已是群组成员
	if isGroupMembers != nil {
		resp.Msg = "success"
		resp.Success = 0
		return nil
	}

	now := time.Now().Unix()
	groupMember := &im.GroupMember{
		GroupID:   req.GroupId,
		UID:       req.Uid,
		Timestamp: int(now),
		Nickname:  req.Nickname,
		Mute:      false,
	}

	if err = il.JoinGroup(groupMember); err != nil {
		resp.Msg = errcode.ErrJoinGroupFailed.Msg
		resp.Success = errcode.ErrJoinGroupFailed.Code
		return
	}

	// 发送用户加入群组的消息
	go il.SendGroupMsg(req.Appid, req.Uid, req.GroupId,
		req.ChatType, req.Method, req.MsgType,
		req.Content,
		req.Token, req.Status, req.Nickname)

	resp.Msg = "success"
	resp.Success = 0

	return nil
}

func (i *ImHandler) DelGroupMember(ctx context.Context, req *pb.DelGroupMemberReq, resp *pb.Response) (err error) {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser == nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser.Mid != req.Uid {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	group, err := il.GetGroupById(req.GroupId, req.Appid)
	if err != nil {
		return err
	}

	if group == nil {
		resp.Msg = errcode.ErrGroupNotExist.Msg
		resp.Success = errcode.ErrGroupNotExist.Code
		return nil
	}

	if group.Master != req.Uid {
		return errcode.ErrAuthFailed
	}

	// 判断请求用户是否为群组成员
	isGroupMembers, err := il.GetGroupMember(req.Uid, req.GroupId)
	if err != nil {
		return err
	}

	// 不是群组成员
	if isGroupMembers == nil {
		resp.Msg = "success"
		resp.Success = 0
		return nil
	}

	// 判断请求用户是否为群组成员
	if err := il.DelGroupMember(isGroupMembers); err != nil {
		resp.Msg = errcode.ErrDelGroupMemberFailed.Msg
		resp.Success = errcode.ErrDelGroupMemberFailed.Code
		return err
	}

	resp.Msg = "success"
	resp.Success = 0
	return nil
}

func (i *ImHandler) AddGroupFile(ctx context.Context, req *pb.AddGroupFileReq, resp *pb.Response) error {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser == nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser.Mid != req.Uid {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	return nil
}

func (i *ImHandler) GetGroupFile(ctx context.Context, req *pb.GroupFileReq, resp *pb.GroupFileResp) error {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		return errcode.ErrAuthFailed
	}

	if esUser == nil {
		return errcode.ErrAuthFailed
	}

	if esUser.Mid != req.Uid {
		return errcode.ErrAuthFailed
	}

	return nil
}

func (i *ImHandler) AddGroupPicture(ctx context.Context, req *pb.AddGroupPictureReq, resp *pb.Response) error {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser == nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser.Mid != req.Uid {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	return nil
}

func (i *ImHandler) GetGroupPicture(ctx context.Context, req *pb.GroupFileReq, resp *pb.GroupPictureResp) error {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		return errcode.ErrAuthFailed
	}

	if esUser == nil {
		return errcode.ErrAuthFailed
	}

	if esUser.Mid != req.Uid {
		return errcode.ErrAuthFailed
	}

	return nil
}

func (i *ImHandler) GetUserUnReadCount(ctx context.Context, req *pb.GetUserUnReadCountReq, resp *pb.GetUserUnReadCountResp) (err error) {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		return errcode.ErrAuthFailed
	}

	if esUser == nil {
		return errcode.ErrAuthFailed
	}

	if esUser.Mid != req.Uid {
		return errcode.ErrAuthFailed
	}


	resp.UnreadCount, err = il.GetUserUnReadCount(req.Appid, req.Uid)
	if err != nil {
		return err
	}

	return nil
}

func (i *ImHandler) UpdateGroupAvatar(ctx context.Context, req *pb.UpdateGroupAvatarReq, resp *pb.Response) (err error) {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser == nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser.Mid != req.Uid {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}


	group, err := il.GetGroupById(req.Gid, req.Appid)
	if err != nil {
		return err
	}

	if group == nil {
		resp.Msg = errcode.ErrUpdateGroupAvatarFailed.Msg
		resp.Success = errcode.ErrUpdateGroupAvatarFailed.Code
		return err
	}

	if group.Master != req.Uid {
		resp.Msg = errcode.ErrUpdateGroupNameFailed.Msg
		resp.Success = errcode.ErrUpdateGroupNameFailed.Code
		return err
	}

	group.ImageURL = req.ImageUrl


	err = il.UpdateGroup(group)
	if err != nil {
		return err
	}

	resp.Success = 0
	resp.Msg = "success"

	return nil
}

func (i *ImHandler) UpdateGroupNotice(ctx context.Context, req *pb.UpdateGroupNoticeReq, resp *pb.Response) (err error) {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser == nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser.Mid != req.Uid {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	group, err := il.GetGroupById(req.Gid, req.Appid)
	if err != nil {
		return err
	}

	if group == nil {
		resp.Msg = errcode.ErrUpdateGroupNoticeFailed.Msg
		resp.Success = errcode.ErrUpdateGroupNoticeFailed.Code
		return err
	}

	if group.Master != req.Uid {
		resp.Msg = errcode.ErrUpdateGroupNameFailed.Msg
		resp.Success = errcode.ErrUpdateGroupNameFailed.Code
		return err
	}

	group.Notice = null.String{String:req.Notice, Valid: true}


	err = il.UpdateGroup(group)
	if err != nil {
		return err
	}

	resp.Success = 0
	resp.Msg = "success"

	return nil
}

func (i *ImHandler) UpdateGroupName(ctx context.Context, req *pb.UpdateGroupNameReq, resp *pb.Response) (err error) {
	il := &logic.ImLogic{}

	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser == nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser.Mid != req.Uid {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	//groupExist, err := il.GetGroupByName(req.Name, req.Appid)
	//if err != nil {
	//	resp.Msg = errcode.ErrCreteGroupFailed.Msg
	//	resp.Success = errcode.ErrCreteGroupFailed.Code
	//	return
	//}
	//
	//if groupExist != nil {
	//	resp.Msg = errcode.ErrGroupNameIsExistFailed.Msg
	//	resp.Success = errcode.ErrGroupNameIsExistFailed.Code
	//	return
	//}

	group, err := il.GetGroupById(req.Gid, req.Appid)
	if err != nil {
		return err
	}

	if group == nil {
		resp.Msg = errcode.ErrUpdateGroupNameFailed.Msg
		resp.Success = errcode.ErrUpdateGroupNameFailed.Code
		return err
	}

	if group.Master != req.Uid {
		resp.Msg = errcode.ErrUpdateGroupNameFailed.Msg
		resp.Success = errcode.ErrUpdateGroupNameFailed.Code
		return err
	}

	group.Name = req.Name

	err = il.UpdateGroup(group)
	if err != nil {
		return err
	}

	resp.Success = 0
	resp.Msg = "success"

	return nil
}

func (i *ImHandler) UpdateGroupJoinAuth(ctx context.Context, req *pb.UpdateGroupJoinAuthReq, resp *pb.Response) (err error) {
	il := &logic.ImLogic{}
	esUser, err := il.GetEsUserInfo(req.UserId)
	if err != nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser == nil {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	if esUser.Mid != req.Uid {
		resp.Msg = errcode.ErrAuthFailed.Msg
		resp.Success = errcode.ErrAuthFailed.Code
		return nil
	}

	group, err := il.GetGroupById(req.Gid, req.Appid)
	if err != nil {
		return err
	}

	if group == nil {
		resp.Msg = errcode.ErrUpdateGroupJoinAuthFailed.Msg
		resp.Success = errcode.ErrUpdateGroupJoinAuthFailed.Code
		return err
	}

	if group.Master != req.Uid {
		resp.Msg = errcode.ErrUpdateGroupNameFailed.Msg
		resp.Success = errcode.ErrUpdateGroupNameFailed.Code
		return err
	}


	group.Auth = int8(req.Auth)

	err = il.UpdateGroup(group)
	if err != nil {
		return err
	}

	resp.Success = 0
	resp.Msg = "success"

	return nil
}
