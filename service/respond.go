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
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	//私聊消息回复
	if tools.IsPrivateMsg(form) {
		msg := form["raw_message"].(string)
		userId := tools.GetIdFromMap(form["user_id"]) //获取对方的QQ号
		//没有信息
		if strings.TrimSpace(msg) == "" {
			t := "有什么事情吗？"
			status := tools.Send(userId, &t, "private") //发送信息
			fmt.Println(status)
			c.JSON(http.StatusOK, gin.H{})
			return
		}

		//导出问答文件
		if msg == "导出问答文件" {
			err := tools.ExportSqlMsg()
			if err != nil {
				tools.DBError(userId, "private")
			}
			t := "导出成功"
			status := tools.Send(userId, &t, "private") //发送信息
			fmt.Println(status)
			c.JSON(http.StatusOK, gin.H{})
			return
		}

		//学习程序触发
		if tools.IsStudy(msg) {
			err := tools.Study(msg)
			//数据库出错
			if err != nil {
				tools.DBError(userId, "private")
				c.JSON(http.StatusBadRequest, gin.H{})
				return
			}
			t := "已经学到啦"
			status := tools.Send(userId, &t, "private") //发送信息
			fmt.Println(status)
			c.JSON(http.StatusOK, gin.H{})
			return
		}

		//正常问答
		text, err := tools.GetRespondWord(msg, int64(0)) //获取回答的语句
		//出问题直接退出
		if err != nil {
			tools.DBError(userId, "private")
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		status := tools.Send(userId, text, "private") //发送信息
		fmt.Println(status)
		c.JSON(http.StatusOK, gin.H{})
	}

	//群聊消息回复
	if tools.IsGroupMsg(form) {
		var status string                               //消息的状态
		var text *string                                //回复内容
		groupId := tools.GetIdFromMap(form["group_id"]) //获取群聊id
		msg := form["raw_message"].(string)             //获取信息本体

		//匿名消息判断
		if tools.IsAnonymous(form) {
			t := "开发大大告诉我，匿名的都是坏蛋，你走开"
			text = &t
			tools.Beautify(text)
			status = tools.Send(groupId, text, "group")
			fmt.Println(status)
			c.JSON(http.StatusOK, gin.H{})
			return
		}

		//被@了
		if tools.BeAt(form["raw_message"]) {
			//获取帮助
			if tools.IsHelp(msg) {
				t := "我只会一点点欸，主要是开发大大太菜了"
				text = &t
				tools.Beautify(text)
				status = tools.Send(groupId, text, "group")
				fmt.Println(status)
			} else {
				var err error
				text, err = tools.GetRespondWord(msg, tools.GetIdFromMap(form["user_id"])) //回答语句获取
				//出问题直接退出
				if err != nil {
					tools.DBError(groupId, "group")
					c.JSON(http.StatusBadRequest, gin.H{})
					return
				}
				tools.Beautify(text)
				status = tools.Send(groupId, text, "group")
				fmt.Println(status)
			}
			//没有被@
		} else {
			//入群打招呼
			if strings.Contains(msg, "大家好") {
				t := "欢迎来到极客勤奋蜂的大家庭"
				text = &t
				tools.Beautify(text)
				status = tools.Send(groupId, text, "group")
				fmt.Println(status)
			}

			//不直接@也有1/10的概率回答此特定的句子
			if tools.DoOrNot(0.1) {
				t := "欢迎大家随时问" + global.MyName + "问题哦"
				status = tools.Send(groupId, &t, "group")
				tools.Beautify(text)
				fmt.Println(status)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{})
}
