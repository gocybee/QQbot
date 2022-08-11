package tools

import (
	"QQbot/runtime"
	"strings"
)

//JudgeAndResp 关键词判断,并返回需要回复的信息
func JudgeAndResp(msg string, uId int64, isAnon bool) string {
	//群聊中被@
	if strings.Contains(msg, "at") {
		//打招呼
		if strings.Contains(msg, "你好啊") {
			return runtime.CodeCQAt(uId) + "，俺现在啥也不知道哦？"
		}
		return runtime.CodeCQAt(uId) + "，俺现在啥也不知道哦？"
	}
	//私聊中啥也不需要
	return "俺现在啥也不知道哦"
}
