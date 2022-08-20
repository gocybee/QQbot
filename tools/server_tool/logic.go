package server_tool

import (
	"QQbot/global"
)

func RespondLogic(text *global.ChanMsg) error {

	// 消息重复
	if text.Repeated {
		ResPondWithText(text.Id, "刚刚才回答过哦", text.Flag)
		return nil
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
		ResPondWithText(id, "我只会一点点欸，主要是开发大大太菜了", global.PrivateFlag)
		return nil
	}

	// 去除表情并回复
	msg = GetUsefulMsg(msg)

	//TODO 获取rasa的回复
	ResPondWithText(id, "我收到消息"+msg, global.PrivateFlag)

	return nil
}

//groupLogic 群聊消息回复逻辑
func groupLogic(id int64, msg string) error {
	// 删除信息中@的部分
	msg = GetUsefulMsg(msg)

	// 获取帮助
	if IsHelp(msg) {
		ResPondWithText(id, "我只会一点点欸，主要是开发大大太菜了", global.GroupFlag)
		return nil
	}

	//TODO 通过rasa获取答案并回复
	ResPondWithText(id, "我收到消息"+msg, global.GroupFlag)

	return nil
}
