package server_tool

import (
	"QQbot/global"
)

// RespondLogic 回答问题
func RespondLogic(text *global.ChanMsg) {
	// 消息重复
	if text.Repeated {
		ResPondWithText(text.Id, "刚刚才回答过哦", text.Flag)
		return
	}

	if IsHelp(text.Msg) {
		ResPondWithText(text.Id, "我只会一点点欸，主要是开发大大太菜了", text.Flag)
		return
	}

	//TODO 第一次则获取rasa的回复和session然后将session加入global.Routing中
	//if text.Session==""{
	//	////发送text.Msg 获取 oldQueId 和 session
	//	//global.Routing[text.RoutingID].OldQueId = oldQueId
	//	//global.Routing[text.RoutingID].Session = session
	//
	//}else{
	//	////如果语义相同
	//	//oldQueId = 语义分析接口
	//	//if text.OldQueId == oldQueId{
	//	//	ResPondWithText(text.Id, "答案不和刚刚一样的嘛", text.Flag)
	//	//	return
	//	//}
	//	////正常回答
	//	//oldQueId, session, response = 问答接口
	//	//ResPondWithText(text.Id, response, text.Flag)
	//	////更新语义id
	//	//global.Routing[text.RoutingID].OldQueId = oldQueId
	//}

	ResPondWithText(text.Id, "我收到消息"+text.Msg, text.Flag)

	return
}
