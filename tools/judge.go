package tools

import (
	"github.com/lithammer/fuzzysearch/fuzzy"
	"strings"
	"sync"
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
func IsHelp(msg string) bool {
	if GetUsefulMsg(msg) == "-help" || GetUsefulMsg(msg) == "帮助" || strings.Contains(msg, "你能干什么") {
		return true
	}
	return false
}

//MatchDistance 获取信息距离数据库中问题的距离
func MatchDistance(msg string, sql string) int {
	mSlice := SplitMsg(msg)
	var X = -1 //匹配度 0-完全匹配
	for _, v := range mSlice {
		x := fuzzy.RankMatch(v, sql)
		if x < 0 {
			continue
		}
		//出现匹配的树则统计距离
		var once sync.Once
		once.Do(func() { X = 0 })
		X += x
	}
	return X
}
