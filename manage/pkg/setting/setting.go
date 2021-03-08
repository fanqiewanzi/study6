package setting

import (
	"github.com/go-ini/ini"
	"log"
)

var (
	Cfg       *ini.File
	JwtSecret string
)

//读取ini文件初始化
func init() {
	var err error
	Cfg, err = ini.Load("conf/manage.ini")
	if err != nil {
		log.Fatalf("Fail to parse ini: %v", err)
	}
	load()
}

//加载JWT配置
func load() {
	sc, err := Cfg.GetSection("JWT")
	if err != nil {
		log.Fatalf("Fail to get: %v", err)
	}
	JwtSecret = sc.Key("SECRET").String()
}
