package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"study6/manage/models"
	"study6/manage/pkg/exception"
	"study6/manage/pkg/util"
)

//登录逻辑
func Login(c *gin.Context) {
	var user models.User
	var token string
	var err error
	code := exception.SUCCESS

	//struct绑定json
	c.ShouldBindJSON(&user)

	//通过用户名查找用户
	u := models.GetUserByName(user)
	if user.Password == u.Password {
		token, err = util.CreatJwt(user.UserName, user.Password)
		if err != nil {
			code = exception.ERROR_AUTH_CREAT_TOKEN
		}
	} else {
		code = exception.ERROR_WRONG_ELEMENT
	}
	c.JSON(code, gin.H{
		"code":    code,
		"message": exception.GetMsg(code),
		"data":    user,
		"token":   token,
	})

}

//注册逻辑
func Register(c *gin.Context) {
	var user models.User
	var code int
	//绑定Json
	c.ShouldBindJSON(&user)
	ok := models.InsertUser(user)
	if ok {
		code = exception.SUCCESS
	} else {
		code = exception.ERROR_SAME_NAME
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": exception.GetMsg(code),
	})
}
