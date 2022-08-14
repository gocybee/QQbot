//package tools 包含所有直接被service引用的函数

package server_tool

import (
	"QQbot/tools/dao_tool"
	"fmt"
)

//ResPondWithAsk 没有实际问题的回答
func ResPondWithAsk(id *int64, flag string) {
	t := "有什么事情吗？"
	status := Send(id, &t, flag) //发送信息
	fmt.Println(status)
}

//ResPondWithDBError 数据库出错返回
func ResPondWithDBError(id *int64, flag string) {
	text := "数据库炸了，寄"
	status := Send(id, &text, flag)
	fmt.Println(status)
}

//ResPondWithText 返回test信息
func ResPondWithText(id *int64, msg string, flag string) {
	Beautify(&msg)
	status := Send(id, &msg, flag)
	fmt.Println(status)
}

//ResPondWithTextPtr 返回test信息
func ResPondWithTextPtr(id *int64, msg *string, flag string) {
	Beautify(msg)
	status := Send(id, msg, flag)
	fmt.Println(status)
}

//RespondWhitSqlAndAI 出去特殊情况外的所有问答
func RespondWhitSqlAndAI(idPtr *int64, msgPtr *string, flag string) {
	//正常问答
	answerPtr := dao_tool.CalculateAnswer(msgPtr) //获取回答的语句
	//没有匹配到答案
	if answerPtr == nil {
		test, err := AIHelp(msgPtr)
		if err != nil {
			ResPondWithText(idPtr, "AI错误", flag)
			return
		}
		ResPondWithText(idPtr, test, flag)
		return
	}
	//匹配到了答案
	ResPondWithTextPtr(idPtr, answerPtr, flag)
}
