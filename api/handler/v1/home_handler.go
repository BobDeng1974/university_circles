package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	hclient "university_circles/api/client/home"
	"university_circles/api/middlewares"
	errcode "university_circles/api/utils/errcode/home"
	"university_circles/api/utils/logger"

	pb "university_circles/api/pb/home"
)

var (
	homeClient = hclient.NewHomeClient()
)

func SavePublishMsg(c *gin.Context) {
	uid := c.Param("uid")
	if uid == "" {
		c.Error(errcode.ErrParam)
		return
	}

	var reqHomePublishMsg *pb.PublishMsg
	if err := c.ShouldBindJSON(&reqHomePublishMsg); err != nil {
		fmt.Println(err)
		logger.Logger.Warn("bind home publish msg failed")
		c.Error(errcode.ErrParam)
		return
	}

	// 判断参数是否正确
	if uid != reqHomePublishMsg.Uid {
		c.Error(errcode.ErrParam)
		return
	}

	resp, err := homeClient.SavePublishMsg(c, reqHomePublishMsg)
	fmt.Println("resp, err", resp, err)
	if err != nil {
		logger.Logger.Warn("api getClassList failed", zap.Error(err))
		c.Error(errcode.ErrPublishMsgFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrPublishMsgFailed)
		return
	}

	if resp.Success == -1 {
		c.Error(errcode.ErrPublishMsgFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
	})
}

func GetHomePublishMsgList(c *gin.Context) {
	var msgType int64
	var orderId int64
	var loadMore bool
	var uid string
	var universityId int64
	var err error

	if m := c.Query("msg_type"); m != "" {
		msgType, err = strconv.ParseInt(m, 10, 64)
		if err != nil {
			c.Error(errcode.ErrParam)
			return
		}
	}
	if o := c.Query("order_id"); o != "" {
		orderId, err = strconv.ParseInt(o, 10, 64)
		if err != nil {
			c.Error(errcode.ErrParam)
			return
		}
	}
	if l := c.Query("load_more"); l != "" {
		loadMore, err = strconv.ParseBool(l)
		if err != nil {
			c.Error(errcode.ErrParam)
			return
		}
	}

	uid = c.Query("uid")

	if u := c.Query("university_id"); u != "" {
		universityId, err = strconv.ParseInt(u, 10, 64)
		if err != nil {
			c.Error(errcode.ErrParam)
			return
		}
	}

	loginSession := c.MustGet("session").(middlewares.LoginSession)
	reqHomePublishMsgList := &pb.AllMsgListRequest{
		MsgType:      msgType,
		OrderId:      orderId,
		LoadMore:     loadMore,
		Uid:          uid,
		UniversityId: universityId,
	}

	fmt.Println("reqHomePublishMsgList===", reqHomePublishMsgList)
	fmt.Println(loginSession)
	if reqHomePublishMsgList.MsgType != 2 {
		reqHomePublishMsgList.Uid = loginSession.Uid
		reqHomePublishMsgList.UniversityId = loginSession.UniversityId
	}

	resp, err := homeClient.GetHomeMsgList(c, reqHomePublishMsgList)
	//fmt.Println("resp, err", resp, err)
	if err != nil {
		logger.Logger.Warn("api getClassList failed", zap.Error(err))
		c.Error(errcode.ErrGetPublishMsgListFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrGetPublishMsgListFailed)
		return
	}

	firstOrderId := 0
	lastOrderId := 0
	count := len(resp.PublishMsgResponseList)
	isLastPage := true
	toastMessage := "已没有新内容了哦～～"
	if count > 0 {
		toastMessage = "为你加载一组新内容"
		isLastPage = false
		firstOrderId = int(resp.PublishMsgResponseList[0].Item.OrderId)
		lastOrderId = int(resp.PublishMsgResponseList[count-1].Item.OrderId)
	}

	data := make([]*pb.PublishMsgListResponse, 0)
	if resp.PublishMsgResponseList != nil {
		data = resp.PublishMsgResponseList
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code":     0,
		"err_msg":      "success",
		"data":         data,
		"toastMessage": toastMessage,
		"loadMoreKey": gin.H{
			"first":      firstOrderId,
			"last":       lastOrderId,
			"isLastPage": isLastPage,
		},
	})
}

func GetUserPublishMsgList(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)
	var reqUserPublishMsgList *pb.UserMsgListRequest

	uid := c.Param("uid")
	if uid == "" {
		c.Error(errcode.ErrParam)
		return
	}
	fmt.Println("uid", uid, c.Params)
	if err := c.ShouldBindJSON(&reqUserPublishMsgList); err != nil {
		fmt.Println(err)
		logger.Logger.Warn("bind home publish msg list failed")
		c.Error(errcode.ErrParam)
		return
	}

	// 判断参数是否正确
	if uid != reqUserPublishMsgList.Uid {
		c.Error(errcode.ErrParam)
		return
	}

	if reqUserPublishMsgList.MsgType != 2 {
		reqUserPublishMsgList.UniversityId = loginSession.UniversityId
	}

	resp, err := homeClient.GetUserMsgList(c, reqUserPublishMsgList)
	fmt.Println("resp, err", resp, err)
	if err != nil {
		logger.Logger.Warn("api getClassList failed", zap.Error(err))
		c.Error(errcode.ErrGetPublishMsgListFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrGetPublishMsgListFailed)
		return
	}

	firstOrderId := 0
	lastOrderId := 0
	count := len(resp.PublishMsgResponseList)
	isLastPage := true
	if count > 0 {
		isLastPage = false
		firstOrderId = int(resp.PublishMsgResponseList[0].OrderId)
		lastOrderId = int(resp.PublishMsgResponseList[count-1].OrderId)
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code":     0,
		"err_msg":      "success",
		"data":         resp.PublishMsgResponseList,
		"toastMessage": "为你加载一组新内容",
		"loadMoreKey": gin.H{
			"first":      firstOrderId,
			"last":       lastOrderId,
			"isLastPage": isLastPage,
		},
	})
}

func GetPublishMsgDetail(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)
	msgReq := &pb.OneMsgRequest{}

	mid := c.Param("mid")
	if mid == "" {
		c.Error(errcode.ErrParam)
		return
	}

	msgReq.Id = mid
	msgReq.Uid = loginSession.Uid
	msgReq.UniversityId = loginSession.UniversityId

	fmt.Println("msgReq", msgReq)

	resp, err := homeClient.GetPublishMsgDetail(c, msgReq)
	fmt.Println("resp, err", resp, err)
	if err != nil {
		logger.Logger.Warn("api GetPublishMsgDetail failed", zap.Error(err))
		c.Error(errcode.ErrGetPublishMsgDetailFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrGetPublishMsgDetailFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"data":     resp,
	})
}

func DelPublishMsg(c *gin.Context) {
	msgId := c.Param("mid")

	loginSession := c.MustGet("session").(middlewares.LoginSession)
	delMsgRequest := &pb.DeleteMsgRequest{
		Id:  msgId,
		Uid: loginSession.Uid,
	}

	resp, err := homeClient.DeletePublishMsg(c, delMsgRequest)
	fmt.Println("resp, err", resp, err)
	if err != nil {
		logger.Logger.Warn("del publish msg failed", zap.Error(err))
		c.Error(errcode.ErrDelPublishMsgFailed)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"data":     "",
	})
}

func SavePublishMsgComment(c *gin.Context) {
	var reqMsgComment *pb.PublishMsgComment
	if err := c.ShouldBindJSON(&reqMsgComment); err != nil {
		fmt.Println(err)
		logger.Logger.Warn("bind home publish msg comment failed")
		c.Error(errcode.ErrParam)
		return
	}

	resp, err := homeClient.SaveMsgComment(c, reqMsgComment)
	fmt.Println("resp, err", resp, err)
	if err != nil {
		logger.Logger.Warn("api SaveMsgComment failed", zap.Error(err))
		c.Error(errcode.ErrPublishMsgCommentFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrPublishMsgCommentFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
	})
}

func DelPublishMsgComment(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)
	var delMsgCommentRequest *pb.DeleteMsgCommentRequest
	if err := c.ShouldBindJSON(&delMsgCommentRequest); err != nil {
		fmt.Println(err)
		logger.Logger.Warn("bind home publish msg comment failed")
		c.Error(errcode.ErrParam)
		return
	}

	delMsgCommentRequest.Uid = loginSession.Uid

	resp, err := homeClient.DeletePublishMsgComment(c, delMsgCommentRequest)
	fmt.Println("resp, err", resp, err)
	if err != nil {
		logger.Logger.Warn("api DeletePublishMsgComment failed", zap.Error(err))
		c.Error(errcode.ErrPublishMsgCommentFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrPublishMsgCommentFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
	})
}

func GetPublishMsgCommentList(c *gin.Context) {
	var reqMsgCommentList *pb.MsgCommentListRequest
	if err := c.ShouldBindJSON(&reqMsgCommentList); err != nil {
		fmt.Println(err)
		logger.Logger.Warn("bind home publish msg comment failed")
		c.Error(errcode.ErrParam)
		return
	}

	resp, err := homeClient.GetMsgCommentList(c, reqMsgCommentList)
	fmt.Println("resp, err", resp, err)
	if err != nil {
		logger.Logger.Warn("api SaveMsgCommentList failed", zap.Error(err))
		c.Error(errcode.ErrPublishMsgCommentFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrPublishMsgCommentFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"data":     resp.PublishMsgComments,
	})
}

func SaveUserOperateMsg(c *gin.Context) {
	var reqUserOperateMsgList *pb.UserOperateCountRequest
	if err := c.ShouldBindJSON(&reqUserOperateMsgList); err != nil {
		fmt.Println(err)
		logger.Logger.Warn("bind home publish msg operate failed")
		c.Error(errcode.ErrParam)
		return
	}

	resp, err := homeClient.SaveUserOperateMsgCount(c, reqUserOperateMsgList)
	fmt.Println("resp, err", resp, err)
	if err != nil {
		logger.Logger.Warn("api SaveUserOperateMsgCount failed", zap.Error(err))
		c.Error(errcode.ErrUserOperatePublishMsgFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrUserOperatePublishMsgFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
	})
}
