package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadTest() {
	r := gin.Default()
	r.POST("/mutiupload", func(context *gin.Context) {
		//从上下文中获取解析的表单
		form, err := context.MultipartForm()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "上传失败",
			})
			return
		}
		//从表单中的param upload[]中读取文件
		files := form.File["upload[]"]
		//将文件一个一个存储
		for _, file := range files {
			context.SaveUploadedFile(file, file.Filename)
		}
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "上传成功",
		})
	})
	r.GET("/getfile", func(context *gin.Context) {
		//从上下文中读取文件路径
		fpath := context.Query("url")
		//重命名，content-disposition是标题
		context.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fpath))
		//content-type是实际返回的类型,application/octet-stream是默认的未知的类型
		context.Writer.Header().Add("Content-Type", "application/octet-stream")
		//将文件路径对应的文件下载下来
		context.File(fpath)
	})
	r.Run()
}
