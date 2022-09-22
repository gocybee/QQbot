package service

import (
	"QQbot/dao"
	"QQbot/global"
	"QQbot/tools/dao_tool"
	"QQbot/tools/routing_tool"
	"QQbot/tools/server_tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostRespond(c *gin.Context) {
	// 获取接收到的信息
	var form map[string]interface{}
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	}

	// 心跳检测的回应
	if server_tool.IsHeartBeat(form) {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	// 生成ReceivedMsg结构体--先判断是否在聊天白名单内-否 则直接return err
	rmPtr, err := global.GetSentenceStruct(form)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	}

	// 不在白名单中则直接退出
	if !(dao.CanChatWith(rmPtr.GetSenderIdStr()) || dao.CanChatWith(rmPtr.GetOppositeIdStr())) {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	// 复读判断
	// 复读当然是继续复读了！
	if server_tool.IsMsgRepeated(rmPtr) {
		server_tool.RespondWithText(rmPtr.GetOppositeIdInt64(), rmPtr.GetMsg(),
			rmPtr.GetGlobalFlag(), false)
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	// 注册并维护协程--私聊信息或者群聊指定信息
	// 维护这轮对话
	if server_tool.IsPrivateMsg(rmPtr.GetGlobalFlag()) || (server_tool.IsGroupMsg(rmPtr.GetGlobalFlag()) && server_tool.BeAt(rmPtr.GetMsg())) {

		// 是否需要 ban 掉回答语句
		if dao_tool.NeedBan(rmPtr.GetSenderIdStr(), rmPtr.GetMsg()) {
			// 获取出 reply 的结构体中的 msg_id
			t := dao_tool.GenerateIdAndAnswerStr(rmPtr.GetMsg(), "")
			err = dao.Banned(t.MsgId)
			if err != nil {
				server_tool.RespondWithText(rmPtr.GetOppositeIdInt64(), "不，我还要说！",
					rmPtr.GetGlobalFlag(), true)
				c.JSON(http.StatusOK, gin.H{"err": err})
				return
			}
			server_tool.RespondWithText(rmPtr.GetOppositeIdInt64(), "知道了知道了",
				rmPtr.GetGlobalFlag(), true)
		}

		// 精简问题--删除多余部分
		rmPtr.ExtractRawMsg()

		// 询问全局并注册
		err = routing_tool.MaintainRouting(rmPtr)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"err": err.Error()})
			return
		}
		// 发送问题
		server_tool.PostChanMsgToRouting(global.Routing[rmPtr.GetSenderIdStr()], rmPtr)
	}

	c.JSON(http.StatusOK, gin.H{})
	return
}
