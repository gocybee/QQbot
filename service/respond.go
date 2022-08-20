package service

import (
	"QQbot/global"
	"QQbot/tools/routing_tool"
	"QQbot/tools/server_tool"
	"github.com/gin-gonic/gin"
	"net/http"
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

	id, msg, err := server_tool.GetIdAndMsg(form, flag)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err})
		return
	}

	//注册并维护协程
	err = routing_tool.MaintainRouting(id, msg, repeated, flag)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err})
		return
	}
	return
}
