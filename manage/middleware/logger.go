package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Log() gin.HandlerFunc {
	//设定存储log文件的路径
	filepath := "log/log"
	//打开路径下的文件，os.O_RDWR指的是权限为读写，os.O_CREATE指的是若不存在路径文件就创建一个文件,os.ModeAppend是只能增加,0755是linux的权限
	sc, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println("err:", err)
	}
	//logrus框架创建一个新的Logger
	logger := logrus.New()
	//指定打开的文件为logger的输出，会输出一个io.copy出去
	logger.Out = sc
	//设置logger的等级为debug
	logger.SetLevel(logrus.DebugLevel)

	return func(c *gin.Context) {
		//开始计时
		startTime := time.Now()
		//调用后续的处理函数
		c.Next()
		//处理函数运行完后停止运行并计算运行时间
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))
		//查找主机名
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		//获取运行状态，客户端ip，用户代理和数据大小
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		//获取请求方法和路径
		method := c.Request.Method
		path := c.Request.RequestURI
		//创建logger的入口，里面的fields数据结构就是一个map[string]interface{}
		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})
		//将数据写入log
		entry.Info()

	}
}
