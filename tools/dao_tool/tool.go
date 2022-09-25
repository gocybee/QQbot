package dao_tool

import (
	"QQbot/global"
	"regexp"
	"strings"
)

// NeedBan 是否触发黑名单-fathers@鼠鼠
func NeedBan(opp string, msgExtracted string) bool {

	for _, v := range global.Controller {
		if v == opp {
			if strings.Contains(msgExtracted, "不能这么说") {
				return true
			}
			return false
		}
	}
	return false
}

// GetReplyMsgId 从他人引用回复的话提取信息
func GetReplyMsgId(status string) string {
	//提取id
	s := regexp.MustCompile("CQ:reply,id=(.\\d*)]")
	_id := s.FindStringSubmatch(status)

	return _id[len(_id)-1]
}

// GenerateIdAndAnswerStr 发送信息后由返回的status生成对应的结构体
func GenerateIdAndAnswerStr(status string, content string) global.AnswerAndIdStruct {
	//提取id
	s := regexp.MustCompile("\"message_id\":(.\\d*)")
	_id := s.FindStringSubmatch(status)

	return global.AnswerAndIdStruct{
		MsgId:   _id[1],
		Content: content,
	}
}
