package server_tool

import (
	"QQbot/global"
	"math/rand"
	"strings"
)

// Beautify 为句子的尾部美化
func Beautify(ctx *string) {
	// 40%的概率做尾部美化
	if DoOrNot(0.4) {
		i := rand.Int()%221 + 1
		// 避开奇怪的表情
		if (i > 40 && i < 92) || (i > 111 && i < 172) || i > 183 {
			i = 179
		}
		*ctx += global.CodeCQFace(int64(i))
	}
}

// RegisterRepeated 基于消息源创建信息记录-便于统计复读
func RegisterRepeated(rmPtr *global.ReceivedMsg) {
	var r = global.Repeat{
		Id:      rmPtr.GetOppositeIdInt64(),
		Content: rmPtr.GetMsg(),
		Flag:    rmPtr.GetGlobalFlag(),
		Times:   1,
	}
	global.Repeated[rmPtr.GetOppositeIdStr()] = &r
}

// Escape 将rasa回复的特殊字符转义为url可发送的
func Escape(rasaMsg string) string {
	s := strings.Split(rasaMsg, " ")
	re := s[0]
	for i := 1; i < len(s); i++ {
		re += "%20" + s[i]
	}
	return re
}
