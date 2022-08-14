package server_tool

import (
	"QQbot/global"
	"fmt"
	"strings"
)

//IsHeartBeat 判断是否为心跳事件
func IsHeartBeat(form map[string]interface{}) bool {
	if form["post_type"] == "meta_event" && form["meta_event_type"] == "heartbeat" {
		return true
	}
	return false
}

//IsPrivateMsg 判断是否为私聊消息
func IsPrivateMsg(form map[string]interface{}) bool {
	if form["post_type"] == "message" && form["message_type"] == "private" {
		return true
	}
	return false
}

//IsGroupMsg 判断是否为群消息
func IsGroupMsg(form map[string]interface{}) bool {
	if form["post_type"] == "message" && form["message_type"] == "group" {
		return true
	}
	return false
}

//IsAnonymous 是否为匿名消息
func IsAnonymous(form map[string]interface{}) bool {
	if form["anonymous"] != nil {
		return true
	}
	return false
}

//IsHelp 是否为帮助
func IsHelp(msg *string) bool {
	if strings.Contains(strings.ToLower(*msg), "help") || *msg == "帮助" {
		return true
	}
	return false
}

//IsStudy 是否触发学习程序
func IsStudy(msg *string) bool {
	if strings.Contains(*msg, "+") {
		return true
	}
	return false
}

//BeAt 看自己是否被@
func BeAt(str *string) bool {
	return strings.Contains(*str, fmt.Sprintf("[CQ:at,qq=%s]", global.MYQQID))
}

//NeedAsk 没有有效信息时是否需要反问
func NeedAsk(msg *string) bool {
	if strings.TrimSpace(*msg) == "" {
		return true
	}
	return false
}

//NeedSqlFire 私聊时需要导出数据库文件
func NeedSqlFire(msg *string) bool {
	if *msg == "导出问答文件" {
		return true
	}
	return false
}

//IsRepeated 是否出现了复读，打断
func IsRepeated(form map[string]interface{}, repeated *bool) (*int64, bool, string) {
	var f = false //标志是否找到消息
	if IsPrivateMsg(form) {
		idPtr, msgPtr, err := GetIdAndMsg(&form, global.PrivateFlag)
		if err != nil {
			return nil, false, ""
		}
		// 遍历寻找
		for i := 0; i < len(global.Re); i++ {
			//找到此人
			if global.Re[i].Id == *idPtr {
				f = true
				//同样的消息
				if global.Re[i].Content == *msgPtr {
					global.Re[i].Times++
					//告诉外界重复信息但不构成复读
					*repeated = true

					//触发复读
					if global.Re[i].Times > global.RepeatLimit {
						//清除内存
						global.Re[i].Times = 0
						return idPtr, true, global.PrivateFlag
					}
				} else {
					//消息更新
					global.Re[i].Content = *msgPtr
					global.Re[i].Times = 1
					return nil, false, ""
				}
			}
		}
		//没有找到则创建消息记录
		if !f {
			var r = global.Repeat{Id: *idPtr, Content: *msgPtr, Flag: global.PrivateFlag, Times: 1}
			global.Re = append(global.Re, r)
			return nil, false, ""
		}
	}
	if IsGroupMsg(form) {
		idPtr, msgPtr, err := GetIdAndMsg(&form, global.GroupFlag)
		if err != nil {
			return nil, false, ""
		}
		for i := 0; i < len(global.Re); i++ {
			//同样的群，同样的消息
			if global.Re[i].Id == *idPtr {
				f = true
				if global.Re[i].Content == *msgPtr {
					global.Re[i].Times++

					//告诉外界重复信息但不构成复读
					*repeated = true

					//触发复读
					if global.Re[i].Times > global.RepeatLimit {
						//清除内存
						global.Re[i].Times = 0
						return idPtr, true, global.GroupFlag
					}

				} else {
					//消息更新
					global.Re[i].Content = *msgPtr
					global.Re[i].Times = 0
					return nil, false, ""
				}
			}
		}
		if !f {
			var r = global.Repeat{Id: *idPtr, Content: *msgPtr, Flag: global.GroupFlag, Times: 1}
			global.Re = append(global.Re, r)
		}
	}
	return nil, false, ""
}
