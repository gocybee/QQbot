package dao

import (
	"QQbot/config"
	"QQbot/global"
	"QQbot/tools"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	msg := "觉得这门语言好难学哦"
	var qa []config.QA
	err := global.DB.Find(&qa).Error
	if err != nil {
		panic(err)
	}
	for _, v := range qa {
		x := tools.MatchDistance(msg, v.Q1)
		fmt.Printf("与 %s 的匹配度为%d\n", v.Q1, x)
		x = tools.MatchDistance(msg, v.Q2)
		fmt.Printf("与 %s 的匹配度为%d\n", v.Q2, x)
		x = tools.MatchDistance(msg, v.Q3)
		fmt.Printf("与 %s 的匹配度为%d\n", v.Q3, x)
	}
}
