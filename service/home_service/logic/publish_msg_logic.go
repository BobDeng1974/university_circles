package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/olivere/elastic/v7"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"time"
	"university_circles/service/home_service/databases/es"
	"university_circles/service/home_service/utils/common"
	"university_circles/service/home_service/utils/errcode"
	"university_circles/service/home_service/utils/logger"
	"unsafe"

	"github.com/volatiletech/sqlboiler/boil"
	"go.uber.org/zap"
	"university_circles/service/home_service/databases/mysql"
	"university_circles/service/home_service/models"
	myRedis "university_circles/service/user_service/databases/redis"

	pb "university_circles/service/home_service/pb/home"
)

const (
	//
	STUDENTLOGINSESSIONPREKEY = "university_circles:student:login:session:"
	// 某个动态的点赞数
	PUBLISHMSGLIKECOUNTCACHEKEY = "university_circles:publish_msg:like:count"
	// 某个动态的评论数
	PUBLISHMSGCOMMENTCOUNTCACHEKEY = "university_circles:publish_msg:comment:count"
	// 某个动态的分享数
	PUBLISHMSGSHARECOUNTCACHEKEY = "university_circles:publish_msg:share:count"
	// 用户点赞过的动态列表
	PUBLISHMSGLIKEPREKEY = "university_circles:publish_msg:like:user:"
	// 用户评论过的动态列表
	PUBLISHMSGCOMMENTPREKEY = "university_circles:publish_msg:comment:user:"
)

type HomeLogic struct{}

func (hl *HomeLogic) SaveHomeToES(msgId string, orderId time.Time, createTime string, req *pb.PublishMsg) (err error) {
	esClient, err := es.NewElasticSearch()
	if err != nil {
		logger.Logger.Warn("New ElasticSearch failed", zap.Error(err))
		return
	}

	// 查询动态里的用户信息
	getUser, err := esClient.Get().Index(common.HOMEPUBLISHMSGUSER).Id(req.Uid).Do(context.Background())
	if err != nil {
		logger.Logger.Warn("home publish msg user byte convert string failed", zap.Any("req", req), zap.Error(err))
	}

	if getUser == nil {
		return errcode.ErrFailed
	}

	sourceUser, err := getUser.Source.MarshalJSON()
	if err != nil {
		logger.Logger.Warn("home publish msg user byte convert string failed", zap.Any("req", req), zap.Error(err))
	}

	fmt.Println("sourceUser", string(sourceUser))

	esUserInfo := &pb.EsUserInfo{}
	if err = json.Unmarshal(sourceUser, esUserInfo); err != nil {
		logger.Logger.Warn("json Unmarshal home publish msg user failed", zap.Any("esUserInfo", string(sourceUser)), zap.Error(err))
		return
	}

	var pictures []*pb.Picture
	for _, picId := range req.Pictures {
		p := common.OSSFILEURLPREFIX + fmt.Sprintf(common.UPLOADFILEPATHPREFIX, 1, picId.ImageId+picId.ImageId)
		pic := &pb.Picture{
			ThumbnailUrl: p + "?x-oss-process=image/resize,m_fixed,h_300,w_2000",
			MiddlePicUrl: p + "?x-oss-process=image/resize,m_fixed,h_400,w_2000",
			PicUrl:       p + "?x-oss-process=image/resize,m_fixed,h_600,w_2000",
			Format:       picId.Format,
			Width:        picId.Width,
			Height:       picId.Height,
		}
		pictures = append(pictures, pic)
	}

	msgRep := &pb.PublishMsgResponse{
		Id:        msgId,
		OrderId:   orderId.Unix(),
		User:      esUserInfo,
		MsgType:   req.MsgType,
		Content:   req.Content,
		Pictures:  pictures,
		LinkInfo:  req.LinkInfo,
		Video:     req.Video,
		Poi:       req.Poi,
		Type:      "ORIGINAL_POST",
		Status:    "NORMAL",
		CreatedAt: createTime,
	}
	fmt.Println("Response", msgRep)
	msgRep.IsCommentForbidden = false
	msgRep.Collected = false

	// 添加动态到es
	_, err = esClient.Index().
		Index(common.HOMEPUBLISHMSG).
		Id(msgRep.Id).
		BodyJson(msgRep).
		Do(context.Background())
	fmt.Println("esClient add", msgRep, err)
	if err != nil {
		logger.Logger.Warn("save home publish msg to es failed", zap.Any("Home", msgRep), zap.Error(err))
		return
	}

	return nil
}

func (hl *HomeLogic) AsyncSaveHomeToDB(msgId string, req *pb.PublishMsg) (err error) {
	var uHome *home.Home
	uHome = &home.Home{
		MSGID:   msgId,
		UID:     req.Uid,
		Content: req.Content,
		MSGType: int8(req.MsgType),
	}

	if err = hl.SaveHome(uHome); err != nil {
		logger.Logger.Warn("save home publish msg to db failed", zap.Any("Home", uHome), zap.Error(err))
		return
	}

	// 存储动态图片url到db
	if req.Pictures != nil {

		for _, picId := range req.Pictures {
			p := common.OSSFILEURLPREFIX + fmt.Sprintf(common.UPLOADFILEPATHPREFIX, 1, picId.ImageId+picId.ImageId)
			msgPicture := &home.HomePicture{
				OwnerID:      msgId,
				ThumbnailUrl: p + "?x-oss-process=image/resize,m_fixed,h_300,w_2000",
				MiddlePicUrl: p + "?x-oss-process=image/resize,m_fixed,h_400,w_2000",
				PicUrl:       p + "?x-oss-process=image/resize,m_fixed,h_600,w_2000",
				Type:         1,
				Width:        *(*int)(unsafe.Pointer(&picId.Width)),
				Height:       *(*int)(unsafe.Pointer(&picId.Height)),
				Format:       picId.Format,
			}

			if err = hl.SavePicture(msgPicture); err != nil {
				logger.Logger.Warn("save home publish msg picture to db failed", zap.Any("picture", picId),
					zap.Error(err))
				return
			}
		}
	}

	// 存储动态地址到db
	if req.Poi != nil {
		var loc string
		if len(req.Poi.Location) > 2 {
			loc = fmt.Sprintf("%f,%f", req.Poi.Location[0].Loc, req.Poi.Location[1].Loc)
		} else {
			loc = ""
		}

		poi := &home.HomePoi{
			OwnerID:          msgId,
			Location:         loc,
			PoiId:            req.Poi.PoiId,
			Countryname:      req.Poi.Countryname,
			Pname:            req.Poi.Pname,
			Cityname:         req.Poi.Cityname,
			Name:             req.Poi.Name,
			FormattedAddress: req.Poi.FormattedAddress,
		}

		if err = hl.SavePoi(poi); err != nil {
			logger.Logger.Warn("save home publish msg poi to db failed", zap.Any("poi", poi), zap.Error(err))
			return
		}
	}

	// 保存关联URL和视频
	var video *home.HomeVideo
	if req.Video != nil {
		msgVideoPicture := &home.HomePicture{
			OwnerID:      msgId,
			ThumbnailUrl: req.Video.Image.ThumbnailUrl,
			MiddlePicUrl: req.Video.Image.MiddlePicUrl,
			PicUrl:       req.Video.Image.PicUrl,
			Type:         2,
		}

		if err = hl.SavePicture(msgVideoPicture); err != nil {
			logger.Logger.Warn("save home publish msg video picture to db failed", zap.Any("picture", msgVideoPicture),
				zap.Error(err))
			return
		}

		video = &home.HomeVideo{
			OwnerID:  msgId,
			Type:     "VIDEO",
			PicID:    int(msgVideoPicture.ID),
			Duration: int(req.Video.Duration),
			Width:    int(req.Video.Width),
			Height:   int(req.Video.Height),
		}
		if err = hl.SaveVideo(video); err != nil {
			logger.Logger.Warn("save home publish msg video to db failed", zap.Any("video", video), zap.Error(err))
			return
		}
	}

	if req.LinkInfo != nil {
		linkInfo := &home.HomeLinkinfo{
			MSGID:      msgId,
			Title:      req.LinkInfo.Title,
			PictureUrl: req.LinkInfo.PictureUrl,
			LinkUrl:    req.LinkInfo.LinkUrl,
			Source:     req.LinkInfo.Source,
			VideoID:    int(video.ID),
		}
		if err = hl.SaveLinkInfo(linkInfo); err != nil {
			logger.Logger.Warn("save home publish msg link info to db failed", zap.Any("linkInfo", linkInfo), zap.Error(err))
			return
		}

	}

	return
}

func (hl *HomeLogic) SaveHome(hpm *home.Home) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}

	if err = hpm.Insert(context.Background(), db, boil.Infer()); err != nil {
		logger.Logger.Warn("insert home publish msg failed", zap.Any("home", hpm), zap.Error(err))
		return
	}
	fmt.Println("logic", hpm)

	return nil
}

func (hl *HomeLogic) SavePicture(pic *home.HomePicture) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	if err = pic.Insert(context.Background(), db, boil.Infer()); err != nil {
		logger.Logger.Warn("insert home publish msg picture failed", zap.Any("picture", pic), zap.Error(err))
		return
	}
	return nil
}

func (hl *HomeLogic) SaveUrlsInText(urlsInText *home.HomeUrlsInText) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	if err = urlsInText.Insert(context.Background(), db, boil.Infer()); err != nil {
		logger.Logger.Warn("insert home publish msg urlsInText failed", zap.Any("urlsInText", urlsInText), zap.Error(err))
		return
	}
	return nil
}

func (hl *HomeLogic) SavePoi(poi *home.HomePoi) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	if err = poi.Insert(context.Background(), db, boil.Infer()); err != nil {
		logger.Logger.Warn("insert home publish msg picture failed", zap.Any("poi", poi), zap.Error(err))
		return
	}
	return nil
}

func (hl *HomeLogic) SaveLinkInfo(linkInfo *home.HomeLinkinfo) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	if err = linkInfo.Insert(context.Background(), db, boil.Infer()); err != nil {
		logger.Logger.Warn("insert home publish msg picture failed", zap.Any("linkInfo", linkInfo), zap.Error(err))
		return
	}
	return nil
}

func (hl *HomeLogic) SaveVideo(video *home.HomeVideo) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	if err = video.Insert(context.Background(), db, boil.Infer()); err != nil {
		logger.Logger.Warn("insert home publish msg video failed", zap.Any("video", video), zap.Error(err))
		return
	}
	return nil
}

func (hl *HomeLogic) Delete(uid, msgId string) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	esClient, err := es.NewElasticSearch()
	if err != nil {
		logger.Logger.Warn("New ElasticSearch failed", zap.Error(err))
	}

	//res, err := esClient.Get().
	//	Index(common.HOMEPUBLISHMSG).
	//	Id(msgId).
	//	Do(context.Background())
	//fmt.Println(res, err)
	//if err != nil {
	//	logger.Logger.Warn("del home_publish_msg comment from es failed", zap.Any("msuidgId", uid), zap.Any("msgId", msgId), zap.Error(err))
	//	return
	//
	//}
	//if res.Source == nil {
	//	logger.Logger.Warn("delete publish msg comment from es failed", zap.Any("uid", uid), zap.Any("msgId", msgId), zap.Error(err))
	//	return
	//}
	//
	//hsmOrigin := &pb.Response{}
	//err = json.Unmarshal(res.Source, &hsmOrigin)
	//logger.Logger.Warn("json Unmarshal home_publish_msg from es failed", zap.Any("source", string(res.Source)), zap.Error(err))
	//
	//// 如果请求删除动态的用户与动态的用户不一致，则返回错误
	//if hsmOrigin.User != nil {
	//	if hsmOrigin.User.Id != uid {
	//		return errcode.ErrDelFailed
	//	}
	//}

	go func() {
		hpm, err := home.Homes(qm.Where("uid = ? and msg_id = ?", uid, msgId)).One(context.Background(), db)
		hpm.Deleted = 1
		_, err = hpm.Update(context.Background(), db, boil.Infer())
		if err != nil {
			logger.Logger.Warn("delete publish msg from db failed", zap.Any("UHome", hpm), zap.Error(err))
			return
		}
		return
	}()

	// 删除es动态
	scriptStr := "ctx._source.status = ctx._source.user.id=params.uid ? 'DELETED': 'NORMAL'"
	userScript := elastic.NewScriptInline(scriptStr).Param("uid", uid)

	_, err = esClient.Update().
		Index(common.HOMEPUBLISHMSG).
		Id(msgId).
		Script(userScript).
		//Doc(map[string]interface{}{"status": "DELETED"}).
		Do(context.Background())
	if err != nil {
		logger.Logger.Warn("delete publish msg from es failed", zap.Any("msgId", msgId), zap.Error(err))
		return
	}

	return nil
}

func (hl *HomeLogic) GetHomeMsgList(msgType, orderId int64, loadMore bool, uid string, universityId int64) (hpmList []*pb.PublishMsgListResponse, err error) {
	esClient, err := es.NewElasticSearch()
	if err != nil {
		logger.Logger.Warn("New ElasticSearch failed", zap.Error(err))
	}

	msgTypeList := make([]interface{}, 2)
	if msgType == 0 || msgType == 1 {
		msgTypeList[0] = 1
		msgTypeList[1] = 4
	} else {
		msgTypeList[0] = 2
		msgTypeList[1] = 4
	}

	fmt.Println("orderId", orderId, loadMore, msgTypeList)
	var res *elastic.SearchResult
	msgTypeQuery := elastic.NewTermsQuery("msg_type", msgTypeList...)
	// 校内，可以查看校内和校外的
	if msgType == 1 {
		universityQuery := elastic.NewMatchQuery("user.university_id", universityId)
		query := elastic.NewBoolQuery().Must(msgTypeQuery, universityQuery)
		if loadMore && orderId != 0 {
			// 下拉加载
			res, err = esClient.Search(common.HOMEPUBLISHMSG).SearchAfter(orderId).Query(query).From(0).Sort("orderId", false).Size(10).
				Do(context.Background())
		} else if !loadMore && orderId != 0 {
			// 上拉加载
			res, err = esClient.Search(common.HOMEPUBLISHMSG).SearchAfter(orderId).Query(query).From(0).Sort("orderId", true).Size(10).
				Do(context.Background())
		} else {
			// 默认
			res, err = esClient.Search(common.HOMEPUBLISHMSG).Query(query).From(0).Sort("orderId", false).Size(10).
				Do(context.Background())
		}

	} else if msgType == 3 {
		// 获取关注的用户列表，筛选对应用户数据
		userFollowList, _ := hl.GetUserFollowList(uid)
		followInterface := make([]interface{}, len(userFollowList))
		var followStr string
		for index, value := range userFollowList {
			followInterface[index] = string(value)
			followStr += value + ","
		}
		fmt.Println("followStr", followStr)
		userFollowQuery := elastic.NewMatchQuery("user.id", followStr)

		// 校内的动态需满足同学校
		msgQuery := elastic.NewMatchQuery("msg_type", 1)
		msgUniversityQuery := elastic.NewMatchQuery("user.university_id", universityId)

		// 关注所有人 and (校内的动态 and 同学校) or (校外或校内外的消息)
		universityQuery := elastic.NewBoolQuery().Must(userFollowQuery, msgQuery, msgUniversityQuery)

		msgTypeQuery := elastic.NewTermsQuery("msg_type", msgTypeList...)

		followMsgTypeQuery := elastic.NewBoolQuery().Must(userFollowQuery, msgTypeQuery)

		followQuery := elastic.NewBoolQuery().Should(universityQuery, followMsgTypeQuery)

		if loadMore && orderId != 0 {
			// 下拉加载
			res, err = esClient.Search(common.HOMEPUBLISHMSG).SearchAfter(orderId).Query(followQuery).From(0).Sort("orderId", false).Size(10).
				Do(context.Background())
		} else if !loadMore && orderId != 0 {
			// 上拉加载
			res, err = esClient.Search(common.HOMEPUBLISHMSG).SearchAfter(orderId).Query(followQuery).From(0).Sort("orderId", true).Size(10).
				Do(context.Background())
		} else {
			// 默认
			res, err = esClient.Search(common.HOMEPUBLISHMSG).Query(followQuery).Sort("orderId", false).Size(10).
				Do(context.Background())
		}
	} else {
		// 所有学校外部的都可以看见
		if loadMore && orderId != 0 {
			// 下拉加载
			res, err = esClient.Search(common.HOMEPUBLISHMSG).SearchAfter(orderId).Query(msgTypeQuery).From(0).Sort("orderId", false).Size(10).
				Do(context.Background())
		} else if !loadMore && orderId != 0 {
			// 上拉加载
			res, err = esClient.Search(common.HOMEPUBLISHMSG).SearchAfter(orderId).Query(msgTypeQuery).From(0).Sort("orderId", true).Size(10).
				Do(context.Background())
		} else {
			// 默认
			res, err = esClient.Search(common.HOMEPUBLISHMSG).Query(msgTypeQuery).From(0).Sort("orderId", false).Size(10).
				Do(context.Background())
		}

	}

	if err != nil {
		logger.Logger.Warn("get home_publish_msg list from es failed", zap.Any("msgType", msgType), zap.Any("orderId", orderId), zap.Any("loadMore", loadMore), zap.Error(err))
		return
	}

	if len(res.Hits.Hits) > 0 {
		for _, hit := range res.Hits.Hits {
			pml := &pb.PublishMsgListResponse{}
			hsm := &pb.PublishMsgResponse{}
			err = json.Unmarshal(hit.Source, &hsm)
			logger.Logger.Warn("json Unmarshal home_publish_msg from es failed", zap.Any("source", string(hit.Source)), zap.Error(err))

			if hsm.User != nil {
				hsm.Liked, _ = hl.CheckUserLikeFromCache(hsm.Id, uid)
			}
			// 增加评论数、点赞数等
			hsm.CommentCount, _ = hl.GetCommentCountFromCache(hsm.Id)
			hsm.LikeCount, _ = hl.GetLikeCountFromCache(hsm.Id)
			hsm.ShareCount, _ = hl.GetShareCountFromCache(hsm.Id)

			// 查询动态里的用户信息
			getUser, err := esClient.Get().Index(common.HOMEPUBLISHMSGUSER).Id(hsm.User.Id).Do(context.Background())
			if err != nil {
				logger.Logger.Warn("get home publish msg user byte convert string failed", zap.Any("hsm", hsm), zap.Error(err))
			}

			sourceUser, err := getUser.Source.MarshalJSON()
			if err != nil {
				logger.Logger.Warn("get home publish msg user byte convert string failed", zap.Any("hsm", hsm), zap.Error(err))
			}

			fmt.Println("sourceUser", string(sourceUser))

			esUserInfo := &pb.EsUserInfo{}
			if err = json.Unmarshal(sourceUser, esUserInfo); err != nil {
				logger.Logger.Warn("json Unmarshal home publish msg user failed", zap.Any("esUserInfo", string(sourceUser)), zap.Error(err))
			}

			hsm.User = esUserInfo

			hsm.IsFollow = false
			if userIsFollow, _ := hl.CheckUserIsFollow(uid, esUserInfo.Id); userIsFollow {
				hsm.IsFollow = true
			}

			pml.Id = hsm.Id
			pml.Item = hsm
			pml.Type = "RECOMMENDED_MESSAGE"
			hpmList = append(hpmList, pml)
		}

	}

	return
}

func (hl *HomeLogic) GetUserMsgList(uid string, orderId, universityId, msgType int64) (hpmList []*pb.PublishMsgResponse, err error) {
	esClient, err := es.NewElasticSearch()
	if err != nil {
		logger.Logger.Warn("New ElasticSearch failed", zap.Error(err))
	}
	var res *elastic.SearchResult

	msgTypeList := make([]interface{}, 2)
	if msgType == 0 || msgType == 1 {
		msgTypeList[0] = 1
		msgTypeList[1] = 4
	} else {
		msgTypeList[0] = 2
		msgTypeList[1] = 4
	}
	msgTypeQuery := elastic.NewTermsQuery("msg_type", msgTypeList...)

	fmt.Println("uid ", uid, orderId)
	userQuery := elastic.NewMatchQuery("user.id", uid)

	userMsgTypeQuery := elastic.NewBoolQuery().Must(userQuery, msgTypeQuery)

	if orderId != 0 {
		res, err = esClient.Search(common.HOMEPUBLISHMSG).SearchAfter(orderId).Query(userMsgTypeQuery).From(0).Sort("orderId", false).Size(10).
			Do(context.Background())
	} else {
		res, err = esClient.Search(common.HOMEPUBLISHMSG).Query(userMsgTypeQuery).Sort("orderId", false).Size(10).
			Do(context.Background())
	}

	if err != nil {
		logger.Logger.Warn("get home_publish_msg list from es failed", zap.Any("uid", uid), zap.Any("orderId", orderId), zap.Error(err))
		return

	}

	if len(res.Hits.Hits) > 0 {
		for _, hit := range res.Hits.Hits {
			hsm := &pb.PublishMsgResponse{}
			err = json.Unmarshal(hit.Source, &hsm)
			logger.Logger.Warn("json Unmarshal home_publish_msg from es failed", zap.Any("source", string(hit.Source)), zap.Error(err))

			if hsm.User != nil {
				hsm.Liked, _ = hl.CheckUserLikeFromCache(hsm.Id, uid)
			}
			// 增加评论数、点赞数等
			hsm.CommentCount, _ = hl.GetCommentCountFromCache(hsm.Id)
			hsm.LikeCount, _ = hl.GetLikeCountFromCache(hsm.Id)
			hsm.ShareCount, _ = hl.GetShareCountFromCache(hsm.Id)

			// 查询动态里的用户信息
			getUser, err := esClient.Get().Index(common.HOMEPUBLISHMSGUSER).Id(hsm.User.Id).Do(context.Background())
			if err != nil {
				logger.Logger.Warn("get home publish msg user byte convert string failed", zap.Any("hsm", hsm), zap.Error(err))
			}

			sourceUser, err := getUser.Source.MarshalJSON()
			if err != nil {
				logger.Logger.Warn("get home publish msg user byte convert string failed", zap.Any("hsm", hsm), zap.Error(err))
			}

			fmt.Println("sourceUser", string(sourceUser))

			esUserInfo := &pb.EsUserInfo{}
			if err = json.Unmarshal(sourceUser, esUserInfo); err != nil {
				logger.Logger.Warn("json Unmarshal home publish msg user failed", zap.Any("esUserInfo", string(sourceUser)), zap.Error(err))
			}

			hsm.User = esUserInfo

			hsm.IsFollow = false
			if userIsFollow, _ := hl.CheckUserIsFollow(uid, esUserInfo.Id); userIsFollow {
				hsm.IsFollow = true
			}

			hpmList = append(hpmList, hsm)
		}

	}

	return
}

func (hl *HomeLogic) GetMsgDetail(uid, msgId string, universityId int64) (hpmList []*pb.PublishMsgResponse, err error) {
	esClient, err := es.NewElasticSearch()
	if err != nil {
		logger.Logger.Warn("New ElasticSearch failed", zap.Error(err))
	}
	var res *elastic.SearchResult

	msgQuery := elastic.NewMatchQuery("id", msgId)
	userQuery := elastic.NewMatchQuery("user.id", uid)

	userMsgQuery := elastic.NewBoolQuery().Must(userQuery, msgQuery)

	res, err = esClient.Search(common.HOMEPUBLISHMSG).Query(userMsgQuery).Size(1).
		Do(context.Background())

	if err != nil {
		logger.Logger.Warn("get home_publish_msg list from es failed", zap.Any("uid", uid), zap.Error(err))
		return

	}

	if len(res.Hits.Hits) > 0 {
		hsm := &pb.PublishMsgResponse{}
		err = json.Unmarshal(res.Hits.Hits[0].Source, &hsm)
		logger.Logger.Warn("json Unmarshal home_publish_msg from es failed", zap.Any("source", string(res.Hits.Hits[0].Source)), zap.Error(err))

		if hsm.User != nil {
			hsm.Liked, _ = hl.CheckUserLikeFromCache(hsm.Id, uid)
		}
		// 增加评论数、点赞数等
		hsm.CommentCount, _ = hl.GetCommentCountFromCache(hsm.Id)
		hsm.LikeCount, _ = hl.GetLikeCountFromCache(hsm.Id)
		hsm.ShareCount, _ = hl.GetShareCountFromCache(hsm.Id)

		// 查询动态里的用户信息
		var getUser *elastic.GetResult
		getUser, err = esClient.Get().Index(common.HOMEPUBLISHMSGUSER).Id(hsm.User.Id).Do(context.Background())
		if err != nil {
			logger.Logger.Warn("get home publish msg user byte convert string failed", zap.Any("hsm", hsm), zap.Error(err))
			return
		}

		var sourceUser []byte
		sourceUser, err = getUser.Source.MarshalJSON()
		if err != nil {
			logger.Logger.Warn("get home publish msg user byte convert string failed", zap.Any("hsm", hsm), zap.Error(err))
			return
		}

		fmt.Println("sourceUser", string(sourceUser))

		var esUserInfo *pb.EsUserInfo
		if err = json.Unmarshal(sourceUser, &esUserInfo); err != nil {
			logger.Logger.Warn("json Unmarshal home publish msg user failed", zap.Any("esUserInfo", string(sourceUser)), zap.Error(err))
			return
		}

		hsm.User = esUserInfo

		hsm.IsFollow = false
		if userIsFollow, _ := hl.CheckUserIsFollow(uid, esUserInfo.Id); userIsFollow {
			hsm.IsFollow = true
		}

		hpmList = append(hpmList, hsm)

	}

	return
}

func (hl *HomeLogic) GetDetail(uid, msgId string, universityId int64) (hpm *pb.PublishMsgResponse, err error) {
	esClient, err := es.NewElasticSearch()
	if err != nil {
		logger.Logger.Warn("New ElasticSearch failed", zap.Error(err))
	}

	res, err := esClient.Get().
		Index(common.HOMEPUBLISHMSG).
		Id(msgId).
		Do(context.Background())
	if err != nil {
		logger.Logger.Warn("get home_publish_msg detail from es failed", zap.Any("msgId", msgId), zap.Error(err))
		return

	}
	source, err := res.Source.MarshalJSON()
	if err != nil {
		logger.Logger.Warn("get home_publish_msg detail Marshal JSON from es failed", zap.Any("msgId", msgId), zap.Error(err))
		return

	}
	hpm = &pb.PublishMsgResponse{}
	err = json.Unmarshal(source, &hpm)
	logger.Logger.Warn("json Unmarshal home_publish_msg from es failed", zap.Any("source", string(res.Source)), zap.Error(err))

	// 校内的动态需判断同学校
	if hpm.MsgType == 1 {
		if hpm.User.UniversityId != universityId {
			logger.Logger.Warn("get home_publish_msg detail by different university failed", zap.Any("msg", hpm), zap.Any("universityId", universityId), zap.Error(err))
			return
		}
	}

	if hpm.User != nil {
		hpm.Liked, _ = hl.CheckUserLikeFromCache(hpm.Id, uid)
	}
	// 增加评论数、点赞数等
	hpm.CommentCount, _ = hl.GetCommentCountFromCache(hpm.Id)
	hpm.LikeCount, _ = hl.GetLikeCountFromCache(hpm.Id)
	hpm.ShareCount, _ = hl.GetShareCountFromCache(hpm.Id)

	// 查询动态里的用户信息
	getUser, err := esClient.Get().Index(common.HOMEPUBLISHMSGUSER).Id(hpm.User.Id).Do(context.Background())
	if err != nil {
		logger.Logger.Warn("get home publish msg user byte convert string failed", zap.Any("hpm", hpm), zap.Error(err))
	}

	sourceUser, err := getUser.Source.MarshalJSON()
	if err != nil {
		logger.Logger.Warn("get home publish msg user byte convert string failed", zap.Any("hpm", hpm), zap.Error(err))
	}

	fmt.Println("sourceUser", string(sourceUser))

	esUserInfo := &pb.EsUserInfo{}
	if err = json.Unmarshal(sourceUser, esUserInfo); err != nil {
		logger.Logger.Warn("json Unmarshal home publish msg user failed", zap.Any("esUserInfo", string(sourceUser)), zap.Error(err))
	}

	hpm.User = esUserInfo

	return hpm, nil
}

func (hl *HomeLogic) AddLikeCountToCache(msgId, uid string) (err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	isLiked, err := hl.CheckUserLikeFromCache(msgId, uid)
	if err != nil {
		logger.Logger.Warn("save publish msg like count to cache failed", zap.Any("msgId", msgId), zap.Any("uid", uid), zap.Error(err))
		return
	}

	if !isLiked {
		cacheKey := PUBLISHMSGLIKECOUNTCACHEKEY
		_, err = rd.Do("HIncrBy", cacheKey, msgId, 1)
		if err != nil {
			logger.Logger.Warn("save publish msg like count to cache failed", zap.Any("cacheKey", cacheKey), zap.Error(err))
			return
		}
	}

	return
}

func (hl *HomeLogic) SubLikeCountToCache(msgId string) (err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := PUBLISHMSGLIKECOUNTCACHEKEY

	count, err := redis.Int64(rd.Do("HGet", cacheKey, msgId))
	if err != nil {
		if err.Error() != "redigo: nil returned" {
			logger.Logger.Warn("get publish msg like count from cache failed", zap.Any("cacheKey", cacheKey), zap.Error(err))
			return
		}
	}
	if count >= 0 {
		_, err = rd.Do("HIncrBy", cacheKey, msgId, -1)
		if err != nil {
			logger.Logger.Warn("save publish msg like count to cache failed", zap.Any("cacheKey", cacheKey), zap.Error(err))
			return
		}
	}

	return
}

func (hl *HomeLogic) GetLikeCountFromCache(msgId string) (count int64, err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := PUBLISHMSGLIKECOUNTCACHEKEY
	count, err = redis.Int64(rd.Do("HGet", cacheKey, msgId))
	if err != nil {
		if err.Error() != "redigo: nil returned" {
			logger.Logger.Warn("get publish msg like count from cache failed", zap.Any("cacheKey", cacheKey), zap.Error(err))
			return
		}
	}

	return count, nil
}

func (hl *HomeLogic) AddCommentCountToCache(msgId string) (err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := PUBLISHMSGCOMMENTCOUNTCACHEKEY
	_, err = rd.Do("HIncrBy", cacheKey, msgId, 1)
	if err != nil {
		logger.Logger.Warn("add publish msg comment count to cache failed", zap.Any("cacheKey", cacheKey), zap.Error(err))
		return
	}

	return
}

func (hl *HomeLogic) SubCommentCountToCache(msgId string) (err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := PUBLISHMSGCOMMENTCOUNTCACHEKEY
	_, err = rd.Do("HIncrBy", cacheKey, msgId, -1)
	if err != nil {
		logger.Logger.Warn("sub publish msg comment count to cache failed", zap.Any("cacheKey", cacheKey), zap.Error(err))
		return
	}

	return
}

func (hl *HomeLogic) GetCommentCountFromCache(msgId string) (count int64, err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := PUBLISHMSGCOMMENTCOUNTCACHEKEY
	count, err = redis.Int64(rd.Do("HGet", cacheKey, msgId))
	if err != nil {
		if err.Error() != "redigo: nil returned" {
			logger.Logger.Warn("get publish msg comment count from cache failed", zap.Any("cacheKey", cacheKey), zap.Error(err))
			return
		}
	}

	return count, nil
}

func (hl *HomeLogic) AddShareCountToCache(msgId string) (err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := PUBLISHMSGSHARECOUNTCACHEKEY
	_, err = rd.Do("HIncrBy", cacheKey, msgId, 1)
	if err != nil {
		logger.Logger.Warn("add publish msg share count to cache failed", zap.Any("cacheKey", cacheKey), zap.Error(err))
		return
	}

	return
}

func (hl *HomeLogic) GetShareCountFromCache(msgId string) (count int64, err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := PUBLISHMSGSHARECOUNTCACHEKEY
	count, err = redis.Int64(rd.Do("HGet", cacheKey, msgId))
	if err != nil {
		if err.Error() != "redigo: nil returned" {
			logger.Logger.Warn("get publish msg share count from cache failed", zap.Any("cacheKey", cacheKey), zap.Error(err))
			return
		}
	}

	return count, nil
}

// 增加用户点赞过的某一个动态
func (hl *HomeLogic) AddUserLikeToCache(msgId, uid string) (err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := PUBLISHMSGLIKEPREKEY + uid
	_, err = rd.Do("zAdd", cacheKey, time.Now().Local().Unix(), msgId)
	if err != nil {
		logger.Logger.Warn("save user like publish msg to cache failed", zap.Any("cacheKey", cacheKey), zap.Error(err))
		return
	}

	return
}

// 删除用户点赞过的某一个动态
func (hl *HomeLogic) SubUserLikeToCache(msgId, uid string) (err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := PUBLISHMSGLIKEPREKEY + uid
	_, err = rd.Do("zRem", cacheKey, msgId)
	if err != nil {
		logger.Logger.Warn("del user like publish msg to cache failed", zap.Any("cacheKey", cacheKey), zap.Error(err))
		return
	}

	return
}

// 获取用户点赞的动态ID列表
func (hl *HomeLogic) GetUserLikeListFromCache(uid string, page, pageNum int) (userMsgIdList []int64, err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := PUBLISHMSGLIKEPREKEY + uid
	var start, stop int
	start = page * pageNum

	if pageNum <= 0 {
		stop = -1
	} else {
		stop = (page + 1) * pageNum
	}
	userMsgIdList, err = redis.Int64s(rd.Do("zRevRange", cacheKey, start, stop))
	if err != nil {
		if err.Error() != "redigo: nil returned" {
			logger.Logger.Warn("get user like publish msg user  to cache failed", zap.Any("cacheKey", cacheKey), zap.Error(err))
			return
		}
	}

	return userMsgIdList, nil
}

// 查询用户是否点赞了动态
func (hl *HomeLogic) CheckUserLikeFromCache(msgId, uid string) (liked bool, err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	liked = false
	cacheKey := PUBLISHMSGLIKEPREKEY + uid

	var isLiked int
	isLiked, err = redis.Int(rd.Do("zRank", cacheKey, msgId))
	if err != nil {
		logger.Logger.Warn("check user is like publish msg user  to cache failed", zap.Any("cacheKey", cacheKey), zap.Any("msgId", msgId), zap.Error(err))
		return
	}

	if isLiked >= 0 {
		liked = true
	}

	return liked, nil
}

// 用户评论过得动态
func (hl *HomeLogic) SaveUserCommentMapToCache(msgId, uid string) (err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := PUBLISHMSGCOMMENTPREKEY + uid
	_, err = rd.Do("zAdd", cacheKey, time.Now().Local().Unix(), msgId)
	if err != nil {
		logger.Logger.Warn("save user comment publish msg to cache failed", zap.Any("cacheKey", cacheKey), zap.Error(err))
		return
	}

	return nil
}

// 获取用户评论过的动态ID列表
func (hl *HomeLogic) GetUserCommentMapFromCache(msgId, uid string, page, pageNum int) (userMsgIdList []int64, err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := PUBLISHMSGCOMMENTPREKEY + uid
	var start, stop int
	start = page * pageNum

	if pageNum <= 0 {
		stop = -1
	} else {
		stop = (page + 1) * pageNum
	}
	userMsgIdList, err = redis.Int64s(rd.Do("zRevRange", cacheKey, start, stop))
	if err != nil {
		if err.Error() != "redigo: nil returned" {
			logger.Logger.Warn("get user comment  publish msg list  to cache failed", zap.Any("cacheKey", cacheKey), zap.Error(err))
			return
		}
	}

	return userMsgIdList, nil
}

// 用户关注列表
func (hl *HomeLogic) GetUserFollowList(uid string) (userFollowList []string, err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	followCacheKey := common.USERFOLLOWList + uid
	userFollowList, err = redis.Strings(rd.Do("zRevRange", followCacheKey, 0, -1))
	fmt.Println("GetUserFollowList", userFollowList, err)
	if err != nil {
		logger.Logger.Warn("get user follow list into redis failed", zap.Any("followCacheKey", followCacheKey), zap.Error(err))
		return
	}

	return userFollowList, nil
}

// 用户是否关注
func (hl *HomeLogic) CheckUserIsFollow(uid, follower string) (userIsFollow bool, err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	userIsFollow = false
	followCacheKey := common.USERFOLLOWList + uid
	userFollow, err := redis.Int(rd.Do("zRank", followCacheKey, follower))
	fmt.Println("CheckUserIsFollow", userFollow, err)
	if err != nil {
		logger.Logger.Warn("get user follow list into redis failed", zap.Any("followCacheKey", followCacheKey), zap.Error(err))
		return
	}

	if userFollow >= 0 {
		userIsFollow = true
	}

	return userIsFollow, nil
}

func (hl *HomeLogic) SaveMsgCommentToDB(comment *home.HomeComment) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	if err = comment.Insert(context.Background(), db, boil.Infer()); err != nil {
		logger.Logger.Warn("insert home publish msg comment failed", zap.Any("comment", comment), zap.Error(err))
		return
	}

	return nil
}

func (hl *HomeLogic) SaveMsgCommentToEs(comment *pb.PublishMsgComment) (err error) {
	esClient, err := es.NewElasticSearch()
	if err != nil {
		logger.Logger.Warn("New ElasticSearch failed", zap.Error(err))
	}

	_, err = esClient.Index().
		Index(common.HOMEPUBLISHMSGCOMMENT).
		Id(comment.Id).
		BodyJson(comment).
		Do(context.Background())
	fmt.Println("esClient add comment", comment, err)
	if err != nil {
		logger.Logger.Warn("save home publish msg to es failed", zap.Any("comment", comment), zap.Error(err))
		return
	}

	return nil
}

func (hl *HomeLogic) SaveMsgCommentPicture(commentPic *home.HomeCommentPicture) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	if err = commentPic.Insert(context.Background(), db, boil.Infer()); err != nil {
		logger.Logger.Warn("insert home publish msg comment failed", zap.Any("commentPic", commentPic), zap.Error(err))
		return
	}

	return nil
}

func (hl *HomeLogic) SaveMsgCommentUrlInText(commentUrlInText *home.HomeCommentUrlsInText) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	if err = commentUrlInText.Insert(context.Background(), db, boil.Infer()); err != nil {
		logger.Logger.Warn("insert home publish msg comment failed", zap.Any("commentPic", commentUrlInText), zap.Error(err))
		return
	}

	return nil
}

func (hl *HomeLogic) GetCommentList(msgId string, orderId int64) (hpmCommentList []*pb.PublishMsgComment, err error) {
	esClient, err := es.NewElasticSearch()
	if err != nil {
		logger.Logger.Warn("New ElasticSearch failed", zap.Error(err))
	}
	var res *elastic.SearchResult

	msgIdQuery := elastic.NewMatchQuery("targetId", msgId)
	statusQuery := elastic.NewMatchQuery("status", "NORMAL")

	commentQuery := elastic.NewBoolQuery().Must(msgIdQuery, statusQuery)
	if orderId != 0 {
		res, err = esClient.Search(common.HOMEPUBLISHMSGCOMMENT).SearchAfter(orderId).Query(commentQuery).From(0).Sort("orderId", false).Size(10).
			Do(context.Background())
	} else {
		res, err = esClient.Search(common.HOMEPUBLISHMSGCOMMENT).Query(commentQuery).From(0).Sort("orderId", false).Size(10).
			Do(context.Background())
	}

	if err != nil {
		logger.Logger.Warn("get home publish msg comment list from es failed", zap.Any("msgId", msgId), zap.Any("orderId", orderId), zap.Error(err))
		return

	}

	if len(res.Hits.Hits) > 0 {
		for _, hit := range res.Hits.Hits {
			hsmComment := &pb.PublishMsgComment{}
			err = json.Unmarshal(hit.Source, &hsmComment)
			logger.Logger.Warn("json Unmarshal home_publish_msg comment from es failed", zap.Any("source", string(hit.Source)), zap.Error(err))

			hpmCommentList = append(hpmCommentList, hsmComment)
		}

	}

	return hpmCommentList, nil
}

func (hl *HomeLogic) DeleteComment(uid, commentId string) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	esClient, err := es.NewElasticSearch()
	if err != nil {
		logger.Logger.Warn("New ElasticSearch failed", zap.Error(err))
	}

	//res, err := esClient.Get().
	//	Index(common.HOMEPUBLISHMSGCOMMENT).
	//	Id(commentId).
	//	Do(context.Background())
	//fmt.Println(res, err)
	//if err != nil {
	//	logger.Logger.Warn("del home_publish_msg comment from es failed", zap.Any("commentId", commentId), zap.Error(err))
	//	return
	//
	//}
	//if res.Source == nil {
	//	logger.Logger.Warn("delete publish msg comment from es failed", zap.Any("commentId", commentId), zap.Error(err))
	//	return
	//}
	//
	//hsmComm := &pb.Comment{}
	//err = json.Unmarshal(res.Source, &hsmComm)
	//logger.Logger.Warn("json Unmarshal home_publish_msg from es failed", zap.Any("source", string(res.Source)), zap.Error(err))
	//
	//// 如果请求删除评论的用户与评论的用户不一致，则返回错误
	//if hsmComm.User != nil {
	//	if hsmComm.User.Id != uid {
	//		return errcode.ErrDelCommentFailed
	//	}
	//}

	go func() {
		hpmComment, err := home.HomeComments(qm.Where("uid = ? and comm_id = ?", uid, commentId)).One(context.Background(), db)
		hpmComment.Deleted = 1
		_, err = hpmComment.Update(context.Background(), db, boil.Infer())
		if err != nil {
			logger.Logger.Warn("delete publish msg comment from db failed", zap.Any("commentId", commentId), zap.Error(err))
			return
		}
		return
	}()

	// 删除es动态评论
	scriptStr := "ctx._source.status = ctx._source.user.id=params.uid ? 'DELETED': 'NORMAL'"
	userScript := elastic.NewScriptInline(scriptStr).Param("uid", uid)

	_, err = esClient.Update().
		Index(common.HOMEPUBLISHMSGCOMMENT).
		Id(commentId).
		Script(userScript).
		//Doc(map[string]interface{}{"status": "DELETED"}).
		Do(context.Background())
	if err != nil {
		logger.Logger.Warn("delete publish msg comment from es failed", zap.Any("commentId", commentId), zap.Error(err))
		return
	}

	return nil
}

func (hl *HomeLogic) GetUserInfoFromEs(uid string) (esUserInfo *pb.EsUserInfo, err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("GetUserInfoFromEs recover", err)
		}
	}()
	esClient, err := es.NewElasticSearch()
	if err != nil {
		logger.Logger.Warn("New ElasticSearch failed", zap.Error(err))
		return nil, err
	}

	// 查询动态里的用户信息
	getUser, err := esClient.Get().Index(common.HOMEPUBLISHMSGUSER).Id(uid).Do(context.Background())
	if err != nil {
		logger.Logger.Warn("home publish msg user byte convert string failed", zap.Any("uid", uid), zap.Error(err))
		return nil, err
	}

	sourceUser, err := getUser.Source.MarshalJSON()
	if err != nil {
		logger.Logger.Warn("home publish msg user byte convert string failed", zap.Any("uid", uid), zap.Error(err))
		return nil, err
	}

	fmt.Println("sourceUser", string(sourceUser))

	esUserInfo = &pb.EsUserInfo{}
	if err = json.Unmarshal(sourceUser, esUserInfo); err != nil {
		logger.Logger.Warn("json Unmarshal home publish msg user failed", zap.Any("esUserInfo", string(sourceUser)), zap.Error(err))
		return nil, err
	}

	return esUserInfo, nil

}
