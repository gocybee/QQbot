package service

import (
	"QQbot/global"
	"QQbot/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func PostRespond(c *gin.Context) {
	//获取接收到的信息
	var form map[string]interface{}
	if c.ShouldBind(&form) != nil {
		return
	}

	//心跳检测的回应
	if tools.IsHeartBeat(form) {
		c.JSONP(http.StatusOK, gin.H{})
		return
	}

	//私聊消息回复
	if tools.IsPrivateMsg(form) {
		msg := form["raw_message"].(string)
		userId := tools.GetIdFromMap(form["user_id"]) //获取对方的QQ号
		text := tools.GetRespondWord(msg, int64(0))   //获取回答的语句
		status := tools.SendPrivate(userId, text)     //发送信息
		fmt.Println(status)
	}

	//群聊消息回复
	if tools.IsGroupMsg(form) {
		var status string                               //消息的状态
		var text string                                 //回复内容
		groupId := tools.GetIdFromMap(form["group_id"]) //获取群聊id
		msg := form["raw_message"].(string)             //获取信息本体

		//匿名消息判断
		if tools.IsAnonymous(form) {
			status = tools.SendGroup(groupId, "开发大大告诉我，匿名的都是坏蛋，你走开")
			fmt.Println(status)
			c.JSON(http.StatusOK, gin.H{})
			return
		}

		//被@了
		if tools.BeAt(form["raw_message"]) {
			//获取帮助
			if tools.GetUsefulMsg(msg) == "-help" || tools.GetUsefulMsg(msg) == "帮助" || strings.Contains(msg, "你能干什么") {
				text = "我只会一点点欸，主要是开发大大太菜了" //回答语句获取
				status = tools.SendGroup(groupId, text)
				fmt.Println(status)
			} else {
				text = tools.GetRespondWord(msg, tools.GetIdFromMap(form["user_id"])) //回答语句获取
				status = tools.SendGroup(groupId, text)
				fmt.Println(status)
			}

			//没有被@
		} else {
			//入群打招呼
			if strings.Contains(msg, "大家好") {
				text = "欢迎来到极客勤奋蜂的大家庭!\n欢迎大家随时问" + global.MyName + "问题哦"
				//TODO 美化句子
				status = tools.SendGroup(groupId, text)
				fmt.Println(status)
			}

			//不直接@也有1/10的概率回答问题
			if tools.DoOrNot(0.1) {
				status = tools.SendGroup(groupId, "欢迎大家随时问"+global.MyName+"问题哦")
				//TODO:Question Answer func
				fmt.Println(status)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{})
}
