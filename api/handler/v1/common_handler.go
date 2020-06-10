package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
	cclient "university_circles/api/client/common"
	"university_circles/api/middlewares"
	pb "university_circles/api/pb/common"
	utils "university_circles/api/utils/common"
	errcode "university_circles/api/utils/errcode/common"

	"university_circles/api/utils/logger"
)

const (
	HOMEFILE         = 1
	USERAVATARFILE   = 2
	USERREGISTERFILE = 3
	CHATPICTURE      = 4
	CHATFILE         = 5
)

var (
	commonClient = cclient.NewCommonClient()
)

func UploadFileToOSS(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	fileType, err := strconv.Atoi(c.PostForm("type"))
	if fileType == 0 || err != nil {
		c.Error(errcode.ErrParam)
		return
	}

	header, err := c.FormFile("file")
	if err != nil {
		logger.Logger.Warn("Get form file failed", zap.Any("loginSession", loginSession), zap.String("filename", header.Filename), zap.Error(err))
		c.Error(errcode.ErrParam)
		return
	}

	dst := header.Filename
	if err := c.SaveUploadedFile(header, dst); err != nil {
		logger.Logger.Warn("save uploaded file failed", zap.Any("loginSession", loginSession), zap.String("filename", header.Filename), zap.Error(err))
		c.Error(errcode.ErrUploadFileToOssFailed)
		return
	}

	imageId := utils.KRand(16, utils.KC_RAND_KIND_ALL) + strconv.FormatInt(time.Now().Unix(), 10)
	//imageId, _ := strconv.Atoi(image)

	var filePath string
	if fileType == HOMEFILE {
		filePath = fmt.Sprintf(utils.UPLOADFILEPATHPREFIX, HOMEFILE, imageId)
	} else if fileType == USERAVATARFILE {
		filePath = fmt.Sprintf(utils.UPLOADFILEPATHPREFIX, USERAVATARFILE, imageId)
	} else if fileType == USERREGISTERFILE {
		filePath = fmt.Sprintf(utils.UPLOADFILEPATHPREFIX, USERREGISTERFILE, imageId)
	} else if fileType == CHATPICTURE {
		filePath = fmt.Sprintf(utils.UPLOADFILEPATHPREFIX, CHATPICTURE, imageId)
	} else if fileType == CHATFILE {
		filePath = fmt.Sprintf(utils.UPLOADFILEPATHPREFIX, CHATFILE, imageId)
	} else {
		c.Error(errcode.ErrUploadFileToOssFailed)
		return
	}

	yunFilePath := filePath + imageId
	err = utils.UploadFileToOSS(dst, yunFilePath)
	if err != nil {
		logger.Logger.Warn("upload file to oss failed", zap.Any("loginSession", loginSession), zap.String("filename", header.Filename), zap.Error(err))
		c.Error(errcode.ErrUploadFileToOssFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
		"image_id": imageId,
	})
}

func UserReport(c *gin.Context) {
	loginSession := c.MustGet("session").(middlewares.LoginSession)

	var reqUserReport *pb.ReportReq
	if err := c.ShouldBindJSON(&reqUserReport); err != nil {
		fmt.Println(err)
		logger.Logger.Warn("bind user report  msg req  failed", zap.Any("user", loginSession))
		c.Error(errcode.ErrParam)
		return
	}

	reqUserReport.Uid = loginSession.Uid

	resp, err := commonClient.Report(c, reqUserReport)
	fmt.Println("resp, err", resp, err)
	if err != nil {
		logger.Logger.Warn("user report msg failed", zap.Any("user", loginSession), zap.Error(err))
		c.Error(errcode.ErrReportMsgFailed)
		return
	}

	if resp == nil {
		c.Error(errcode.ErrReportMsgFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "success",
	})
}
