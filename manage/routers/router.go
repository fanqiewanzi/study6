package routers

import (
	"github.com/gin-gonic/gin"
	"study6/manage/middleware"
	"study6/manage/routers/api"
)

//初始化gin引擎
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Log())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/student")
	{
		apiv1.GET("/getallgrade", api.GetAllGrade)
		apiv1.POST("/insertgrade", api.InsertGrade)
		apiv1.POST("/setgrade", api.SetGrade)
		apiv1.GET("/sortgrade", api.SortGrade)
		apiv1.DELETE("/delete", api.Delete)
	}

	return r
}
