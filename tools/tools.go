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

//NeedResp 看自己是否被@判断需不需要回答
func NeedResp(str interface{}) bool {
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
