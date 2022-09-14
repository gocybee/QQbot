// package tools 包含所有直接被service引用的函数

package server_tool

import (
	"QQbot/global"
	"fmt"
	"io"
	"net/http"
	"unsafe"
)

// ResPondWithText 返回test信息
func ResPondWithText(id int64, msg string, flag string, needBeautify bool) {
	if needBeautify {
		Beautify(&msg)
	}
	status := send(&id, &msg, flag)
	fmt.Println(status)
}

// ResPondWithPhoto 返回非闪照的图片
func ResPondWithPhoto(id int64, fileName string, url string, flag string) {
	msg := global.CodeCQPhoto(fileName, url)
	status := send(&id, &msg, flag)
	fmt.Println(status)
}

// ResPondWithTextAndPhoto 返回信息及非闪照的图片
func ResPondWithTextAndPhoto(id int64, msg string, fileName string, url string, flag string) {
	msg += global.CodeCQPhoto(fileName, url)
	status := send(&id, &msg, flag)
	fmt.Println(status)
}

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
