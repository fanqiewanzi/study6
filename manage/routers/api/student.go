package api

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"study6/manage/models"
)

//查询所有学生成绩
func GetAllGrade(context *gin.Context) {
	data := make(map[string]interface{})
	data["lists"] = models.GetStudent()
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "查询成功",
		"data": data,
	})
}

//插入学生信息
func InsertGrade(context *gin.Context) {
	var stu models.Studentgrade
	//结构体与json表单进行数据绑定
	context.ShouldBindJSON(&stu)
	ok := models.InsertStudent(stu)
	if ok {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "插入成功",
			"data": stu,
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "插入失败",
			"data": stu,
		})
	}
}

//根据学号设置成绩
func SetGrade(context *gin.Context) {
	var stu models.Studentgrade
	stu.Id = com.StrTo(context.Query("id")).MustInt()
	stu.Grade, _ = com.StrTo(context.Query("grade")).Float64()
	ok := models.SetGrade(stu.Id, stu.Grade)
	if ok {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "更新成功",
			"data": stu,
		})
	}
	context.JSON(http.StatusBadRequest, gin.H{
		"code": http.StatusBadRequest,
		"msg":  "更新失败",
		"data": stu,
	})
}

//升序输出所有学生成绩
func SortGrade(context *gin.Context) {
	data := make(map[string]interface{})
	data["lists"] = models.SortGrade()
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "查询成功",
		"data": data,
	})
}

//根据学号删除学生
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
