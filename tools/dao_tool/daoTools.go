//表示数据库的上一级即调用，即处理数据库信息相关的函数

package dao_tool

import (
	"QQbot/dao"
	"QQbot/global"
	"bufio"
	"errors"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"io"
	"os"
	"strings"
	"sync"
)

//CalculateAnswer 计算出距离最小的答案
func CalculateAnswer(msg *string) *string {

	var min = 100 //记录最小值进行比较
	var answer string

	for _, v := range global.QAs {
		x := matchDistance(*msg, v.Q1) //计算距离
		if x < min && x >= 0 {         //满足要求
			min = x           //更新最小值
			answer = v.Answer //更新结果
		}
		x = matchDistance(*msg, v.Q2)
		if x < min && x >= 0 {
			min = x
			answer = v.Answer
		}
		x = matchDistance(*msg, v.Q3)
		if x < min && x >= 0 {
			min = x
			answer = v.Answer
		}
	}

	//answer为空指针则string.Contains()报错
	if strings.Contains(answer, "resource") {
		url := strings.Split(answer, "+")
		answerPtr := loadResource(url[1])
		return answerPtr
	}

	//距离过远则舍弃答案
	if min > global.DistanceLimit {
		return nil
	}

	return &answer
}

//MatchDistance 获取某一个信息距离数据库中问题的距离
func matchDistance(msg string, sql string) int {
	mSlice := SplitMsg(msg)
	var X = -1 //匹配度 0-完全匹配
	for _, v := range mSlice {
		x := fuzzy.RankMatch(v, sql)
		if x < 0 {
			continue
		}
		//出现匹配的树则统计距离
		var once sync.Once
		once.Do(func() { X = 0 })
		X += x
	}
	return X
}

//SplitMsg 将信息拆分成两个字，便于模糊匹配
func SplitMsg(msg string) []string {
	msg += " "
	var res []string
	var m = []rune(msg)
	for i := 0; i < len(m)-2; i++ {
		t := string(m[i]) + string(m[i+1])
		res = append(res, t)
	}
	return res
}

//UpdateQA 作为一个常运行的协程：更新本能地的缓存数据
func UpdateQA() error {
	return dao.SelectQA()
}

//Study 开启学习功能，更新本地数据库
func Study(msg *string) error {
	qaS, err := CodeQA(msg)
	if err != nil {
		return err
	}
	err = dao.AddQA(qaS)
	if err != nil {
		return err
	}
	return UpdateQA()
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

//CodeQA 将学习的问题包装成结构体
func CodeQA(msg *string) (global.QA, error) {
	qa := strings.Split(*msg, "+") //0-三个问题，1-答案
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
