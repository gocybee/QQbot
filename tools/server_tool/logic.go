package server_tool

import (
	"QQbot/global"
	"QQbot/tools/analysis_tool"
	"QQbot/tools/rasa_tool"
	"fmt"
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

	if text.Msg == "" || PunctualOnly(text.Msg) {
		ResPondWithText(text.Id, "叫我干哈", text.Flag, true)
		return
	}

	intention := analysis_tool.IntentionJudge(text)
	// ChAT 交予rasa
	if intention == global.CHAT {
		answer := rasa_tool.GetRasaAnswer(text.RoutingID, text.Msg)
		if answer == "" {
			ResPondWithText(text.Id, "还不懂这句话的意思哦", text.Flag, true)
		}
		ResPondWithText(text.Id, answer, text.Flag, false)
		return
	} else {
		//其余搜寻答案即可
		an := analysis_tool.SelectAnswer(intention)

		fmt.Println("答案：：\n" + an)

		ResPondWithText(text.Id, an, text.Flag, true)
		return
	}
}
