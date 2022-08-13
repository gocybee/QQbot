package service

import (
	"QQbot/dao"
	"QQbot/global"
	"QQbot/runtime"
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
		text, err := GetRespondWord(msg, int64(0))    //获取回答的语句
		//出问题直接退出
		if err != nil {
			text = "数据库炸了，寄"
			tools.Beautify(&text)
			status := tools.SendPrivate(userId, text) //发送信息
			fmt.Println(status)
			c.JSONP(http.StatusBadRequest, gin.H{})
			return
		}
		status := tools.SendPrivate(userId, text) //发送信息
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
			text = "开发大大告诉我，匿名的都是坏蛋，你走开"
			tools.Beautify(&text)
			status = tools.SendGroup(groupId, text)
			fmt.Println(status)
			c.JSON(http.StatusOK, gin.H{})
			return
		}

		//被@了
		if tools.BeAt(form["raw_message"]) {
			//获取帮助
			if tools.IsHelp(msg) {
				text = "我只会一点点欸，主要是开发大大太菜了" //回答语句获取
				tools.Beautify(&text)
				status = tools.SendGroup(groupId, text)
				fmt.Println(status)
			} else {
				var err error
				text, err = GetRespondWord(msg, tools.GetIdFromMap(form["user_id"])) //回答语句获取
				//出问题直接退出
				if err != nil {
					text = "数据库炸了，寄"
					tools.Beautify(&text)
					status = tools.SendGroup(groupId, text)
					fmt.Println(status)
					c.JSONP(http.StatusBadRequest, gin.H{})
					return
				}
				tools.Beautify(&text)
				status = tools.SendGroup(groupId, text)
				fmt.Println(status)
			}

			//没有被@
		} else {
			//入群打招呼
			if strings.Contains(msg, "大家好") {
				text = "欢迎来到极客勤奋蜂的大家庭!\n欢迎大家随时问" + global.MyName + "问题哦"
				tools.Beautify(&text)
				status = tools.SendGroup(groupId, text)
				fmt.Println(status)
			}

			//不直接@也有1/10的概率回答此特定的句子
			if tools.DoOrNot(0.1) {
				status = tools.SendGroup(groupId, "欢迎大家随时问"+global.MyName+"问题哦")
				tools.Beautify(&text)
				fmt.Println(status)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{})
}

//GetRespondWord 回复消息可能@也可能不@
func GetRespondWord(msg string, uId int64) (string, error) {
	var text string
	//有50%的几率@回去
	if tools.DoOrNot(0.5) {
		text += runtime.CodeCQAt(uId)
	}
	//打招呼
	if strings.Contains(msg, "你好") {
		text += "你好你好鸭"
		return text, nil
	}
	//模糊查询
	t, err := dao.SelectQA(tools.GetUsefulMsg(msg))
	if err != nil {
		return "", err
	}
	text += t
	return text, nil
}
