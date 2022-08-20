package server_tool

import (
	"QQbot/global"
	"fmt"
	"strings"
)

// IsHeartBeat 判断是否为心跳事件
func IsHeartBeat(form map[string]interface{}) bool {
	if form["post_type"] == "meta_event" && form["meta_event_type"] == "heartbeat" {
		return true
	}
	return false
}

// IsPrivateMsg 判断是否为私聊消息
func IsPrivateMsg(form map[string]interface{}) bool {
	if form["post_type"] == "message" && form["message_type"] == "private" {
		return true
	}
	return false
}

// IsGroupMsg 判断是否为群消息
func IsGroupMsg(form map[string]interface{}) bool {
	if form["post_type"] == "message" && form["message_type"] == "group" {
		return true
	}
	return false
}

// IsAnonymous 是否为匿名消息
func IsAnonymous(form map[string]interface{}) bool {
	if form["anonymous"] != nil {
		return true
	}
	return false
}

// IsHelp 是否为帮助
func IsHelp(msg string) bool {
	if strings.Contains(strings.ToLower(msg), "help") || msg == "帮助" {
		return true
	}
	return false
}

// BeAt 看自己是否被@
func BeAt(str string) bool {
	return strings.Contains(str, fmt.Sprintf("[CQ:at,qq=%s]", global.MYQQID))
}

// IsRepeated 是否出现了复读，打断
func IsRepeated(form map[string]interface{}, repeated *bool) (int64, bool, string) {
	if IsPrivateMsg(form) {
		_, id, msg, err := GetIdAndMsg(form, global.PrivateFlag) //私聊时两id信息相同
		if err != nil {
			return 0, false, ""
		}
		return GetPossibleRepeatedMsg(id, msg, global.PrivateFlag, repeated)
	}
	if IsGroupMsg(form) {
		_, id, msg, err := GetIdAndMsg(form, global.GroupFlag) //群内复读不需要考虑个人
		if err != nil {
			return 0, false, ""
		}
		return GetPossibleRepeatedMsg(id, msg, global.GroupFlag, repeated)
	}
	return 0, false, ""
}
