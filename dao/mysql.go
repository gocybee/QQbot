package dao

import (
	"QQbot/global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.Mysql.UId, global.Mysql.Password, global.Mysql.Address, global.Mysql.Database,
	)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}

	//禁用复数
	db.SingularTable(true)

	// 判断是否有聊天白名单
	if !db.HasTable(&global.ChatWhiteList{}) {
		err = db.CreateTable(&global.ChatWhiteList{}).Error
		if err != nil {
			return err
		}
	}

	// 判断是否有回答黑名单
	if !db.HasTable(&global.BanedAnswerList{}) {
		err = db.CreateTable(&global.BanedAnswerList{}).Error
		if err != nil {
			return err
		}
	}

	// 判断是否有信息记录
	if !db.HasTable(&global.AnswerAndId{}) {
		err = db.CreateTable(&global.AnswerAndId{}).Error
		if err != nil {
			return err
		}
	}

	global.DB = db
	return nil
}
