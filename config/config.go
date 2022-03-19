package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

/**
配置文件映射 ``表示标签，yaml表示映射到字段mysql
*/
type Config struct {
	Mysql MySqlCfg `yaml:"mysql"`
}

/**
Mysql 数据库配置
*/
type MySqlCfg struct {
	Username string
	Password string
	Database string
	Host     string
	Port     string
}

var Cfg Config

func init() {
	file, err := ioutil.ReadFile("F:\\projects\\goprojects\\src\\blog-go\\config\\config.yml")
	if err != nil {
		log.Fatal(err)
	}

	//yaml文件内容影射到结构体中
	err1 := yaml.Unmarshal(file, &Cfg)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func GetMySqlCfg() (mysql MySqlCfg) {
	mysqlCfg := Cfg.Mysql
	mysql = mysqlCfg
	return mysql
}
