package server_tool

import (
	"QQbot/global"
	"QQbot/tools/analysis_tool"
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
		ResPondWithText(text.Id, "我只会一点点欸，主要是开发大大太菜了(help写出来也没几个字)", text.Flag, true)
		return
	}

	if text.Msg == "" || PunctualOnly(text.Msg) {
		if DoOrNot(0.5) {
			ResPondWithText(text.Id, "咋了", text.Flag, true)
		} else {
			ResPondWithText(text.Id, "叫我干哈", text.Flag, true)
		}

		return
	}

	intention := analysis_tool.IntentionJudge(text)
	// CHAT 交予rasa
	if intention == global.CHAT {
		answer, err := rasa_tool.GetRasaAnswer(text.Session, text.Msg)
		ResPondWithText(text.Id, answer, text.Flag, false)
		// 没有回复结果

		if answer == "" || err != nil {
			ResPondWithText(text.Id, "后台又双叒叕不和我玩了", text.Flag, false)
			ResPondWithPhoto(text.Id, "b564900eded645e5c523f5534b14ab1b.image", "https://gchat.qpic.cn/gchatpic_new/918845478/920689543-2538037296-B564900EDED645E5C523F5534B14AB1B/0?term=3", text.Flag)
		}
		return

	} else {
		// 其余搜寻答案即可
		an := analysis_tool.SelectAnswer(intention)

		ResPondWithText(text.Id, an, text.Flag, true)
		return
	}
}

// PostChanMsgToRouting 向全局的协程发送回复目标的信息
func PostChanMsgToRouting(l *global.RoutingMsg, rmPtr *global.ReceivedMsg) {
	t := global.ChanMsg{
		Id:        rmPtr.GetOppositeIdInt64(),
		Msg:       rmPtr.GetMsg(),
		Flag:      rmPtr.GetGlobalFlag(),
		Repeated:  rmPtr.IsRepeated(),
		Session:   l.Session,
		RoutingID: rmPtr.GetSenderIdStr(),
	}
	l.C <- &t
}
