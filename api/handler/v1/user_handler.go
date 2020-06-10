package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
	"time"
	"university_circles/api/middlewares"
	pb "university_circles/api/pb/user"
	"university_circles/api/utils/common"
	"university_circles/api/utils/logger"

	uclient "university_circles/api/client/user"
	errcode "university_circles/api/utils/errcode/user"
)

var (
	userClient = uclient.NewUserClient()
)

func StudentRegister(c *gin.Context) {
	var user pb.UserRegisterReq

	if err := c.ShouldBindJSON(&user); err != nil {
		c.Error(errcode.ErrParam)
		return
	}

	if user.Phone == "" || user.Username == "" || user.Code == "" {
		c.Error(errcode.ErrParam)
		return
	}
	// 去除空格
	user.Username = strings.Replace(user.Username, " ", "", -1)
	user.Email = strings.Replace(user.Email, " ", "", -1)
	user.Phone = strings.Replace(user.Phone, " ", "", -1)
	user.IdCardNumber = strings.Replace(user.IdCardNumber, " ", "", -1)
	user.UserNo = strings.Replace(user.UserNo, " ", "", -1)

	resp, err := userClient.StudentRegister(c, &user)
	if err != nil {
		logger.Logger.Warn("api user register failed", zap.Error(err))
		c.Error(errcode.ErrRegisterFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrRegisterFailed)
		return
	}

	if resp.Success == 0 {
		c.JSON(http.StatusOK, gin.H{
			"err_code": 0,
			"err_msg":  "success",
		})
	} else if resp.Success == errcode.ErrRegisterPhoneExist.Code {
		c.Error(errcode.ErrRegisterPhoneExist)
	} else if resp.Success == errcode.ErrRegisterNicknameExist.Code {
		c.Error(errcode.ErrRegisterNicknameExist)
	} else if resp.Success == errcode.ErrVerifyCodeNotExist.Code {
		c.Error(errcode.ErrVerifyCodeNotExist)
	} else if resp.Success == errcode.ErrVerifyCodeCompareFailed.Code {
		c.Error(errcode.ErrVerifyCodeCompareFailed)
	} else if resp.Success == errcode.ErrRegisterStuNoExist.Code {
		c.Error(errcode.ErrRegisterStuNoExist)
	} else if resp.Success == errcode.ErrRegisterIDCardNumberExist.Code {
		c.Error(errcode.ErrRegisterIDCardNumberExist)
	} else if resp.Success == errcode.ErrRegisterUniversityNotExist.Code {
		c.Error(errcode.ErrRegisterUniversityNotExist)
	} else if resp.Success == errcode.ErrRegisterProfessionNotExist.Code {
		c.Error(errcode.ErrRegisterProfessionNotExist)
	} else if resp.Success == errcode.ErrRegisterEmailExist.Code {
		c.Error(errcode.ErrRegisterEmailExist)
	} else if resp.Success == errcode.ErrRegisterPhoneNotExist.Code {
		c.Error(errcode.ErrRegisterPhoneNotExist)
	} else if resp.Success == errcode.ErrUserNameVerifyInvalid.Code {
		c.Error(errcode.ErrUserNameVerifyInvalid)
	} else if resp.Success == errcode.ErrUserIdCardVerifyInvalid.Code {
		c.Error(errcode.ErrUserIdCardVerifyInvalid)
	} else if resp.Success == errcode.ErrUserPhoneVerifyInvalid.Code {
		c.Error(errcode.ErrUserPhoneVerifyInvalid)
	} else {
		c.Error(errcode.ErrRegisterFailed)
	}
	return

}

func UpdateStudentInfo(c *gin.Context) {
	var stu pb.UpdateStudentInfoReq

	uid := c.Param("uid")
	if uid == "" {
		c.Error(errcode.ErrParam)
		return
	}

	if err := c.ShouldBindJSON(&stu); err != nil {
		c.Error(errcode.ErrParam)
		return
	}

	// 判断参数是否正确
	if uid != stu.Id {
		c.Error(errcode.ErrParam)
		return
	}

	resp, err := userClient.UpdateStudentInfo(c, &stu)
	if err != nil {
		logger.Logger.Warn("user update failed", zap.Error(err))
		c.Error(errcode.ErrDbQuery)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrModifyUserInfoFailed)
		return
	}

	if resp.Success == errcode.ErrRegisterNicknameExist.Code {
		c.Error(errcode.ErrRegisterNicknameExist)
		return
	}

	if resp.Success == errcode.ErrUserNotExist.Code {
		c.Error(errcode.ErrUserNotExist)
		return
	}

	if resp.Success == 0 {
		c.JSON(http.StatusOK, gin.H{
			"err_code": 0,
			"err_msg":  "success",
		})
		return
	} else {
		c.Error(errcode.ErrModifyUserInfoFailed)
	}
	return

}

func UpdateUserPhone(c *gin.Context) {
	var u pb.UpdateUserPhoneReq

	uid := c.Param("uid")
	if uid == "" {
		c.Error(errcode.ErrParam)
		return
	}

	if err := c.ShouldBindJSON(&u); err != nil {
		c.Error(errcode.ErrParam)
		return
	}

	// 判断参数是否正确
	if uid != u.Uid {
		c.Error(errcode.ErrParam)
		return
	}

	resp, err := userClient.UpdateUserPhone(c, &u)
	if err != nil {
		logger.Logger.Warn("user update phone failed", zap.Error(err))
		c.Error(errcode.ErrDbQuery)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrModifyUserInfoFailed)
		return
	}

	if resp.Success == errcode.ErrUserNotExist.Code {
		c.Error(errcode.ErrUserNotExist)
		return
	}

	if resp.Success == 0 {
		c.JSON(http.StatusOK, gin.H{
			"err_code": 0,
			"err_msg":  "success",
		})
		return
	} else {
		c.Error(errcode.ErrModifyUserInfoFailed)
	}
	return

}

func UpdateUserPassword(c *gin.Context) {
	var u pb.UpdateUserPasswordReq

	if err := c.ShouldBindJSON(&u); err != nil {
		c.Error(errcode.ErrParam)
		return
	}

	resp, err := userClient.UpdateUserPassword(c, &u)
	if err != nil {
		logger.Logger.Warn("user update password failed", zap.Error(err))
		c.Error(errcode.ErrDbQuery)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrModifyUserInfoFailed)
		return
	}

	if resp.Success == errcode.ErrUserNotExist.Code {
		c.Error(errcode.ErrUserNotExist)
		return
	}

	if resp.Success == errcode.ErrVerifyCodeNotExist.Code {
		c.Error(errcode.ErrVerifyCodeNotExist)
		return
	}

	if resp.Success == errcode.ErrVerifyCodeCompareFailed.Code {
		c.Error(errcode.ErrVerifyCodeCompareFailed)
		return
	}

	if resp.Success == errcode.ErrUserIsNotVerify.Code {
		c.Error(errcode.ErrUserIsNotVerify)
		return
	}

	if resp.Success == 0 {
		c.JSON(http.StatusOK, gin.H{
			"err_code": 0,
			"err_msg":  "success",
		})
		return
	} else {
		c.Error(errcode.ErrModifyUserInfoFailed)
	}
	return

}

func UpdateUserAvatar(c *gin.Context) {
	var u pb.UpdateUserAvatarReq

	uid := c.Param("uid")
	if uid == "" {
		c.Error(errcode.ErrParam)
		return
	}

	if err := c.ShouldBindJSON(&u); err != nil {
		c.Error(errcode.ErrParam)
		return
	}

	// 判断参数是否正确
	if uid != u.Uid {
		c.Error(errcode.ErrParam)
		return
	}

	resp, err := userClient.UpdateUserAvatar(c, &u)
	if err != nil {
		logger.Logger.Warn("user update avatar failed", zap.Error(err))
		c.Error(errcode.ErrDbQuery)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrModifyUserInfoFailed)
		return
	}

	if resp.Success == errcode.ErrUserNotExist.Code {
		c.Error(errcode.ErrUserNotExist)
		return
	}

	if resp.Success == 0 {
		c.JSON(http.StatusOK, gin.H{
			"err_code": 0,
			"err_msg":  "success",
		})
		return
	} else {
		c.Error(errcode.ErrModifyUserInfoFailed)
	}
	return

}

type UserLogin struct {
	Username string `json:"username", form:"username" binding:"required"`
	Password string `json:"password", form:"password" binding:"required"`
	Type     int    `json:"type", form:"type" binding:"required"`
	// 设备类型，1.移动设备(mobile)  2.桌面端(desktop)  3.web端(web)
	DeviceType  string    `json:"device_type", form:"device_type" binding:"required"`
}

func PwdLoginHandler(c *gin.Context) {
	var userLogin UserLogin
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.Error(errcode.ErrParam)
		return
	}

	if userLogin.Type != 1 && userLogin.Type != 2 {
		c.Error(errcode.ErrParam)
		return
	}

	login := &pb.PwdLoginReq{
		Username: userLogin.Username,
		Password: userLogin.Password,
	}

	resp, err := userClient.PwdLogin(c, login)
	if err != nil {
		logger.Logger.Warn("api user login failed", zap.Any("user", login), zap.Error(err))
		c.Error(errcode.ErrDbQuery)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrLoginUserNameOrPwdFailed)
		return
	}

	if resp.Success == 0 {
		token, _ := utils.SetLoginToken(resp.Uid, resp.Nickname, userLogin.DeviceType, resp.UniversityId, )
		c.Header("token", token)
		c.JSON(http.StatusOK, gin.H{
			"err_code": 0,
			"err_msg":  "success",
			"token":    token,
			"uid":      resp.Uid,
			"unid":     resp.UniversityId,
			"mid":      resp.Mid,
			"name":     resp.Nickname,
			"avatar":   resp.Avatar,
		})

		return
	} else if resp.Success == errcode.ErrLoginUserNameFailed.Code {
		c.Error(errcode.ErrLoginUserNameFailed)
		return
	} else if resp.Success == errcode.ErrLoginUserPwdFailed.Code {
		c.Error(errcode.ErrLoginUserPwdFailed)
		return
	} else if resp.Success == errcode.ErrUserIsNotVerify.Code {
		c.Error(errcode.ErrUserIsNotVerify)
		return
	} else {
		c.Error(errcode.ErrLoginUserNameOrPwdFailed)
		return
	}

	return
}

func GetVerifyCode(c *gin.Context) {
	phone := c.Param("phone")

	if phone == "" {
		c.Error(errcode.ErrParam)
		return
	}

	verifyCode := &pb.VerifyCodeReq{
		Phone: phone,
	}
	resp, err := userClient.GetVerifyCode(c, verifyCode)
	if err != nil {
		logger.Logger.Warn("api get verify code failed", zap.Any("verifyCode", verifyCode), zap.Error(err))
		c.Error(errcode.ErrVerifyCodeFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrSendVerifyCodeFailed)
		return
	}

	if resp.Status == 0 {
		c.JSON(http.StatusOK, gin.H{
			"err_code": 0,
			"err_msg":  "success",
		})
	} else if resp.Status == errcode.ErrBusinessDayLimitControlFailed.Code {
		c.Error(errcode.ErrBusinessDayLimitControlFailed)
		return
	} else if resp.Status == errcode.ErrBusinessMinuteLimitControlFailed.Code {
		c.Error(errcode.ErrBusinessMinuteLimitControlFailed)
		return
	} else {
		c.Error(errcode.ErrSendVerifyCodeFailed)
		return
	}

	return
}

func CheckVerifyCode(c *gin.Context) {
	phone := c.Param("phone")

	if phone == "" {
		c.Error(errcode.ErrParam)
		return
	}

	code := c.Query("code")

	verifyCode := &pb.VerifyCodeRegReq{
		Phone: phone,
		Code:  code,
	}

	resp, err := userClient.CheckVerifyCode(c, verifyCode)
	if err != nil {
		logger.Logger.Warn("api get verify code failed", zap.Any("verifyCode", verifyCode), zap.Error(err))
		c.Error(errcode.ErrVerifyCodeFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrVerifyCodeCompareFailed)
		return
	}

	if resp.Success == 0 {
		c.JSON(http.StatusOK, gin.H{
			"err_code": 0,
			"err_msg":  "success",
		})
	} else {
		c.Error(errcode.ErrVerifyCodeCompareFailed)
		return
	}

	return
}

type UserCodeLogin struct {
	Phone string `json:"phone", form:"phone" binding:"required"`
	Code  string `json:"code", form:"code" binding:"required"`
	Type  int    `json:"type", form:"type" binding:"required"`
	// 设备类型，1.移动设备(mobile)  2.桌面端(desktop)  3.web端(web)
	DeviceType  string    `json:"device_type", form:"device_type" binding:"required"`
}

func VerifyCodeLogin(c *gin.Context) {
	var userCodeLogin UserCodeLogin
	if err := c.ShouldBindJSON(&userCodeLogin); err != nil {
		c.Error(errcode.ErrParam)
		return
	}

	if userCodeLogin.Type != 1 && userCodeLogin.Type != 2 {
		c.Error(errcode.ErrParam)
		return
	}

	verifyCodeLogin := &pb.VerifyCodeLoginReq{
		Phone: userCodeLogin.Phone,
		Code:  userCodeLogin.Code,
	}

	resp, err := userClient.VerifyCodeLogin(c, verifyCodeLogin)
	if err != nil {
		logger.Logger.Warn("api Verify Code Login failed", zap.Any("verifyCodeLogin", verifyCodeLogin), zap.Error(err))
		c.Error(errcode.ErrLoginFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrUserNotExist)
		return
	}

	if resp.Success == errcode.ErrUserNotExist.Code {
		c.Error(errcode.ErrUserNotExist)
		return
	}

	if resp.Success == errcode.ErrVerifyCodeNotExist.Code {
		c.Error(errcode.ErrVerifyCodeNotExist)
		return
	}

	if resp.Success == errcode.ErrVerifyCodeCompareFailed.Code {
		c.Error(errcode.ErrVerifyCodeCompareFailed)
		return
	}

	if resp.Success == errcode.ErrUserIsNotVerify.Code {
		c.Error(errcode.ErrUserIsNotVerify)
		return
	}

	if resp.Success == errcode.ErrLoginFailed.Code {
		c.Error(errcode.ErrLoginFailed)
		return
	}

	if resp.Success == 0 {
		token, _ := utils.SetLoginToken(resp.Uid, resp.Nickname, userCodeLogin.DeviceType, resp.UniversityId)
		c.Header("token", token)
		c.JSON(http.StatusOK, gin.H{
			"err_code": 0,
			"err_msg":  "success",
			"token":    token,
			"uid":      resp.Uid,
			"unid":     resp.UniversityId,
			"mid":      resp.Mid,
			"name":     resp.Nickname,
			"avatar":   resp.Avatar,
		})

		return
	}

	return
}

func GetStudentInfoById(c *gin.Context) {
	uid := c.Param("uid")
	if uid == "" {
		c.Error(errcode.ErrParam)
		return
	}

	logoutReq := &pb.GetStudentByIdReq{
		Id: uid,
	}
	resp, err := userClient.GetStudentInfoById(c, logoutReq)
	if err != nil {
		logger.Logger.Warn("api Get Student Info by id failed", zap.Any("uid", uid), zap.Error(err))
		c.Error(errcode.ErrLogoutFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrUserNotExist)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"data":     resp,
	})

}

func SaveUserFollow(c *gin.Context) {
	uid := c.Param("uid")
	if uid == "" {
		c.Error(errcode.ErrParam)
		return
	}

	var userFollowReq *pb.UserFollowOperateReq
	if err := c.ShouldBindJSON(&userFollowReq); err != nil {
		fmt.Println(err)
		logger.Logger.Warn("bind home publish msg failed")
		c.Error(errcode.ErrParam)
		return
	}

	// 判断参数是否正确
	if uid != userFollowReq.Uid {
		c.Error(errcode.ErrParam)
		return
	}

	resp, err := userClient.SaveUserFollow(c, userFollowReq)
	if err != nil {
		logger.Logger.Warn("save user follow failed", zap.Any("userFollowReq", userFollowReq), zap.Error(err))
		c.Error(errcode.ErrLogoutFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrUserFollowFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
	})

}

func CancelUserFollow(c *gin.Context) {
	uid := c.Param("uid")
	if uid == "" {
		c.Error(errcode.ErrParam)
		return
	}

	var userFollowReq *pb.UserFollowOperateReq
	if err := c.ShouldBindJSON(&userFollowReq); err != nil {
		fmt.Println(err)
		logger.Logger.Warn("bind home publish msg failed")
		c.Error(errcode.ErrParam)
		return
	}

	// 判断参数是否正确
	if uid != userFollowReq.Uid {
		c.Error(errcode.ErrParam)
		return
	}

	resp, err := userClient.CancelUserFollow(c, userFollowReq)
	if err != nil {
		logger.Logger.Warn("cancel user follow failed", zap.Any("userFollowReq", userFollowReq), zap.Error(err))
		c.Error(errcode.ErrLogoutFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrUserFollowFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"data":     resp,
	})

}

func GetStudentInfoByUsername(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.Error(errcode.ErrParam)
		return
	}

	logoutReq := &pb.GetStudentByUsernameReq{
		Username: username,
	}
	resp, err := userClient.GetStudentInfoByUsername(c, logoutReq)
	if err != nil {
		logger.Logger.Warn("api Get Student Info by username failed", zap.Any("username", username), zap.Error(err))
		c.Error(errcode.ErrLogoutFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrUserNotExist)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"data":     resp,
	})

}

func Logout(c *gin.Context) {
	uid := c.Param("uid")

	if uid == "" {
		c.Error(errcode.ErrParam)
		return
	}

	logoutReq := &pb.LogoutReq{
		Id: uid,
	}
	resp, err := userClient.Logout(c, logoutReq)
	if err != nil {
		logger.Logger.Warn("api logout failed", zap.Any("uid", uid), zap.Error(err))
		c.Error(errcode.ErrLogoutFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrUserNotExist)
		return
	}

	if resp.Success == errcode.ErrUserNotExist.Code {
		c.Error(errcode.ErrUserNotExist)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
	})
}

func RefreshLoginToken(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)
	uid := c.Param("uid")

	if uid == "" {
		c.Error(errcode.ErrParam)
		return
	}

	token, err := utils.RefreshLoginToken(uid, loginSession.ScreenName, loginSession.UniversityId)
	if err != nil {
		c.Error(errcode.ErrLoginSessionInvalid)
		return
	}

	c.Header("token", token)
	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"data":     token,
	})
}

func GetUniversityList(c *gin.Context) {
	getUniversityListReq := &pb.GetUniversityListReq{}
	resp, err := userClient.GetUniversityList(c, getUniversityListReq)
	if err != nil {
		logger.Logger.Warn("api getUniversityList failed", zap.Error(err))
		c.Error(errcode.ErrGetUniversityListFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrGetUniversityListFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"data":     resp.GetUniversityList,
	})
}

func GetUniversity(c *gin.Context) {
	university := c.Query("university")

	if university == "" {
		c.Error(errcode.ErrParam)
		return
	}
	getUniversityReq := &pb.GetUniversityReq{
		University: university,
	}
	resp, err := userClient.GetUniversity(c, getUniversityReq)
	if err != nil {
		logger.Logger.Warn("api getUniversityList failed", zap.Error(err))
		c.Error(errcode.ErrGetUniversityListFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrGetUniversityListFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"data":     resp.GetUniversityList,
	})
}

func GetCollegeList(c *gin.Context) {
	unid := c.Param("university_id")

	if unid == "" {
		c.Error(errcode.ErrParam)
		return
	}

	getCollegeListReq := &pb.GetCollegeListReq{
		UniversityId: unid,
	}
	resp, err := userClient.GetCollegeList(c, getCollegeListReq)
	if err != nil {
		logger.Logger.Warn("api getCollegeList failed", zap.Error(err))
		c.Error(errcode.ErrGetCollegeListFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrGetCollegeListFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"data":     resp.GetCollegeList,
	})
}

func GetProfessionList(c *gin.Context) {
	unid := c.Param("university_id")
	collegeId := c.Param("college_id")

	if unid == "" {
		c.Error(errcode.ErrParam)
		return
	}

	getScienceListReq := &pb.GetProfessionListReq{
		UniversityId: unid,
		CollegeId:    collegeId,
	}
	resp, err := userClient.GetProfessionList(c, getScienceListReq)
	if err != nil {
		logger.Logger.Warn("api getScienceList failed", zap.Error(err))
		c.Error(errcode.ErrGetScienceListFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrGetScienceListFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"data":     resp,
	})
}

func GetClassList(c *gin.Context) {
	unid := c.Param("university_id")
	collegeId := c.Param("college_id")
	professionId := c.Param("profession_id")

	if unid == "" {
		c.Error(errcode.ErrParam)
		return
	}

	getClassListReq := &pb.GetClassListReq{
		UniversityId: unid,
		CollegeId:    collegeId,
		ProfessionId: professionId,
	}
	resp, err := userClient.GetClassList(c, getClassListReq)
	if err != nil {
		logger.Logger.Warn("api getClassList failed", zap.Error(err))
		c.Error(errcode.ErrGetClassListFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrGetClassListFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"data":     resp.GetClassList,
	})
}

func UploadAvatar(c *gin.Context) {
	fileType, err := strconv.Atoi(c.PostForm("type"))
	if fileType == 0 || err != nil {
		c.Error(errcode.ErrParam)
		return
	}

	header, err := c.FormFile("file")
	if err != nil {
		logger.Logger.Warn("Get form file failed", zap.String("filename", header.Filename), zap.Error(err))
		c.Error(errcode.ErrParam)
		return
	}

	dst := header.Filename
	if err := c.SaveUploadedFile(header, dst); err != nil {
		logger.Logger.Warn("save uploaded avatar file failed", zap.String("filename", header.Filename), zap.Error(err))
		c.Error(errcode.ErrUploadFileToOssFailed)
		return
	}

	imageId := utils.KRand(16, utils.KC_RAND_KIND_ALL) + strconv.FormatInt(time.Now().Unix(), 10)
	//imageId, _ := strconv.Atoi(image)

	var filePath string
	if fileType == USERAVATARFILE {
		filePath = fmt.Sprintf(utils.UPLOADFILEPATHPREFIX, USERAVATARFILE, imageId)
	} else if fileType == USERREGISTERFILE {
		filePath = fmt.Sprintf(utils.UPLOADFILEPATHPREFIX, USERREGISTERFILE, imageId)
	} else {
		c.Error(errcode.ErrUploadFileToOssFailed)
		return
	}

	yunFilePath := filePath + imageId
	err = utils.UploadFileToOSS(dst, yunFilePath)
	if err != nil {
		logger.Logger.Warn("upload file to oss failed", zap.String("filename", header.Filename), zap.Error(err))
		c.Error(errcode.ErrUploadFileToOssFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"image_id": imageId,
	})
}

func TeacherRegister(c *gin.Context) {
	var user pb.UserRegisterReq

	if err := c.ShouldBindJSON(&user); err != nil {
		c.Error(errcode.ErrParam)
		return
	}

	if user.Phone == "" || user.Username == "" || user.Code == "" {
		c.Error(errcode.ErrParam)
		return
	}

	// 去除空格
	user.ScreenName = strings.Replace(user.ScreenName, " ", "", -1)
	user.Email = strings.Replace(user.Email, " ", "", -1)
	user.Phone = strings.Replace(user.Phone, " ", "", -1)
	user.IdCardNumber = strings.Replace(user.IdCardNumber, " ", "", -1)
	user.UserNo = strings.Replace(user.UserNo, " ", "", -1)

	resp, err := userClient.TeacherRegister(c, &user)
	if err != nil {
		logger.Logger.Warn("api user register failed", zap.Error(err))
		c.Error(errcode.ErrRegisterFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrRegisterFailed)
		return
	}

	if resp.Success == 0 {
		c.JSON(http.StatusOK, gin.H{
			"err_code": 0,
			"err_msg":  "success",
		})
	} else if resp.Success == errcode.ErrRegisterPhoneExist.Code {
		c.Error(errcode.ErrRegisterPhoneExist)
	} else if resp.Success == errcode.ErrRegisterNicknameExist.Code {
		c.Error(errcode.ErrRegisterNicknameExist)
	} else if resp.Success == errcode.ErrVerifyCodeNotExist.Code {
		c.Error(errcode.ErrVerifyCodeNotExist)
	} else if resp.Success == errcode.ErrVerifyCodeCompareFailed.Code {
		c.Error(errcode.ErrVerifyCodeCompareFailed)
	} else if resp.Success == errcode.ErrRegisterStuNoExist.Code {
		c.Error(errcode.ErrRegisterStuNoExist)
	} else if resp.Success == errcode.ErrRegisterIDCardNumberExist.Code {
		c.Error(errcode.ErrRegisterIDCardNumberExist)
	} else if resp.Success == errcode.ErrRegisterUniversityNotExist.Code {
		c.Error(errcode.ErrRegisterUniversityNotExist)
	} else if resp.Success == errcode.ErrRegisterProfessionNotExist.Code {
		c.Error(errcode.ErrRegisterProfessionNotExist)
	} else if resp.Success == errcode.ErrRegisterEmailExist.Code {
		c.Error(errcode.ErrRegisterEmailExist)
	} else if resp.Success == errcode.ErrRegisterPhoneNotExist.Code {
		c.Error(errcode.ErrRegisterPhoneNotExist)
	} else if resp.Success == errcode.ErrUserNameVerifyInvalid.Code {
		c.Error(errcode.ErrUserNameVerifyInvalid)
	} else if resp.Success == errcode.ErrUserIdCardVerifyInvalid.Code {
		c.Error(errcode.ErrUserIdCardVerifyInvalid)
	} else if resp.Success == errcode.ErrUserPhoneVerifyInvalid.Code {
		c.Error(errcode.ErrUserPhoneVerifyInvalid)
	} else {
		c.Error(errcode.ErrRegisterFailed)
	}
	return
}

func UpdateTeacherInfo(c *gin.Context) {
	var teacher pb.UpdateTeacherInfoReq

	uid := c.Param("uid")
	if uid == "" {
		c.Error(errcode.ErrParam)
		return
	}

	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.Error(errcode.ErrParam)
		return
	}

	// 判断参数是否正确
	if uid != teacher.Id {
		c.Error(errcode.ErrParam)
		return
	}

	resp, err := userClient.UpdateTeacherInfo(c, &teacher)
	if err != nil {
		logger.Logger.Warn("user update failed", zap.Error(err))
		c.Error(errcode.ErrDbQuery)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrModifyUserInfoFailed)
		return
	}

	if resp.Success == errcode.ErrRegisterNicknameExist.Code {
		c.Error(errcode.ErrRegisterNicknameExist)
		return
	}

	if resp.Success == errcode.ErrUserNotExist.Code {
		c.Error(errcode.ErrUserNotExist)
		return
	}

	if resp.Success == 0 {
		c.JSON(http.StatusOK, gin.H{
			"err_code": 0,
			"err_msg":  "success",
		})
		return
	} else {
		c.Error(errcode.ErrModifyUserInfoFailed)
	}
	return

}

func QueryUser(c *gin.Context) {
	var user pb.QueryUserReq

	if err := c.ShouldBindJSON(&user); err != nil {
		c.Error(errcode.ErrParam)
		return
	}

	user.ReqStr = strings.Replace(user.ReqStr, " ", "", -1)
	resp, err := userClient.QueryUser(c, &user)
	if err != nil {
		logger.Logger.Warn("user search failed", zap.Error(err))
		c.Error(errcode.ErrQueryUserFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrQueryUserFailed)
		return
	}

	data := resp.UserInfos
	if data == nil {
		data = make([]*pb.QueryUser, 0)
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"data": data,
	})
	return

}

func CheckUserIsExist(c *gin.Context) {

	phone := c.Query("phone")
	username := c.Query("username")

	// 去除空格
	username = strings.Replace(username, " ", "", -1)
	phone = strings.Replace(phone, " ", "", -1)

	checkInfo := &pb.CheckUserIsExistReq{
		Phone: phone,
		Username:  username,
	}

	resp, err := userClient.CheckUserIsExist(c, checkInfo)
	if err != nil {
		logger.Logger.Warn("api check user is exist failed", zap.Any("checkInfo", checkInfo), zap.Error(err))
		c.Error(errcode.ErrVerifyCodeFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrParam)
		return
	}

	if resp.Success == 0 {
		c.JSON(http.StatusOK, gin.H{
			"err_code": 0,
			"err_msg":  "success",
		})
	} else if resp.Success == errcode.ErrRegisterPhoneExist.Code {
		c.Error(errcode.ErrRegisterPhoneExist)
	} else if resp.Success == errcode.ErrRegisterNicknameExist.Code {
		c.Error(errcode.ErrRegisterNicknameExist)
	}

	return
}

