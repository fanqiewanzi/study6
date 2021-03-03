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
	maps := make(map[string]interface{})
	maps["id"] = com.StrTo(context.Query("id")).MustInt()
	maps["name"] = context.Query("name")
	maps["grade"] = com.StrTo(context.Query("grade")).MustFloat64()
	models.InsertStudent(maps)
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "插入成功",
		"data": maps,
	})
}
func SetGrade(context *gin.Context) {

}

func SortGrade(context *gin.Context) {

}
func Delete(context *gin.Context) {

}
