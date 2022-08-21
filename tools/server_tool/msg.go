package server_tool

import (
	"QQbot/global"
	"math/rand"
)

// Beautify 为句子的尾部美化
func Beautify(ctx *string) {
	// 60%的概率做尾部美化
	if DoOrNot(0.6) {
		i := rand.Int()%221 + 1
		// 避开奇怪的表情
		if (i > 40 && i < 92) || (i > 111 && i < 172) || i > 183 {
			i = 179
		}
		*ctx += global.CodeCQFace(int64(i))
	}
}

// RegisterRepeated 基于消息源创建信息记录
func RegisterRepeated(rmPtr *global.ReceivedMsg) {
	var r = global.Repeat{
		Id:      rmPtr.GetOppositeIdInt64(),
		Content: rmPtr.GetMsg(),
		Flag:    rmPtr.GetGlobalFlag(),
		Times:   1}
	global.Repeated[rmPtr.GetOppositeIdStr()] = &r
}
