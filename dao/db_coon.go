package dao

import (
	"blog-go/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var globalDb *gorm.DB

func init() {
	InitMysqlCfgCoon()
}

func InitMysqlCfgCoon() {
	mysqlCfg := config.GetMySqlCfg()
	//"user:pass@tcp(127.0.0.1:3306)/blog_go?charset=utf8mb4&parseTime=True&loc=Local"
	url := mysqlCfg.Username + ":" + mysqlCfg.Password + "@tcp(" + mysqlCfg.Host + ":" + mysqlCfg.Port + ")/" + mysqlCfg.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	} else {
		log.Println("db connection succedssed")
	}

	globalDb = db
}
