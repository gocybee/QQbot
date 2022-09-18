package server_tool

import (
	"QQbot/global"
	"QQbot/tools/rasa_tool"
)

// RespondLogic 回答问题
func RespondLogic(text *global.ChanMsg) {

	// 消息重复
	if text.Repeated {
		RespondWithText(text.Id, "你刚刚发过了哦", text.Flag, true)
		return
	}

	if text.Msg == "" || PunctualOnly(text.Msg) {
		RespondWithText(text.Id, "干嘛？", text.Flag, true)
		return
	}

	answer, err := rasa_tool.GetRasaAnswer(text.Session, text.Msg)
	if err != nil || answer == "" {
		// rasa机器人无法回答
		RespondWithText(text.Id, "我不到啊", text.Flag, true)

	} else {
		// 获取到回答的消息
		RespondWithText(text.Id, answer, text.Flag, false)
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
