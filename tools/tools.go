package tools

import (
	"QQbot/global"
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

//GetIdFromMap 从接受到的表单中提取出用户或者群聊Id
func GetIdFromMap(id interface{}) int64 {
	return int64(id.(float64))
}

//DoOrNot 生成随机数换算为概率--输入小数,现两位，默认0.5
func DoOrNot(p float32) bool {
	i := rand.Int() % 100
	if i < int(p*100) {
		return true
	}
	return false
}

//GetUsefulMsg 删去@自己部分（CQcode部分），获取消息的可被分析部分
func GetUsefulMsg(msg interface{}) string {
	str := msg.(string)
	return strings.TrimFunc(str, func(u rune) bool {
		if string(u) == fmt.Sprintf("[CQ:at,qq=%s]", global.MYQQID) {
			return true
		}
		return false
	})
}

//Beautify 为句子的头和尾美化
func Beautify(ctx *string) {
	//60%的概率做前部美化
	if DoOrNot(0.6) {
		i := rand.Int() % (len(global.Add))
		*ctx = global.Add[i] + *ctx
	}
	//60%的概率做尾部美化
	if DoOrNot(0.6) {
		i := rand.Int()%221 + 1
		*ctx += global.CodeCQFace(int64(i))
	}
}

//CodeQA 将学习的问题包装成结构体
func CodeQA(msg string) (global.QA, error) {
	qa := strings.Split(msg, "+") //0-三个问题，1-答案
	question := strings.Split(qa[0], " ")
	var q [3]string
	//问题初始化
	for i := 0; i < len(question); i++ {
		q[i] = question[i]
	}
	if q[0] == "" || qa[1] == "" {
		return global.QA{}, errors.New("数据读取错误")
	}
	return global.QA{
		Q1:     q[0],
		Q2:     q[1],
		Q3:     q[2],
		Answer: qa[1],
	}, nil
}
