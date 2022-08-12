package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Mysql `yaml:"mysql"`
}

type Mysql struct {
	DBName   string `yaml:"dbname"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int32  `yaml:"port"`
}

var (
	Cfg *Config //mysql配置文件信息
	DB  *gorm.DB
)

func init() {
	err := loadMysqlCfg() //加载数据库
	if err != nil {
		panic(err)
	}
	err = initDB() //链接数据库
	if err != nil {
		panic(err)
	}
}

func loadMysqlCfg() error {
	var config Config
	yamlFile, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return err
	}
	Cfg = &config
	return nil
}

func initDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", Cfg.UserName, Cfg.Password, Cfg.Host, Cfg.Port, Cfg.DBName)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	DB = db
	return nil
}
