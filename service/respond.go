package service

import (
	"QQbot/global"
	"QQbot/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
	"unsafe"
)

func PostRespond(c *gin.Context) {
	var form map[string]interface{}
	if c.ShouldBind(&form) != nil {
		return
	}
	fmt.Println(form)

	//心跳检测的回应
	if form["post_type"] == "meta_event" && form["meta_event_type"] == "heartbeat" {
		c.JSONP(http.StatusOK, gin.H{})
		return
	}

	//私聊消息回复
	if form["post_type"] == "message" && form["message_type"] == "private" {
		msg := form["raw_message"].(string)

		text := tools.JudgeAndResp(msg, int64(0), false) //获取回答的语句
		userId := tools.GetIdFromMap(form["user_id"])    //获取对方的QQ号
		status := sendPrivate(userId, text)              //发送信息

		fmt.Println(status)
	}
	//群聊消息回复
	if form["post_type"] == "message" && form["message_type"] == "group" {
		groupId := tools.GetIdFromMap(form["group_id"]) //获取群聊id
		msg := form["raw_message"].(string)             //获取信息本体

		//获取帮助
		if msg == "-help" || msg == "帮助" || strings.Contains(msg, "你能干什么") {

		}

		//入群打招呼
		if strings.Contains(msg, "大家好") {
			text := "欢迎来到极客勤奋蜂的大家庭!"
			status := sendGroup(groupId, text)
			fmt.Println(status)
		}

		//不直接@也有1/10的概率回答问题
		if tools.DoOrNot(0.1) {
			status := sendGroup(groupId, "欢迎大家随时问"+global.MyName+"问题哦")
			//TODO:Question Answer func
			fmt.Println(status)
		}
		c.JSON(http.StatusOK, gin.H{})
	}
	//at机器人回答问题
	if form["post_type"] == "message" && form["message_type"] == "group" && tools.NeedResp(form["raw_message"]) {
		var (
			status string  //消息的状态
			isAnon = false //是否收到了匿名消息
			text   string  //回复内容
		)
		groupId := tools.GetIdFromMap(form["group_id"]) //获取群聊id
		msg := form["raw_message"].(string)             //获取信息本体
		//匿名消息判断
		if form["anonymous"] != "" {
			isAnon = true
		}

		text = tools.JudgeAndResp(msg, tools.GetIdFromMap(form["user_id"]), isAnon) //回答语句获取
		status = sendGroup(groupId, text)

		fmt.Println(status)
	}

	c.JSON(http.StatusOK, gin.H{})
}

//sendPrivate 私发消息
func sendPrivate(qq int64, msg string) string {
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

//sendGroup 群发消息
func sendGroup(gId int64, msg string) string {
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
