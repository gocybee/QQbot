package tools

import (
	"QQbot/global"
	"fmt"
	"math/rand"
	"strings"
)

//GetIdFromMap 从接受到的表单中提取出用户或者群聊Id
func GetIdFromMap(id interface{}) int64 {
	return int64(id.(float64))
}

//DoOrNot 生成随机数换算为概率--输入小数,现两位，默认0.5
func DoOrNot(p float32) bool {
	i := rand.Int() % 100
	if i < int(p*100) {
		return true
	}
	return false
}

//GetUsefulMsg 删去@自己部分（CQcode部分），获取消息的可被分析部分
func GetUsefulMsg(msg interface{}) string {
	str := msg.(string)
	return strings.TrimFunc(str, func(u rune) bool {
		if string(u) == fmt.Sprintf("[CQ:at,qq=%s]", global.MYQQID) {
			return true
		}
		return false
	})
}

//Beautify 为句子的头和尾美化
func Beautify(ctx *string) {
	//60%的概率做前部美化
	if DoOrNot(0.6) {
		i := rand.Int() % (len(global.Add))
		*ctx = global.Add[i] + *ctx
	}
	//60%的概率做尾部美化
	if DoOrNot(0.6) {
		i := rand.Int()%221 + 1
		*ctx += global.CodeCQFace(int64(i))
	}
}
