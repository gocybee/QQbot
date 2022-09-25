package dao

import (
	"QQbot/global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Init() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.Mysql.User, global.Mysql.Password, global.Mysql.Address, global.Mysql.DbName,
	)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 禁用复数
	db.SingularTable(true)

	// 判断是否有聊天白名单
	if !db.HasTable(&global.ChatWhiteListStruct{}) {
		err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&global.ChatWhiteListStruct{}).Error
		if err != nil {
			return err
		}
	}

	// 判断是否有回答黑名单
	if !db.HasTable(&global.BannedAnswerListStruct{}) {
		err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&global.BannedAnswerListStruct{}).Error
		if err != nil {
			return err
		}
	}

	// 判断是否有信息记录
	if !db.HasTable(&global.AnswerAndIdStruct{}) {
		err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&global.AnswerAndIdStruct{}).Error
		if err != nil {
			return err
		}
	}

	global.DB = db

	//将聊天白名单写入数据库
	err = writChatWhiteList(global.ChatList)
	if err != nil {
		return err
	}

	return nil
}
