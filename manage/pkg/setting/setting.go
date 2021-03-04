package setting

import (
	"github.com/go-ini/ini"
	"log"
)

var (
	Cfg *ini.File
)

func init() {
	var err error
	Cfg, err = ini.Load("D:\\GoProject\\study6\\manage\\conf\\manage.ini")
	if err != nil {
		log.Fatalf("Fail to parse ini: %v", err)
	}
}
