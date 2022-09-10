// Package analysis_tool 分析意图
package analysis_tool

import (
	"QQbot/global"
	"math/rand"
	"strings"
)

//whichIntention 判断意图关键词中是否含有特定词语
func whichIntention(str string, keys []string) bool {
	for _, v := range keys {
		if strings.Contains(str, v) {
			return true
		}
	}
	return false
}

// IntentionJudge 通过精简后的语句判断表层意图
// 返回值 表层意图 回答的描述信息
func IntentionJudge(cPtr *global.ChanMsg) string {
	msg := cPtr.Msg

	// 情感倾向
	if whichIntention(msg, global.IntentionKey.LikeKey) {
		return global.LIKE
	}
	//3G故事
	if whichIntention(msg, []string{"3G"}) {
		return global.THREE
	}
	// 学校相关
	if whichIntention(msg, global.IntentionKey.SchoolKey) {
		return global.SCHOOL
	}
	// 其他工作室相关
	if whichIntention(msg, global.IntentionKey.StudioKey) {
		return global.STUDIO
	}
	// 零基础
	if whichIntention(msg, global.IntentionKey.QffFreshmenKey) {
		return global.QffFreshmen
	}
	// 勤奋蜂相关
	if whichIntention(msg, global.IntentionKey.QffKey) {
		// 招新简章
		if whichIntention(msg, global.IntentionKey.QffRecruitKey) {
			return global.QffRecruit
		}
		// 刷不刷人
		if whichIntention(msg, global.IntentionKey.QffStayKey) {
			return global.QffStay
		}
		// 学长学姐
		if whichIntention(msg, global.IntentionKey.QffSeniorStuKey) {
			return global.QffSenior
		}
		return global.QFF
	}
	return global.CHAT
}

// SelectAnswer 搜索出确定的答案
func SelectAnswer(class string) string {
	ansArr := global.AnswerMap[class]
	x := rand.Int() % len(ansArr)
	return ansArr[x]
}
