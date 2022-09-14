package server_tool

import (
	"QQbot/global"
	"QQbot/tools/analysis_tool"
	"QQbot/tools/rasa_tool"
	"time"
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
		ResPondWithText(text.Id, "叫我干哈", text.Flag, true)
		return
	}

	intention := analysis_tool.IntentionJudge(text)
	// ChAT 交予rasa
	if intention == global.CHAT {
		var answer = make(chan string, 1)
		//超过0.5s没有获取回复算超时
		select {
		case <-time.After(time.Millisecond):
			ResPondWithText(text.Id, "阿勒，后台失联了", text.Flag, false)

			ResPondWithPhoto(text.Id, "90e4a8323deb495bdef7086b618269e7.image", "https://gchat.qpic.cn/gchatpic_new/2505772098/920689543-3065492934-90E4A8323DEB495BDEF7086B618269E7/0?term=3", text.Flag)
		case answer <- rasa_tool.GetRasaAnswer(text.Session, text.Msg):
			ResPondWithText(text.Id, <-answer, text.Flag, false)
		}
		//没有回复结果
		if <-answer == "" {
			ResPondWithTextAndPhoto(text.Id, "问了一下后台，结果是。。。%0a我不会", "b564900eded645e5c523f5534b14ab1b.image", "https://gchat.qpic.cn/gchatpic_new/918845478/920689543-2538037296-B564900EDED645E5C523F5534B14AB1B/0?term=3", text.Flag)
		}

		return

	} else {
		//其余搜寻答案即可
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
