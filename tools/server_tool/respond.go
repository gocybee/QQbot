// package tools 包含所有直接被service引用的函数

package server_tool

import (
	"QQbot/global"
	"QQbot/tools/dao_tool"
	"fmt"
)

// ResPondWithAsk 没有实际问题的回答
func ResPondWithAsk(id int64, flag string) {
	t := "叫我干哈"
	Beautify(&t)
	status := send(&id, &t, flag) // 发送信息
	fmt.Println(status)
}

// ResPondWithDBError 数据库出错返回
func ResPondWithDBError(id int64, flag string) {
	text := "数据库炸了，寄"
	status := send(&id, &text, flag)
	fmt.Println(status)
}

// ResPondWithText 返回test信息
func ResPondWithText(id int64, msg string, flag string) {
	Beautify(&msg)
	status := send(&id, &msg, flag)
	fmt.Println(status)
}

// ResPondWithTextPtr 返回test信息
func ResPondWithTextPtr(id int64, msg string, flag string) {
	Beautify(&msg)
	status := send(&id, &msg, flag)
	fmt.Println(status)
}

// RespondWhitSqlAndAI 出去特殊情况外的所有问答
func RespondWhitSqlAndAI(idPtr int64, msgPtr string, flag string) {
	// 正常问答
	answerPtr := dao_tool.CalculateAnswer(msgPtr) // 获取回答的语句
	// 没有匹配到答案
	if answerPtr == "" {
		test, err := AIHelp(&msgPtr)
		if err != nil {
			ResPondWithText(idPtr, "AI错误", flag)
			return
		}
		ResPondWithText(idPtr, test, flag)
		return
	}
	// 匹配到了答案
	ResPondWithTextPtr(idPtr, answerPtr, flag)
}

// ResPondWithPhoto 返回非闪照的图片
func ResPondWithPhoto(id *int64, fileName string, url string, flag string) {
	msg := global.CodeCQPhoto(fileName, url)
	status := send(id, &msg, flag)
	fmt.Println(status)
}

// ResPondWithTextAndPhoto 返回信息及非闪照的图片
func ResPondWithTextAndPhoto(id int64, msg string, fileName string, url string, flag string) {
	msg += global.CodeCQPhoto(fileName, url)
	status := send(&id, &msg, flag)
	fmt.Println(status)
}
