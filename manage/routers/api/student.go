package api

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"study6/manage/models"
)

func GetAllGrade(context *gin.Context) {
	data := make(map[string]interface{})
	data["lists"] = models.GetStudent()
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "查询成功",
		"data": data,
	})
}

func InsertGrade(context *gin.Context) {
	var stu models.Studentgrade
	stu.Id = com.StrTo(context.Query("id")).MustInt()
	stu.Name = context.Query("name")
	stu.Grade, _ = com.StrTo(context.Query("grade")).Float64()
	models.InsertStudent(stu)
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "插入成功",
		"data": stu,
	})
}
func SetGrade(context *gin.Context) {
	var stu models.Studentgrade
	stu.Id = com.StrTo(context.Query("id")).MustInt()
	stu.Grade, _ = com.StrTo(context.Query("grade")).Float64()
	models.SetGrade(stu.Id, stu.Grade)
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "更新成功",
		"data": stu,
	})
}

func SortGrade(context *gin.Context) {
	data := make(map[string]interface{})
	data["lists"] = models.SortGrade()
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "查询成功",
		"data": data,
	})
}
func Delete(context *gin.Context) {
	id := com.StrTo(context.Query("id")).MustInt()
	if models.DeleteById(id) {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "删除成功",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "删除失败",
		})
	}
}
