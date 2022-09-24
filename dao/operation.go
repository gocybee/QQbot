package dao

import (
	"QQbot/global"
	"github.com/jinzhu/gorm"
)

var number = 0 //记录信息条数

// Banned 在所有情况下设置不能说的句子
func Banned(msgId string) error {
	// 在全局记录中寻找id对应的信息
	var c global.AnswerAndIdStruct
	err := global.DB.Model(&global.AnswerAndIdStruct{}).Where("msg_id = ?", msgId).Find(&c).Error
	if err != nil {
		return err
	}
	// 将其写入回答黑名单
	var t = global.BannedAnswerListStruct{
		Baned: c.Content,
	}
	return global.DB.Model(&global.BannedAnswerListStruct{}).Create(&t).Error
}

// Filter 将rasa的回答过滤一遍，
func Filter(answer *string) {
	var t global.BannedAnswerListStruct
	if !global.DB.Model(&global.BannedAnswerListStruct{}).Where("baned = ?", *answer).First(&t).RecordNotFound() {
		*answer = "这。。。不好说"
	}
}

// CanChatWith 是否在聊天白名单内
func CanChatWith(opp string) bool {
	// debug标志
	if global.Debug == true {
		return true
	}

	var t global.ChatWhiteListStruct
	// 白名单的判断
	if !global.DB.Model(&global.ChatWhiteListStruct{}).Where("uid = ?", opp).First(&t).RecordNotFound() {
		return true
	}
	return false
}

// WritIdAndAnswer 将信息写入数据库-只存500条
func WritIdAndAnswer(x global.AnswerAndIdStruct) error {
	var err error
	number++
	if number >= 500 {
		err = global.DB.Model(&global.AnswerAndIdStruct{}).Where("id=?", number-499).Update(&x).Error
	} else {
		err = global.DB.Model(&global.AnswerAndIdStruct{}).Create(&x).Error
	}
	if err != nil {
		return err
	}
	return nil
}

// witChatWhiteList 初始化数据库时初始化白名单
func writChatWhiteList(uid []string) error {
	if len(uid) == 0 {
		return nil
	}

	//删除已经初始化的信息-测试时不用删库
	global.DB.Model(&global.ChatWhiteListStruct{}).Delete(&global.ChatWhiteListStruct{})

	err := global.DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range uid {
			t := global.ChatWhiteListStruct{Uid: v}
			err := global.DB.Model(&global.ChatWhiteListStruct{}).Create(&t).Error
			if err != nil {
				return err
			}
		}
		// 返回 nil 提交事务
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
