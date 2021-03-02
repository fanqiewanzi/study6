package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func LoadTest()  {
	r:=gin.Default()
	r.POST("/upload", func(context *gin.Context) {
		file,err:=context.FormFile("upload")
		if err!=nil{
			context.String(http.StatusBadRequest,"上传失败")
			return
		}
		if err:=context.SaveUploadedFile(file,file.Filename);err!=nil{
			context.String(http.StatusBadRequest,"\"保存失败 Error:%s\"", err.Error())
			return
		}
		context.String(http.StatusOK,"上传成功")
	})
	r.POST("/mutiupload", func(context *gin.Context) {
		form,err:=context.MultipartForm()
		if err!=nil{
			context.JSON(http.StatusBadRequest,"上传失败")
			return
		}
		files:=form.File["upload[]"]
		for _,file:=range files{
			context.SaveUploadedFile(file,file.Filename)
		}
		context.JSON(http.StatusOK,"上传成功")
	})
	r.GET("/getfile", func(context *gin.Context) {
		fpath:=context.Query("url")
		tp:=path.Base(fpath)
		context.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", tp))//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		context.Writer.Header().Add("Content-Type", "application/octet-stream")
		context.File(fpath)
	})
	r.Run()
}
