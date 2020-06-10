package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"university_circles/api/middlewares"
)

func ClientEngine() *gin.Engine {
	router := gin.Default()

	gin.SetMode(gin.DebugMode)

	///**  global interceptor */
	// 全局中间件
	router.Use(middlewares.RequestLogger())
	router.Use(middlewares.ResponseHandler())
	router.Use(middlewares.Cors())

	api := router.Group("/api/v1")

	/**  user handler */
	userGroup := api.Group("/user")
	userGroup.POST("/student/register", StudentRegister)
	userGroup.POST("/teacher/register", TeacherRegister)
	//userGroup.POST("/student/add", AddStudentInfo)
	userGroup.POST("/login/pwd", PwdLoginHandler)
	userGroup.POST("/login/phone", VerifyCodeLogin)
	userGroup.GET("/getCode/:phone", GetVerifyCode)
	userGroup.GET("/checkCode/:phone", CheckVerifyCode)
	userGroup.POST("/upload/avatar", UploadAvatar)
	userGroup.GET("/check/reg", CheckUserIsExist)
	userGroup.Use(middlewares.LoginSessionAuth())
	userGroup.GET("/logout/:uid", Logout)
	userGroup.GET("/refresh/token/:uid", RefreshLoginToken)
	userGroup.GET("/student/detail/:uid", GetStudentInfoById)
	userGroup.GET("/student/username/:username", GetStudentInfoByUsername)
	userGroup.POST("/student/update/info/:uid", UpdateStudentInfo)
	userGroup.POST("/teacher/update/info/:uid", UpdateTeacherInfo)
	userGroup.POST("/query", QueryUser)

	userGroup.GET("/university/list/name", GetUniversity)
	userGroup.GET("/university/list", GetUniversityList)
	userGroup.GET("/college/list/:university_id", GetCollegeList)
	userGroup.GET("/science/list/:university_id/:college_id", GetProfessionList)
	userGroup.GET("/class/list/:university_id/:college_id/:profession_id", GetClassList)

	userGroup.POST("/update/phone/:uid", UpdateUserPhone)
	userGroup.POST("/update/password/:uid", UpdateUserPassword)
	userGroup.POST("/update/avatar/:uid", UpdateUserAvatar)
	// 关注
	userGroup.POST("/follow/:uid", SaveUserFollow)
	// 取消关注
	userGroup.POST("/cancel/follow/:uid", CancelUserFollow)

	/**  home handler */
	homeGroup := api.Group("/home")
	homeGroup.Use(middlewares.LoginSessionAuth())
	homeGroup.POST("/publish/msg/release/:uid", SavePublishMsg)
	homeGroup.POST("/publish/msg/delete/:mid", DelPublishMsg)
	homeGroup.GET("/publish/msg/list", GetHomePublishMsgList)
	homeGroup.GET("/publish/msg/list/user/:uid", GetUserPublishMsgList)
	homeGroup.GET("/publish/msg/detail/:mid", GetPublishMsgDetail)

	homeGroup.POST("/publish/msg/operate", SaveUserOperateMsg)
	homeGroup.POST("/publish/msg/operate/record/list")

	homeGroup.POST("/publish/msg/comment", SavePublishMsgComment)
	homeGroup.POST("/publish/msg/comment/delete/:uid", DelPublishMsgComment)
	homeGroup.GET("/publish/msg/comment/list", GetPublishMsgCommentList)

	// 通用接口
	commonGroup := api.Group("/common")
	commonGroup.Use(middlewares.LoginSessionAuth())
	// 上传图片
	commonGroup.POST("/upload/file", UploadFileToOSS)
	commonGroup.POST("/report", UserReport)

	// IM
	IMGroup := api.Group("/im")
	IMGroup.POST("/auth/token", IMAuthToken)
	IMGroup.Use(middlewares.LoginSessionAuth())
	IMGroup.POST("/user/unread/count", GetUserUnreadCount)
	IMGroup.POST("/peer/msg", SendPeerMsg)
	//IMGroup.POST("/friend/add", AddFriends)
	IMGroup.POST("/friend/del", DelFriends)
	IMGroup.GET("/friend/list", GetAllFriends)
	IMGroup.POST("/friend/update/remark", UpdateFriendRemark)

	IMGroup.POST("/blacklist/add", AddFriendBlackList)
	IMGroup.POST("/blacklist/del", DelFriendBlackList)
	IMGroup.GET("/blacklist/list", GetBlackList)

	IMGroup.POST("/group/add", CreateGroup)
	IMGroup.POST("/group/join", JoinGroup)
	IMGroup.POST("/group/del", DelGroup)
	IMGroup.GET("/group/list", GetGroupList)
	IMGroup.GET("/group/member/list", GetGroupMemberList)
	IMGroup.GET("/group/member/del", DelGroupMember)
	//IMGroup.POST("/group/update/avatar", UpdateGroupAvatar)
	IMGroup.POST("/group/update/notice", UpdateGroupNotice)
	IMGroup.POST("/group/update/name", UpdateGroupName)
	//IMGroup.POST("/group/update/auth", UpdateGroupJoinAuth)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  "Api route not found",
		})
	})

	return router
}
