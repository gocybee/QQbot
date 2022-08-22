package rasa_tool

import (
	"QQbot/global"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
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

//GetRasaAnswer 向rasa发送问题并收到回复
func GetRasaAnswer(session string, text string) string {
	q := global.RasaPost{Sender: session, Message: text}
	reader, err := json.Marshal(q)
	if err != nil {
		return err.Error()
	}
	resp, err := http.Post(global.RasaURL+"/webhooks/rest/webhook", "application/json", strings.NewReader(string(reader)))
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}

	var an []map[string]interface{}
	err = json.Unmarshal(data, &an)
	if err != nil {
		return err.Error()
	}

	fmt.Println("原始回答：", *(*string)(unsafe.Pointer(&data)))

	var ans string //储存回答的语句
	for _, v := range an {
		if a, ok := v["text"]; ok {
			ans, err = strconv.Unquote(strings.Replace(strconv.Quote(a.(string)), `\\u`, `\u`, -1))
			if err != nil {
				return err.Error()
			}
		}
	}
	return ans
}

// GetAnalysisId 获取语义分析的id
func GetAnalysisId(text string) string {
	q := fmt.Sprintf("\"text\":%s", text)

	reader, err := json.Marshal(q)
	if err != nil {
		return err.Error()
	}

	resp, err := http.Post("http://0.0.0.0:5005/model/parse", "application/json", strings.NewReader(string(reader)))
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
