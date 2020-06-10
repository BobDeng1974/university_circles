package middlewares

import (
	"go.uber.org/zap"
	"net/http"
	"university_circles/api/utils/common"
	errcode "university_circles/api/utils/errcode/user"
	"university_circles/api/utils/logger"

	"github.com/gin-gonic/gin"
)

type LoginSession struct {
	Uid          string `json:"uid"`
	ScreenName   string `json:"screenName"`
	UniversityId int64  `json:"university_id"`
}

// LoginSessionAuth middleware is check ipcc voice agent already logged in
func LoginSessionAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		token = c.DefaultQuery("token", c.PostForm("token"))
		if token == "" {
			token = c.GetHeader("token")
		}

		if token == "" {
			c.AbortWithError(http.StatusUnauthorized, c.Error(errcode.ErrLoginSessionInvalid))
			return
		}

		tokenClaims, err := utils.ParseToken(token)
		if err != nil {
			logger.Logger.Warn("get user token from redis failed", zap.Any("token", token), zap.Error(err))
			c.AbortWithError(http.StatusUnauthorized, c.Error(errcode.ErrLoginSessionInvalid))
			return
		}

		localToken, err := utils.GetLoginToken(tokenClaims.Uid)
		if localToken == "" {
			logger.Logger.Warn("user token nil", zap.Any("uid", tokenClaims.Uid), zap.Error(err))
			c.AbortWithError(http.StatusOK, c.Error(errcode.ErrUserNotLogin))
			return
		}

		loginSession := LoginSession{
			Uid:          tokenClaims.Uid,
			ScreenName:   tokenClaims.ScreenName,
			UniversityId: tokenClaims.UniversityId,
		}

		// 判断请求用户ID是否正确
		uid := c.Param("uid")
		if uid != "" {
			if uid != loginSession.Uid {
				c.AbortWithError(http.StatusForbidden, c.Error(errcode.ErrReqForbidden))
				return
			}
		}

		// 判断请求用户名是否正确
		username := c.Param("username")
		if username != "" {
			if username != loginSession.ScreenName {
				c.AbortWithError(http.StatusForbidden, c.Error(errcode.ErrReqForbidden))
				return
			}
		}

		// token刷新
		//if strings.Compare(token, localToken) >= 0 {
		//	if tokenClaims.ExpiresAt > 0 && (tokenClaims.ExpiresAt-time.Now().Unix()) < 300 {
		//		utils.RefreshLoginToken(tokenClaims.Uid, tokenClaims.ScreenName)
		//	}
		//} else {
		//	c.AbortWithError(http.StatusOK, c.Error(errcode.ErrLoginSessionInvalid))
		//	return
		//}

		c.Set("session", loginSession)
		c.Next()
	}
}
