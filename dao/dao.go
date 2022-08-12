package dao

import (
	"QQbot/config"
	"QQbot/global"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

//SelectQA 对确定的问题进行模糊查询-fuzzySearch
func SelectQA(msg string) (string, error) {
	var qa []config.QA
	var maxX = 0 //记录最大值进行比较
	var answer string
	err := global.DB.Find(&qa).Error
	if err != nil {
		return "", err
	}
	for _, v := range qa {
		x := fuzzy.RankMatch(msg, v.Q1) //计算匹配度
		if x >= maxX {
			answer = v.Answer
		}
		x = fuzzy.RankMatch(msg, v.Q2) //计算匹配度
		if x >= maxX {
			answer = v.Answer
		}
		x = fuzzy.RankMatch(msg, v.Q3) //计算匹配度
		if x >= maxX {
			answer = v.Answer
		}
	}
	return answer, nil
}
