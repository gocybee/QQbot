package tools

import (
	"QQbot/global"
	"bufio"
	"errors"
	"io"
	"math/rand"
	"os"
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
	var x [2]int
	str := msg.(string)
	res := []rune(str)

	for i := 0; i < len(res); i++ {
		if res[i] == '[' {
			x[0] = i
		}
		if res[i] == ']' {
			x[1] = i
		}
		if x[0] != 0 || x[1] != 0 {
			res = []rune(string(res[:x[0]]) + string(res[x[1]+1:]))
			x[0], x[1] = 0, 0
		}
	}
	an := strings.TrimSpace(string(res))
	return an
}

//Beautify 为句子的头和尾美化
func Beautify(ctx *string) {
	//60%的概率做尾部美化
	if DoOrNot(0.6) {
		i := rand.Int()%221 + 1
		//避开奇怪的表情
		if (i > 40 && i < 92) || (i > 111 && i < 172) || i > 183 {
			i = 179
		}
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

//TODO 错误处理
func loadResource(FileName string) *string {
	file, err := os.OpenFile(global.ResourceURL+FileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var result string
	// 按行处理txt
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		result += string(line)
		result += "%0A"
	}
	return &result

}
