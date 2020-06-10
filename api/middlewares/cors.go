package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Cors across domain
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
		   c.Next()后就执行真实的路由函数，路由函数执行完成之后继续执行后续的代码
		*/
		c.Next()
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Action, Module, X-PINGOTHER, Content-Type, Content-Disposition")
	}
}
