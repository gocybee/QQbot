// Package analysis_tool 分析意图
package analysis_tool

import (
	"QQbot/global"
	"math/rand"
	"strings"
)

// IntentionJudge 通过精简后的语句判断表层意图
// 返回值 表层意图 回答的描述信息
func IntentionJudge(msg string) string {
	// 首先判断环境
	if ok, flag := isRelatedToQFFQue(msg); ok {
		return flag
	}

	// 情感倾向
	if whichIntention(msg, global.IntentionKey.LikeKey) {
		return global.LIKE
	}

	// 3G故事
	if whichIntention(msg, []string{"3G", "3g"}) {
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

	return global.CHAT
}

// SelectAnswer 搜索出确定的答案
func SelectAnswer(class string) string {
	ansArr := global.AnswerMap[class]
	return ansArr[rand.Intn(len(ansArr))]
}

// whichIntention 判断意图关键词中是否含有特定词语
func whichIntention(str string, keys []string) bool {
	for _, v := range keys {
		if strings.Contains(str, v) {
			return true
		}
	}
	return false
}

// isRelatedToQFFQue 已有环境时，判断是否确实是问qff环境中的内容
func isRelatedToQFFQue(msg string) (bool, string) {
	// 招新简章-qff
	if whichIntention(msg, global.IntentionKey.QffRecruitKey) {
		return true, global.QffRecruit
	}
	// 刷不刷人-qff
	if whichIntention(msg, global.IntentionKey.QffStayKey) {
		return true, global.QffStay
	}
	// 学长学姐-qff
	if whichIntention(msg, global.IntentionKey.QffSeniorStuKey) {
		return true, global.QffSenior
	}
	// 考核-qff
	if whichIntention(msg, global.IntentionKey.QffExam) {
		return true, global.QffExam
	}
	// 授课-qff
	if whichIntention(msg, global.IntentionKey.QffClass) {
		return true, global.QffClass
	}
	return false, ""
}
