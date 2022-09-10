// Package analysis 分析意图
package analysis

import (
	"QQbot/global"
	"strings"
)

// IntentionJudge 通过精简后的语句判断表层意图
// 返回值 表层意图 回答的描述信息
func IntentionJudge(rmPtr *global.ReceivedMsg) (int, *global.PostMsg) {
	msg := rmPtr.GetMsg()
	pm := global.GeneratePostMsg(rmPtr)

	// 工作室
	if strings.Contains(msg, "工作室") || strings.Contains(msg, "红岩") || strings.Contains(msg, "组织") {
		// 涉及本工作室
		if strings.Contains(msg, "勤奋蜂") || strings.Contains(msg, "我们") || strings.Contains(msg, "咱们") {
			//涉及零基础
			if strings.Contains(msg, "才入门") || strings.Contains(msg, "零基础") || strings.Contains(msg, "没基础") {
				return global.QFF_FRESHMEN, pm
			}
			//涉及刷人
			if strings.Contains(msg, "留几个人") || strings.Contains(msg, "刷人") {
				return global.QFF_RECRUIT, pm
			}
			//涉及招新的问题
			//涉及学长学姐的问题

			return global.UNKNOWN, pm
		}

		return global.STUDIO, pm
	}

	// 涉及学校
	if strings.Contains(msg, "我们学校") || strings.Contains(msg, "重邮") || strings.Contains(msg, "重庆邮电大学") {
		return global.SCHOOL, pm
	}

	// 贴贴
	if strings.Contains(msg, "喜欢你") || strings.Contains(msg, "爱你") {
		return global.LIKE, pm
	}

	// 聊天
	return global.CHAT, pm
}
