//表示数据库的上一级即调用，即处理数据库信息相关的函数

package tools

import (
	"QQbot/dao"
	"QQbot/global"
	"fmt"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"strings"
	"sync"
)

//CalculateAnswer 计算出距离最小的答案
func CalculateAnswer(msg string) (*string, error) {
	//没有信息
	if strings.TrimSpace(msg) == "" {
		t := "有什么事情吗？"
		return &t, nil
	}

	var min = 100 //记录最小值进行比较
	var answer *string

	for _, v := range global.QAs {
		x := matchDistance(msg, v.Q1) //计算距离
		if x < min && x >= 0 {        //满足要求
			min = x            //更新最小值
			answer = &v.Answer //更新结果
		}
		x = matchDistance(msg, v.Q2)
		if x < min && x >= 0 {
			min = x
			answer = &v.Answer
		}
		x = matchDistance(msg, v.Q3)
		if x < min && x >= 0 {
			min = x
			answer = &v.Answer
		}
	}

	//大规模资源
	if strings.Contains(*answer, "resource") {
		url := strings.Split(*answer, "+")
		answer = loadResource(url[1])
	}

	//距离过远则舍弃答案
	if min > global.DistanceLimit {
		//调用AI
		help, err := AIHelp(msg)
		return &help, err
	}
	return answer, nil
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
func Study(msg string) error {
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

//DBError 数据库出错返回
func DBError(id int64, flag string) {
	var status string
	text := "数据库炸了，寄"
	Send(id, &text, flag)
	fmt.Println(status)
}
