package tools

import (
	"QQbot/dao"
	"QQbot/runtime"
	"fmt"
	"io"
	"net/http"
	"strings"
	"unsafe"
)

//GetRespondWord 回复消息可能@也可能不@
func GetRespondWord(msg string, uId int64) string {
	var text string
	//有50%的几率@回去
	if DoOrNot(0.5) {
		text += runtime.CodeCQAt(uId)
	}
	//打招呼
	if strings.Contains(msg, "你好") {
		text += "你好你好鸭"
		return text
	}
	//模糊查询
	t, err := dao.SelectQA(GetUsefulMsg(msg))
	if err != nil {
		text += "我还不知道，等待接入AI"
	}
	text += t
	return text
}

//SendPrivate 私发消息
func SendPrivate(qq int64, msg string) string {
	url := fmt.Sprintf("http://127.0.0.1:5700/send_private_msg?user_id=%d&message=%s", qq, msg)
	resp, err := http.Get(url)
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

//SendGroup 群发消息
func SendGroup(gId int64, msg string) string {
	url := fmt.Sprintf("http://127.0.0.1:5700/send_group_msg?group_id=%d&message=%s", gId, msg)
	resp, err := http.Get(url)
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
