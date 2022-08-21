package rasa_tool

import (
	"QQbot/global"
	"io"
	"net/http"
	"unsafe"
)

// AskRasa 向协程发送信息
func AskRasa(l *global.RoutingMsg, rmPtr *global.ReceivedMsg) {
	t := global.ChanMsg{
		Id:        rmPtr.GetOppositeIdInt64(),
		Msg:       rmPtr.GetMsg(),
		Flag:      rmPtr.GetGlobalFlag(),
		Repeated:  rmPtr.IsRepeated(),
		Session:   l.Session,
		OldQueId:  l.OldQueId,
		RoutingID: rmPtr.GetSenderIdStr(),
	}
	l.C <- &t
}

//PostQuestion 向rasa发送问题
func PostQuestion() string {
	resp, err := http.Get(global.PostQuestionToRasaURL)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return *(*string)(unsafe.Pointer(&data))
}

//GetAnswer 收到rasa的回复
func GetAnswer() string {
	resp, err := http.Get(global.GetRasaAnswerURL)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return *(*string)(unsafe.Pointer(&data))
}
