package dao

import (
	"QQbot/config"
	"QQbot/global"
	"QQbot/tools"
)

//SelectQA 对确定的问题进行模糊查询-fuzzySearch
func SelectQA(msg string) (string, error) {
	var qa []config.QA
	var min = 100 //记录最小值进行比较
	err := global.DB.Find(&qa).Error
	if err != nil {
		return "", err
	}
	var answer string
	for _, v := range qa {
		x := tools.MatchDistance(msg, v.Q1) //计算距离
		if x < min && x >= 0 {              //满足要求
			min = x           //更新最小值
			answer = v.Answer //更新结果
		}
		x = tools.MatchDistance(msg, v.Q2)
		if x < min && x >= 0 {
			min = x
			answer = v.Answer
		}
		x = tools.MatchDistance(msg, v.Q3)
		if x < min && x >= 0 {
			min = x
			answer = v.Answer
		}
	}
	if min == 100 {
		return "我还不知道哦", nil
	}
	//TODO: 设置距离限制，不然乱回答
	return answer, nil
}
