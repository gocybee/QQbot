package tools

import (
	"QQbot/global"
	"math/rand"
	"strings"
)

//GetIdFromMap 从接受到的表单中提取出用户或者群聊Id
func GetIdFromMap(id interface{}) int64 {
	return int64(id.(float64))
}

//BeAt 看自己是否被@
func BeAt(str interface{}) bool {
	msg := str.(string)
	return strings.Contains(msg, "at") && strings.Contains(msg, global.MYQQID)
}

//DoOrNot 生成随机数换算为概率--输入小数,现两位，默认0.5
func DoOrNot(p float32) bool {
	i := rand.Int() % 100
	if i < int(p*100) {
		return true
	}
	return false
}

//GetUsefulMsg 删去@部分（CQcode部分），获取消息的可被分析部分
func GetUsefulMsg(msg interface{}) string {
	code := strings.LastIndex(msg.(string), "]")
	if code == -1 {
		return msg.(string)
	}
	str := msg.(string)
	return string([]byte(str)[code+2:]) //信息的code后含有一个空格
}

//SplitMsg 将信息拆分成两个字，便于模糊匹配
func SplitMsg(msg string) []string {
	msg += " "
	var res []string
	var m = []rune(msg)
	for i := 0; i < len(m)-2; i++ {
		t := string(m[i]) + string(m[i+1])
		res = append(res, t)
	}
	return res
}
