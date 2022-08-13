package tools

import (
	"fmt"
	"io"
	"net/http"
	"unsafe"
)

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
