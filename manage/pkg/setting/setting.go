package setting

import (
	"github.com/go-ini/ini"
	"log"
)

var (
	Cfg *ini.File
)

//读取ini文件初始化
func init() {
	var err error
	Cfg, err = ini.Load("conf/manage.ini")
	if err != nil {
		log.Fatalf("Fail to parse ini: %v", err)
	}
}
