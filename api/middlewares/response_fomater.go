package middlewares

import (
	"net/http"
	"university_circles/api/utils/errcode"

	"github.com/gin-gonic/gin"
)

// ResponseHandler across domain
func ResponseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"err_code": http.StatusInternalServerError,
					"err_msg":  err,
					"data":     "",
				})
				return
			}
		}()
		/*
		   c.Next()后就执行真实的路由函数，路由函数执行完成之后继续执行后续的代码
		*/
		c.Next()
		err := c.Errors.ByType(gin.ErrorTypeAny).Last()
		if err != nil {
			if err.Meta != nil {
				c.JSON(http.StatusOK, err.Meta)
			} else {
				if e, ok := err.Err.(errcode.StandardError); ok {
					c.JSON(http.StatusOK, gin.H{
						"err_code": e.Code,
						"err_msg":  e.Msg,
					})
					c.Abort()
				} else {
					c.JSON(http.StatusOK, gin.H{
						"err_code": http.StatusInternalServerError,
						"err_msg":  err.Error(),
						"data":     "",
					})
				}
			}
		}
	}
}
