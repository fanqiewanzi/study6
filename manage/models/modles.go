package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"study6/manage/pkg/setting"
)

var db *gorm.DB

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	sc, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType = sc.Key("TYPE").String()
	dbName = sc.Key("NAME").String()
	user = sc.Key("USER").String()
	password = sc.Key("PASSWORD").String()
	host = sc.Key("HOST").String()
	tablePrefix = sc.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix
	}

}
func Close() {
	defer db.Close()
}
