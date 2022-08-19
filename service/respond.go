package service

import (
	"QQbot/global"
	"QQbot/tools/dao_tool"
	"QQbot/tools/server_tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func PostRespond(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{}) // 啊对对对
	// 获取接收到的信息
	var form map[string]interface{}
	if c.ShouldBind(&form) != nil {
		return
	}
	// 心跳检测的回应
	if server_tool.IsHeartBeat(form) {
		return
	}

	// 记录是否出现重复问题
	var repeated = false
	// 复读判断
	if idPtr, ok, flag := server_tool.IsRepeated(form, &repeated); ok {
		server_tool.ResPondWithTextAndPhoto(idPtr, "复读打咩", global.RefuseFileName, global.RefuseURL, flag)
		return
	}

	// 私聊消息回复
	if server_tool.IsPrivateMsg(form) {

		idPtr, msgPtr, err := server_tool.GetIdAndMsg(form, global.PrivateFlag)
		if err != nil {
			return
		}
		// 去除表情
		msgPtr = server_tool.GetUsefulMsg(msgPtr)
		// 消息重复
		if repeated {
			server_tool.ResPondWithText(idPtr, "刚刚才回答过哦", global.PrivateFlag)
			return
		}
		// 没有信息
		if server_tool.NeedAsk(msgPtr) {
			server_tool.ResPondWithAsk(idPtr, global.PrivateFlag)
			return
		}

		// 导出问答文件
		if server_tool.NeedSqlFire(msgPtr) {
			// 导出文件
			err = server_tool.ExportSqlMsg()
			if err != nil {
				server_tool.ResPondWithDBError(idPtr, global.PrivateFlag)
				return
			}
			// 导出成功
			server_tool.ResPondWithText(idPtr, "导出成功", global.PrivateFlag)

			return
		}

		// 学习程序触发
		if server_tool.IsStudy(msgPtr) {
			// 信息写入数据库
			err = dao_tool.Study(msgPtr)
			// 数据库出错
			if err != nil {
				server_tool.ResPondWithDBError(idPtr, global.PrivateFlag)
				return
			}

			server_tool.ResPondWithText(idPtr, "学到了", global.PrivateFlag)

			return
		}

		// 正常问答
		server_tool.RespondWhitSqlAndAI(idPtr, msgPtr, global.PrivateFlag)
		return
	}

	// 群聊消息回复
	if server_tool.IsGroupMsg(form) {
		idPtr, msgPtr, err := server_tool.GetIdAndMsg(form, global.GroupFlag)
		if err != nil {
			return
		}

		// 匿名消息判断
		if server_tool.IsAnonymous(form) {
			server_tool.ResPondWithText(idPtr, "开发大大告诉我，匿名的都是坏蛋，你走开", global.GroupFlag)
			return
		}

		// 被@了
		if server_tool.BeAt(msgPtr) {
			// 删除信息中@的部分
			msgPtr = server_tool.GetUsefulMsg(msgPtr)

			// 消息重复
			if repeated {
				server_tool.ResPondWithText(idPtr, "刚刚才回答过哦", global.GroupFlag)
				return
			}

			// 没有信息
			if server_tool.NeedAsk(msgPtr) {
				server_tool.ResPondWithAsk(idPtr, global.GroupFlag)
				return
			}

			// 获取帮助
			if server_tool.IsHelp(msgPtr) {
				server_tool.ResPondWithText(idPtr, "我只会一点点欸，主要是开发大大太菜了", global.GroupFlag)
				return
			}
			// 正常问答
			server_tool.RespondWhitSqlAndAI(idPtr, msgPtr, global.GroupFlag)

			// 没有被@
		} else {
			// 入群打招呼
			if strings.Contains(msgPtr, "大家好") {
				server_tool.ResPondWithText(idPtr, "欢迎来到极客勤奋蜂的大家庭", global.GroupFlag)
				return
			}

			// 不直接@也有1/10的概率回答此特定的句子
			if server_tool.DoOrNot(0.1) {
				server_tool.ResPondWithText(idPtr, "欢迎大家随时问"+global.MyName+"问题哦", global.GroupFlag)
			}
		}
	}
	return
}
