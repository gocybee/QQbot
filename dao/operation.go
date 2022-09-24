package dao

import (
	"QQbot/global"
)

var number = 0 //记录信息条数

// Banned 在所有情况下设置不能说的句子
func Banned(msgId string) error {
	// 在全局记录中寻找id对应的信息
	var c global.AnswerAndId
	err := global.DB.Model(&global.AnswerAndId{}).Where("msg_id = ?", msgId).Find(&c).Error
	if err != nil {
		return err
	}
	// 将其写入回答黑名单
	var t = global.BanedAnswerList{
		Baned: c.Content,
	}
	return global.DB.Model(&global.BanedAnswerList{}).Create(&t).Error
}

// Filter 将rasa的回答过滤一遍，
func Filter(answer *string) {
	var t global.BanedAnswerList
	if !global.DB.Model(&global.BanedAnswerList{}).Where("baned = ?", *answer).First(&t).RecordNotFound() {
		*answer = "这。。。不好说"
	}
}

// CanChatWith 是否在聊天白名单内
func CanChatWith(opp string) bool {
	var t global.ChatWhiteList

	// 其他白名单的判断
	if !global.DB.Model(&global.ChatWhiteList{}).Where("uid = ?", opp).First(&t).RecordNotFound() {
		return true
	}
	return false
}

// WritIdAndAnswer 将信息写入数据库-只存500条
func WritIdAndAnswer(x global.AnswerAndId) error {
	var err error
	number++
	if number >= 500 {
		x.Id = uint(number - 499)
		err = global.DB.Model(&global.AnswerAndId{}).Where("id=?", number-499).Update(&x).Error
	} else {
		err = global.DB.Model(&global.AnswerAndId{}).Create(&x).Error
	}
	if err != nil {
		return err
	}
	return nil
}
