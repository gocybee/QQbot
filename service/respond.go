package service

import (
	"QQbot/global"
	"QQbot/tools/routing_tool"
	"QQbot/tools/server_tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func PostRespond(c *gin.Context) {
	// 获取接收到的信息
	var form map[string]interface{}
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err})
		return
	}

	// 心跳检测的回应
	if server_tool.IsHeartBeat(form) {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	// 记录是否出现重复问题
	var repeated = false
	// 复读判断
	if idPtr, ok, flag := server_tool.IsRepeated(form, &repeated); ok {
		server_tool.ResPondWithTextAndPhoto(idPtr, "复读打咩", global.RefuseFileName, global.RefuseURL, flag)
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	var flag string
	if server_tool.IsPrivateMsg(form) {
		flag = global.PrivateFlag
	} else if server_tool.IsGroupMsg(form) {
		flag = global.GroupFlag
	}

	//获取信息的重要部分
	senderId, id, msg, err := server_tool.GetIdAndMsg(form, flag)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err})
		return
	}

	//注册并维护协程--私聊信息或者群聊指定信息
	if server_tool.IsPrivateMsg(form) || (server_tool.IsGroupMsg(form) && server_tool.BeAt(msg)) {
		err = routing_tool.MaintainRouting(senderId, id, msg, repeated, flag)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"err": err})
			return
		}
	} else {
		// 没有被@
		if strings.Contains(msg, "大家好") {
			server_tool.ResPondWithText(id, "欢迎来到极客勤奋蜂的大家庭", global.GroupFlag)
			return
		}

		// 不直接@也有1/10的概率回答此特定的句子
		if server_tool.DoOrNot(0.1) {
			server_tool.ResPondWithText(id, "欢迎大家随时问"+global.MyName+"问题哦", global.GroupFlag)
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{})
	return
}
