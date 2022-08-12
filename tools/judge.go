package tools

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
