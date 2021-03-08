package api

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"study6/manage/models"
	"study6/manage/pkg/exception"
)

//查询所有学生成绩
func GetAllGrade(context *gin.Context) {
	data := make(map[string]interface{})
	code := exception.SUCCESS
	data["lists"] = models.GetStudent()
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  exception.GetMsg(code),
		"data": data,
	})
}

//插入学生信息
func InsertGrade(context *gin.Context) {
	var stu models.Studentgrade
	var code int
	//结构体与json表单进行数据绑定
	context.ShouldBindJSON(&stu)
	ok := models.InsertStudent(stu)
	if ok {
		code = exception.SUCCESS
	} else {
		code = exception.ERROR
	}
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  exception.GetMsg(code),
		"data": stu,
	})
}

//根据学号设置成绩
func SetGrade(context *gin.Context) {
	var stu models.Studentgrade
	var code int
	stu.Id = com.StrTo(context.Query("id")).MustInt()
	stu.Grade, _ = com.StrTo(context.Query("grade")).Float64()
	ok := models.SetGrade(stu.Id, stu.Grade)
	if ok {
		code = exception.SUCCESS
	} else {
		code = exception.ERROR
	}
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  exception.GetMsg(code),
		"data": stu,
	})
}

//升序输出所有学生成绩
func SortGrade(context *gin.Context) {
	data := make(map[string]interface{})
	data["lists"] = models.SortGrade()
	code := exception.SUCCESS
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  exception.GetMsg(code),
		"data": data,
	})
}

//根据学号删除学生
func Delete(context *gin.Context) {
	id := com.StrTo(context.Query("id")).MustInt()
	var code int
	if models.DeleteById(id) {
		code = exception.SUCCESS
	} else {
		code = exception.ERROR
	}
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  exception.GetMsg(code),
	})
}
