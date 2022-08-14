//package tools 包含所有直接被service引用的函数

package tools

import (
	"QQbot/global"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"unsafe"
)

//Send 发小规模消息
func Send(qq int64, msg *string, flag string) string {
	var target string
	if flag == "group" {
		target = "group"
	} else {
		target = "user"
	}
	url := fmt.Sprintf(global.SendMsgURL+"/send_"+flag+"_msg?"+target+"_id=%d&message=%s", qq, *msg)
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

//SendHuge 发大量消息
//func SendHuge(gId int64, form *string, flag string) string {
//	url := global.SendMsgURL + "/send_" + flag + "_msg"
//	resp, err := http.Post(url, form)
//	if err != nil {
//		return err.Error()
//	}
//	defer resp.Body.Close()
//
//	data, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return err.Error()
//	}
//	return *(*string)(unsafe.Pointer(&data))
//}

//IsHeartBeat 判断是否为心跳事件
func IsHeartBeat(form map[string]interface{}) bool {
	if form["post_type"] == "meta_event" && form["meta_event_type"] == "heartbeat" {
		return true
	}
	return false
}

//IsPrivateMsg 判断是否为私聊消息
func IsPrivateMsg(form map[string]interface{}) bool {
	if form["post_type"] == "message" && form["message_type"] == "private" {
		return true
	}
	return false
}

//IsGroupMsg 判断是否为群消息
func IsGroupMsg(form map[string]interface{}) bool {
	if form["post_type"] == "message" && form["message_type"] == "group" {
		return true
	}
	return false
}

//IsAnonymous 是否为匿名消息
func IsAnonymous(form map[string]interface{}) bool {
	if form["anonymous"] != nil {
		return true
	}
	return false
}

//IsHelp 是否为帮助
func IsHelp(msg string) bool {
	if GetUsefulMsg(msg) == "-help" || GetUsefulMsg(msg) == "帮助" || strings.Contains(msg, "你能干什么") {
		return true
	}
	return false
}

//IsStudy 是否触发学习程序
func IsStudy(msg string) bool {
	if strings.Contains(msg, "+") {
		return true
	}
	return false
}

//BeAt 看自己是否被@
func BeAt(str interface{}) bool {
	msg := str.(string)
	return strings.Contains(msg, "at") && strings.Contains(msg, global.MYQQID)
}

//GetRespondWord 回复消息可能@也可能不@
func GetRespondWord(msg string, uId int64) (*string, error) {
	var text string
	//有50%的几率@回去
	if DoOrNot(0.5) {
		text += global.CodeCQAt(uId)
	}
	//打招呼
	if strings.Contains(msg, "你好") {
		text += "你好你好鸭"
		return &text, nil
	}
	//模糊查询
	t, err := CalculateAnswer(GetUsefulMsg(msg))
	if err != nil {
		return nil, err
	}
	text += *t
	return &text, nil
}

//ExportSqlMsg 导出所有的学习的信息
func ExportSqlMsg() error {
	data, err := yaml.Marshal(global.QAs)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("oldMsg.yml", data, 0777)
}

//AIHelp 获取AI帮助
func AIHelp(msg string) (string, error) {
	url := fmt.Sprintf("http://api.qingyunke.com/api.php?key=free&appid=0&msg=%s", msg)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var AIType global.AI
	err = json.Unmarshal(data, &AIType)
	if err != nil {
		return "", err
	}
	return AIType.Content, nil
}
