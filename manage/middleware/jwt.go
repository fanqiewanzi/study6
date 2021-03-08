package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"study6/manage/pkg/exception"
	"study6/manage/pkg/util"
	"time"
)

func JWT() gin.HandlerFunc {

	return func(context *gin.Context) {
		code := exception.SUCCESS
		var data interface{}

		//从URL中读取token字符串
		token := context.Query("token")

		//若token为空则直接返回
		//不为空则解析token
		if token == "" {
			code = exception.ERROR_AUTH
		} else {
			claim, err := util.ParseToken(token)
			data = claim
			//解析失败
			if err != nil {
				code = exception.ERROR_AUTH_CHECK_TOKEN_FAIL
				//TOKEN过期
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = exception.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		//返回Json
		if code != exception.SUCCESS {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  exception.GetMsg(code),
				"data": data,
			})
			//停止调用挂起链的处理函数
			context.Abort()
			return
		}

		//通过TOKEN验证
		context.Next()
	}
}
