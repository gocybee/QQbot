package config

import (
	"QQbot/global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Mysql `yaml:"mysql"` //数据库配置
	Res   []QA           `yaml:"res"` //问答初始化
}

type Mysql struct {
	DBName   string `yaml:"dbname"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int32  `yaml:"port"`
}

//QA 用于创建数据库并储存相关的信息
type QA struct {
	ID     int32  `gorm:"AUTO_INCREMENT" gorm:"id"`
	Q1     string `gorm:"type:char(25)" yaml:"q1" gorm:"q1"`
	Q2     string `gorm:"type:char(25)" yaml:"q2" gorm:"q2"`
	Q3     string `gorm:"type:char(25)" yaml:"q3" gorm:"q3"`
	Answer string `gorm:"type:char(255)" yaml:"answer" gorm:"answer"`
}

var cfg *Config //mysql配置文件信息,数据库初始信息

func init() {
	err := loadCfg() //加载数据库
	if err != nil {
		panic(err)
	}
	err = initDB() //链接数据库
	if err != nil {
		panic(err)
	}
	err = loadQA() //加载问题和答案
	if err != nil {
		panic(err)
	}
}

func loadCfg() error {
	var config Config
	yamlFile, err := ioutil.ReadFile(global.CfgFileURL)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return err
	}
	cfg = &config
	return nil
}

func initDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := gorm.Open("mysql", dsn)
	db.SingularTable(true) //禁用复数表名
	if err != nil {
		return err
	}
	global.DB = db
	return nil
}

func loadQA() error {
	if !global.DB.HasTable(&QA{}) {
		err := global.DB.CreateTable(&QA{}).Error
		if err != nil {
			return err
		}
		//逐条插入语句和答案
		err = global.DB.Transaction(func(tx *gorm.DB) error {
			for _, v := range cfg.Res {
				err = global.DB.Create(&v).Error
				if err != nil {
					return err
				}
			}
			return nil
		})
	}
	return nil
}
