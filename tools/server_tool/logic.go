package server_tool

import (
	"QQbot/global"
	"strings"
)

func RespondLogic(text *global.ChanMsg) error {
	// 消息重复
	if text.Repeated {
		ResPondWithText(text.Id, "刚刚才回答过哦", text.Flag)
	}

	// 私聊消息回复
	if text.Flag == global.PrivateFlag {
		err := privateLogic(text.Id, text.Msg)
		if err != nil {
			return err
		}
	}

	// 群聊消息回复
	if text.Flag == global.GroupFlag {
		err := groupLogic(text.Id, text.Msg)
		if err != nil {
			return err
		}
	}

	return nil
}

//privateLogic 私聊消息回复逻辑
func privateLogic(id int64, msg string) error {
	// 获取帮助
	if IsHelp(msg) {
		ResPondWithText(id, "我只会一点点欸，主要是开发大大太菜了", global.GroupFlag)
	}

	// 去除表情并回复
	msg = GetUsefulMsg(msg)

	//TODO 获取rasa的回复

	return nil
}

//groupLogic 群聊消息回复逻辑
func groupLogic(id int64, msg string) error {
	// 被@了
	if BeAt(msg) {
		// 删除信息中@的部分
		msg = GetUsefulMsg(msg)

		// 获取帮助
		if IsHelp(msg) {
			ResPondWithText(id, "我只会一点点欸，主要是开发大大太菜了", global.GroupFlag)
		}

		//TODO 通过rasa获取答案并回复

		// 没有被@
	} else {
		// 入群打招呼
		if strings.Contains(msg, "大家好") {
			ResPondWithText(id, "欢迎来到极客勤奋蜂的大家庭", global.GroupFlag)
		}

		// 不直接@也有1/10的概率回答此特定的句子
		if DoOrNot(0.1) {
			ResPondWithText(id, "欢迎大家随时问"+global.MyName+"问题哦", global.GroupFlag)
		}
	}
	return nil
}
