package server_tool

import (
	"QQbot/global"
	"QQbot/tools/rasa_tool"
)

// RespondLogic 回答问题
func RespondLogic(text *global.ChanMsg) {
	// 消息重复
	if text.Repeated {
		ResPondWithText(text.Id, "刚刚才回答过哦", text.Flag, true)
		return
	}

	if IsHelp(text.Msg) {
		ResPondWithText(text.Id, "我只会一点点欸，主要是开发大大太菜了", text.Flag, true)
		return
	}

	//TODO 第一次则获取rasa的回复和session然后将session加入global.Routing中
	answer := rasa_tool.GetRasaAnswer(text.RoutingID, text.Msg)
	if answer == "" {
		ResPondWithText(text.Id, "还不懂这句话的意思哦", text.Flag, true)
	}
	ResPondWithText(text.Id, answer, text.Flag, false)

	return
}
