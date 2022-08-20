package link_rasa

import (
	"QQbot/global"
	"io"
	"net/http"
	"unsafe"
)

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
