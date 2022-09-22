package dao_tool

import (
	"QQbot/global"
	"strings"
)

// NeedBan 是否触发黑名单-fathers@鼠鼠
func NeedBan(opp string, msgExtracted string) bool {
	for _, v := range global.Fathers {
		if v == opp {
			if strings.Contains(msgExtracted, "不可以") {
				return true
			}
			return false
		}
	}
	return false
}

// GenerateIdAndAnswerStr 由返回的 status 生成对应的结构体
func GenerateIdAndAnswerStr(status string, content string) global.AnswerAndId {
	m := []byte(status)
	var r int
	for i := 0; i < len(m); i++ {
		if m[i] == ']' {
			r = i
			break
		}
	}

	id := string(m[13:r])

	return global.AnswerAndId{
		MsgId:   id,
		Content: content,
	}
}
