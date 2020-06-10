package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	iclient "university_circles/api/client/im"
	"university_circles/api/middlewares"
	pb "university_circles/api/pb/im"

	"university_circles/api/utils/common"
	errcode "university_circles/api/utils/errcode/im"
	"university_circles/api/utils/logger"
)

var (
	imClient = iclient.NewImClient()
)

type AuthUserMessage struct {
	UID        int64  `json:"uid" form:"uid"`
	PlatformId int64  `json:"platform_id" form:"platform_id"`
	DeviceId   string `json:"device_id" form:"device_id"`
	AppId      int64  `json:"app_id" form:"app_id"`
}

func IMAuthToken(c *gin.Context) {
	var authUserMessage AuthUserMessage
	if err := c.ShouldBindJSON(&authUserMessage); err != nil {
		c.Error(errcode.ErrIMAuthTokenFailed)
		return
	}

	fmt.Println("authUserMessage, err",authUserMessage)

	logger.Logger.Info("user im auth info", zap.Any("auth info ", authUserMessage))

	// 获取旧的token
	//oldToken, err := utils.GetUserIMToken(authUserMessage.UID)
	//fmt.Println("oldToken, err",oldToken, err)
	//if err != nil {
	//	logger.Logger.Warn("user hget im auth old token failed", zap.Any("auth info ", authUserMessage), zap.Error(err))
	//	c.Error(errcode.ErrIMAuthTokenFailed)
	//	return
	//}

	token := utils.GenIMToken(authUserMessage.AppId, authUserMessage.UID)

	fMap := make(map[string]interface{})
	fMap["user_id"] = authUserMessage.UID
	fMap["app_id"] = authUserMessage.AppId
	fMap["notification_on"] = int8(1)
	fMap["forbidden"] = 1

	// 保存新的token
	if err := utils.SaveIMToken(token, fMap); err != nil {
		logger.Logger.Warn("user hmset im auth token failed", zap.Any("auth info ", authUserMessage), zap.Error(err))
		c.Error(errcode.ErrIMAuthTokenFailed)
		return
	}

	if err := utils.SetUserIMToken(token, authUserMessage.UID); err != nil {
		fmt.Println("SetUserIMToken",err)
		logger.Logger.Warn("user hset im auth token failed", zap.Any("auth info ", authUserMessage), zap.Error(err))
		c.Error(errcode.ErrIMAuthTokenFailed)
		return
	}

	// 删除旧的token
	//if oldToken != "" {
	//	go utils.DelIMToken(oldToken)
	//}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"data":     token,
	})

}

func GetUserUnreadCount(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	var unreadCount *pb.GetUserUnReadCountReq
	if err := c.ShouldBindJSON(&unreadCount); err != nil {
		fmt.Println(err)
		c.Error(errcode.ErrAddFriendFailed)
		return
	}

	unreadCount.UserId = loginSession.Uid

	resp, err := imClient.GetUserUnReadCount(c, unreadCount)
	fmt.Println("GetUserUnreadCount resp, err", resp, err)
	if err != nil {
		logger.Logger.Warn("GetUserUnreadCount failed", zap.Error(err))
		c.Error(errcode.ErrAddFriendFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrAddFriendFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"data": resp.UnreadCount,
	})

}

func SendPeerMsg(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	var req *pb.SendPeerMsgReq
	var token string
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		c.Error(errcode.ErrAddFriendFailed)
		return
	}

	req.UserId = loginSession.Uid

	if req.ChatType == "addFriend" ||
		req.ChatType == "refusedFriend"{
		req.ChatType = "system"
	} else if req.ChatType == "agreedFriend" {
		req.ChatType = "system"
		friend := &pb.AddFriendReq{
			Uid: req.Uid,
			FriendUid: req.ReceiverUid,
			Appid: req.Appid,
			UserId: loginSession.Uid,
		}

		resp, err := imClient.AddFriends(c, friend)
		if err != nil {
			logger.Logger.Warn("AddFriends failed", zap.Error(err))
			c.Error(errcode.ErrAddFriendFailed)
			return
		}

		if resp == nil {
			c.Error(errcode.ErrAddFriendFailed)
			return
		}
	}

	// 获取token
	token, err = utils.GetUserIMToken(req.Uid)
	if err != nil {
		logger.Logger.Warn("user hget im auth old token failed", zap.Any("req", req), zap.Error(err))
		c.Error(errcode.ErrIMAuthTokenFailed)
		return
	}

	if token == "" {
		token = utils.GenIMToken(req.Appid, req.Uid)

		fMap := make(map[string]interface{})
		fMap["user_id"] = req.Uid
		fMap["app_id"] = req.Appid
		fMap["notification_on"] = 1
		fMap["forbidden"] = 1

		if err := utils.SaveIMToken(token, fMap); err != nil {
			fmt.Println(err)
			logger.Logger.Warn("SendAddFriendMsg gen im auth token failed", zap.Any("token info ", req), zap.Error(err))
			c.Error(errcode.ErrIMAuthTokenFailed)
			return
		}
	}

	req.Token = token

	resp, err := imClient.SendPeerMsg(c, req)
	if err != nil {
		logger.Logger.Warn("AddFriends failed", zap.Error(err))
		c.Error(errcode.ErrAddFriendFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrAddFriendFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
	})

}

func AddFriends(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	var friend *pb.AddFriendReq
	if err := c.ShouldBindJSON(&friend); err != nil {
		fmt.Println(err)
		c.Error(errcode.ErrAddFriendFailed)
		return
	}

	friend.UserId = loginSession.Uid

	logger.Logger.Info("add friend info", zap.Any("user", loginSession), zap.Any("friend", friend))

	resp, err := imClient.AddFriends(c, friend)
	fmt.Println("AddFriends resp, err", resp, err)
	if err != nil {
		logger.Logger.Warn("AddFriends failed", zap.Error(err))
		c.Error(errcode.ErrAddFriendFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrAddFriendFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
	})

}

func DelFriends(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	var friend *pb.DelFriendReq
	if err := c.ShouldBindJSON(&friend); err != nil {
		fmt.Println(err)
		c.Error(errcode.ErrDelFriendFailed)
		return
	}

	friend.UserId = loginSession.Uid

	logger.Logger.Info("del friend info", zap.Any("user", loginSession), zap.Any("friend", friend))

	resp, err := imClient.DelFriends(c, friend)
	fmt.Println("DelFriends resp, err", resp, err)
	if err != nil {
		logger.Logger.Warn("DelFriends failed", zap.Error(err))
		c.Error(errcode.ErrDelFriendFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrDelFriendFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
	})

}

func UpdateFriendRemark(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	var friend *pb.UpdateFriendRemarkReq
	if err := c.ShouldBindJSON(&friend); err != nil {
		c.Error(errcode.ErrDelFriendFailed)
		return
	}

	friend.UserId = loginSession.Uid

	logger.Logger.Info("update friend remark", zap.Any("user", loginSession), zap.Any("friend", friend))

	resp, err := imClient.UpdateFriendRemark(c, friend)
	if err != nil {
		c.Error(errcode.ErrDelFriendFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrDelFriendFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
	})

}


func AddFriendBlackList(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	var friend *pb.AddFriendBlackListReq
	if err := c.ShouldBindJSON(&friend); err != nil {
		fmt.Println(err)
		c.Error(errcode.ErrAddFriendBlackListFailed)
		return
	}

	friend.UserId = loginSession.Uid

	logger.Logger.Info("add friend black list info", zap.Any("user", loginSession), zap.Any("friend", friend))

	resp, err := imClient.AddFriendBlackList(c, friend)
	fmt.Println("AddFriendBlackList resp, err", resp, err)
	if err != nil {
		logger.Logger.Warn("AddFriendBlackList failed", zap.Error(err))
		c.Error(errcode.ErrAddFriendBlackListFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrAddFriendBlackListFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
	})

}

func DelFriendBlackList(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	var friend *pb.DelFriendBlackListReq
	if err := c.ShouldBindJSON(&friend); err != nil {
		fmt.Println(err)
		c.Error(errcode.ErrDelFriendBlackListFailed)
		return
	}

	friend.UserId = loginSession.Uid

	logger.Logger.Info("del friend black list info", zap.Any("user", loginSession), zap.Any("friend", friend))

	resp, err := imClient.DelFriendBlackList(c, friend)
	if err != nil {
		logger.Logger.Warn("DelFriendBlackList failed", zap.Error(err))
		c.Error(errcode.ErrDelFriendBlackListFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrDelFriendBlackListFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
	})

}

func GetAllFriends(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	var userId int64
	var appId int64
	var err error
	if uid := c.Query("uid"); uid != "" {
		userId, err = strconv.ParseInt(uid, 10, 64)
		if err != nil {
			c.Error(errcode.ErrParam)
			return
		}
	}

	if aid := c.Query("appid"); aid != "" {
		appId, err = strconv.ParseInt(aid, 10, 64)
		if err != nil {
			c.Error(errcode.ErrParam)
			return
		}
	}

	friendList := &pb.FriendListReq{
		Uid:   userId,
		Appid: appId,
		UserId: loginSession.Uid,
	}

	logger.Logger.Info("get friend list info", zap.Any("user", loginSession))

	resp, err := imClient.GetAllFriends(c, friendList)
	if err != nil {
		logger.Logger.Warn("GetAllFriends failed", zap.Error(err))
		c.Error(errcode.ErrGetAllFriendFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrGetAllFriendFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"data":     resp.FriendUids,
	})
}

func GetBlackList(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	var blacklist *pb.BlackListReq
	if err := c.ShouldBindJSON(&blacklist); err != nil {
		c.Error(errcode.ErrGetBlackListFailed)
		return
	}

	blacklist.UserId = loginSession.Uid

	logger.Logger.Info("get friend list info", zap.Any("user", loginSession))

	resp, err := imClient.GetBlackList(c, blacklist)
	if err != nil {
		logger.Logger.Warn("GetAllFriends failed", zap.Error(err))
		c.Error(errcode.ErrGetBlackListFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrGetBlackListFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"data":     resp.BlacklistUids,
	})

}

func CreateGroup(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	var group *pb.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		fmt.Println(err)
		c.Error(errcode.ErrCreteGroupFailed)
		return
	}

	group.UserId = loginSession.Uid

	logger.Logger.Info("create group", zap.Any("user", loginSession))

	resp, err := imClient.CreateGroup(c, group)
	fmt.Println("CreateGroup resp, err", resp, err)
	if err != nil {
		logger.Logger.Warn("CreateGroup failed", zap.Error(err))
		c.Error(errcode.ErrCreteGroupFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrGetBlackListFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"group_id": resp.GroupId,
	})

}

func DelGroup(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	var group *pb.DelGroupReq
	if err := c.ShouldBindJSON(&group); err != nil {
		fmt.Println(err)
		c.Error(errcode.ErrDelGroupFailed)
		return
	}

	group.UserId = loginSession.Uid
	logger.Logger.Info("del group info", zap.Any("user", loginSession))

	resp, err := imClient.DelGroup(c, group)
	fmt.Println("DelGroup resp, err", resp, err)
	if err != nil {
		logger.Logger.Warn("DelGroup failed", zap.Error(err))
		c.Error(errcode.ErrDelGroupFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrDelGroupFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
	})

}

func GetGroupMemberList(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	var groupMember *pb.GroupMemberListReq
	if err := c.ShouldBindJSON(&groupMember); err != nil {
		fmt.Println(err)
		c.Error(errcode.ErrGetGroupMemberFailed)
		return
	}
	groupMember.UserId = loginSession.Uid

	logger.Logger.Info("del group info", zap.Any("user", loginSession))

	resp, err := imClient.GetGroupMemberList(c, groupMember)
	fmt.Println("DelGroup resp, err", resp, err)
	if err != nil {
		logger.Logger.Warn("DelGroup failed", zap.Error(err))
		c.Error(errcode.ErrGetGroupMemberFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrGetGroupMemberFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"data":     gin.H{
			"list": resp.Users,
			"num": resp.Num,
		},
	})

}

func GetGroupList(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	var groupList *pb.GroupListReq
	if err := c.ShouldBindJSON(&groupList); err != nil {
		fmt.Println(err)
		c.Error(errcode.ErrGroupListFailed)
		return
	}

	groupList.UserId = loginSession.Uid

	logger.Logger.Info("get group list info", zap.Any("user", loginSession))

	resp, err := imClient.GetGroupList(c, groupList)
	fmt.Println("GetGroupList resp, err", resp, err)
	if err != nil {
		logger.Logger.Warn("DelGroup failed", zap.Error(err))
		c.Error(errcode.ErrGroupListFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrGroupListFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"data":     resp.Groups,
	})

}

func JoinGroup(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	var joinGroup *pb.JoinGroupReq
	if err := c.ShouldBindJSON(&joinGroup); err != nil {
		fmt.Println(err)
		c.Error(errcode.ErrJoinGroupFailed)
		return
	}

	joinGroup.UserId = loginSession.Uid

	logger.Logger.Info("join group info", zap.Any("user", loginSession), zap.Any("req", joinGroup))

	resp, err := imClient.JoinGroup(c, joinGroup)
	if err != nil {
		logger.Logger.Warn("JoinGroup failed", zap.Error(err))
		c.Error(errcode.ErrJoinGroupFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrJoinGroupFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
	})

}

func DelGroupMember(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	var delGroupMember *pb.DelGroupMemberReq
	if err := c.ShouldBindJSON(&delGroupMember); err != nil {
		fmt.Println(err)
		c.Error(errcode.ErrJoinGroupFailed)
		return
	}

	delGroupMember.UserId = loginSession.Uid

	logger.Logger.Info("del group member info", zap.Any("user", loginSession), zap.Any("req", delGroupMember))

	resp, err := imClient.DelGroupMember(c, delGroupMember)
	if err != nil {
		logger.Logger.Warn("DelGroupMember failed", zap.Error(err))
		c.Error(errcode.ErrJoinGroupFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrJoinGroupFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
	})

}

func UpdateGroupAvatar(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	var group *pb.UpdateGroupAvatarReq
	if err := c.ShouldBindJSON(&group); err != nil {
		c.Error(errcode.ErrParam)
		return
	}

	logger.Logger.Info("update group avatar", zap.Any("user", loginSession), zap.Any("req", group))

	group.UserId = loginSession.Uid

	resp, err := imClient.UpdateGroupAvatar(c, group)
	if err != nil {
		logger.Logger.Warn("UpdateGroupAvatar failed", zap.Error(err))
		c.Error(errcode.ErrUpdateGroupAvatarFailed)
		return
	}


	if resp == nil {
		c.Error(errcode.ErrUpdateGroupAvatarFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
	})

}

func UpdateGroupNotice(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	var group *pb.UpdateGroupNoticeReq
	if err := c.ShouldBindJSON(&group); err != nil {
		c.Error(errcode.ErrParam)
		return
	}

	group.UserId = loginSession.Uid

	logger.Logger.Info("update group avatar", zap.Any("user", loginSession), zap.Any("req", group))

	resp, err := imClient.UpdateGroupNotice(c, group)
	if err != nil {
		logger.Logger.Warn("UpdateGroupNotice failed", zap.Error(err))
		c.Error(errcode.ErrUpdateGroupNoticeFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrUpdateGroupNoticeFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
	})

}

func UpdateGroupName(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	var group *pb.UpdateGroupNameReq
	if err := c.ShouldBindJSON(&group); err != nil {
		c.Error(errcode.ErrParam)
		return
	}

	group.UserId = loginSession.Uid

	logger.Logger.Info("update group avatar", zap.Any("user", loginSession), zap.Any("req", group))

	resp, err := imClient.UpdateGroupName(c, group)
	if err != nil {
		logger.Logger.Warn("UpdateGroupNotice failed", zap.Error(err))
		c.Error(errcode.ErrUpdateGroupNameFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrUpdateGroupNameFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
	})

}

func UpdateGroupJoinAuth(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	var group *pb.UpdateGroupJoinAuthReq
	if err := c.ShouldBindJSON(&group); err != nil {
		c.Error(errcode.ErrParam)
		return
	}

	group.UserId = loginSession.Uid

	logger.Logger.Info("update group join auth", zap.Any("user", loginSession), zap.Any("req", group))

	resp, err := imClient.UpdateGroupJoinAuth(c, group)
	if err != nil {
		logger.Logger.Warn("UpdateGroupJoinAuth failed", zap.Error(err))
		c.Error(errcode.ErrUpdateGroupJoinAuthFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrUpdateGroupJoinAuthFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
	})

}









