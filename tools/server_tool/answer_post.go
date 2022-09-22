// package tools 包含所有直接被service引用的函数

package server_tool

import (
	"QQbot/dao"
	"QQbot/global"
	"QQbot/tools/dao_tool"
	"fmt"
	"io"
	"net/http"
	"unsafe"
)

// RespondWithText 返回test信息
func RespondWithText(id int64, msg string, flag string, needBeautify bool) {
	if needBeautify {
		Beautify(&msg)
	}
	status := send(id, msg, flag)
	fmt.Println(status)
}

// RespondWithPhoto 返回非闪照的图片
func RespondWithPhoto(id int64, fileName string, url string, flag string) {
	msg := global.CodeCQPhoto(fileName, url)
	status := send(id, msg, flag)
	fmt.Println(status)
}

// RespondWithTextAndPhoto 返回信息及非闪照的图片
func RespondWithTextAndPhoto(id int64, msg string, fileName string, url string, flag string) {
	msg += global.CodeCQPhoto(fileName, url)
	status := send(id, msg, flag)
	fmt.Println(status)
}

// send 发消息
func send(qq int64, msg string, flag string) string {
	var target string
	if flag == "group" {
		target = "group"
	} else {
		target = "user"
	}
	url := fmt.Sprintf("%s/send_%s_msg?%s_id=%d&message=%s",
		global.SendMsgURL, flag, target, qq, msg)
	// _url := global.SendMsgURL + "/send_" + flag + "_msg?" + target + "_"
	// format := fmt.Sprintf("id=%d&message=%s", qq, msg)
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	an := *(*string)(unsafe.Pointer(&data))

	// 记录id和信息
	temp := dao_tool.GenerateIdAndAnswerStr(an, msg)
	err = dao.WritIdAndAnswer(temp)
	if err != nil {
		an += err.Error()
	}

	return an
}
