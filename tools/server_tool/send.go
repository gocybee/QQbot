package server_tool

import (
	"QQbot/global"
	"fmt"
	"io"
	"net/http"
	"unsafe"
)

// send 发小规模消息
func send(qq *int64, msg *string, flag string) string {
	var target string
	if flag == "group" {
		target = "group"
	} else {
		target = "user"
	}
	_url := global.SendMsgURL + "/send_" + flag + "_msg?" + target + "_"
	format := fmt.Sprintf("id=%d&message=%s", *qq, *msg)
	resp, err := http.Get(_url + format)
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
