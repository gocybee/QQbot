package server_tool

import (
	"QQbot/global"
	"fmt"
	"math/rand"
	"strings"
	"unicode"
)

// IsHeartBeat 判断是否为心跳事件
func IsHeartBeat(form map[string]interface{}) bool {
	if form["post_type"] == "meta_event" && form["meta_event_type"] == "heartbeat" {
		return true
	}
	return false
}

// IsPrivateMsg 判断是否为私聊消息
func IsPrivateMsg(flag string) bool {
	if flag == "private" {
		return true
	}
	return false
}

// IsGroupMsg 判断是否为群消息
func IsGroupMsg(flag string) bool {
	if flag == "group" {
		return true
	}
	return false
}

// IsHelp 是否为帮助
func IsHelp(mt string) bool {
	if strings.Contains(strings.ToLower(mt), "help") || mt == "帮助" {
		return true
	}
	return false
}

// BeAt 看自己是否被@
func BeAt(mt string) bool {
	return strings.Contains(mt, fmt.Sprintf("[CQ:at,qq=%s]", global.MYQQID))
}

//PunctualOnly 是否只有符号
func PunctualOnly(str string) bool {
	s := []rune(str)
	for _, v := range s {
		if !unicode.IsPunct(v) {
			return false
		}
	}
	return true
}

// IsMsgRepeated 是否出现了语言内容重复型复读
//返回 是否出现复读
func IsMsgRepeated(rmPtr *global.ReceivedMsg) bool {
	opp := rmPtr.GetOppositeIdStr()
	if r, ok := global.Repeated[opp]; ok {
		r.Lock()
		defer r.Unlock()
		if r.Content == rmPtr.GetMsg() {
			//标记信息重复
			rmPtr.Repeated()
			r.Times++
			if r.Times > global.RepeatLimit {
				//重置
				delete(global.Repeated, opp)
				return true
			}
		} else {
			//更新信息
			r.Content = rmPtr.GetMsg()
		}
	} else {
		//需要创建信息记录
		RegisterRepeated(rmPtr)
	}
	return false
}

// DoOrNot 生成随机数换算为概率--输入小数,现两位
func DoOrNot(p float32) bool {
	i := rand.Int() % 100
	if i < int(p*100) {
		return true
	}
	return false
}
