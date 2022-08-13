//表示数据库的上一级即调用，即处理数据库信息相关的函数

package tools

import (
	"QQbot/dao"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"sync"
)

//CalculateAnswer 计算出距离最小的答案
func CalculateAnswer(msg string) (string, error) {
	qa, err := dao.SelectQA()
	if err != nil {
		return "数据库炸了，寄", err
	}

	var min = 100 //记录最小值进行比较
	var answer string

	for _, v := range *qa {
		x := matchDistance(msg, v.Q1) //计算距离
		if x < min && x >= 0 {        //满足要求
			min = x           //更新最小值
			answer = v.Answer //更新结果
		}
		x = matchDistance(msg, v.Q2)
		if x < min && x >= 0 {
			min = x
			answer = v.Answer
		}
		x = matchDistance(msg, v.Q3)
		if x < min && x >= 0 {
			min = x
			answer = v.Answer
		}
	}
	if min == 100 {
		return "我还不知道哦[CQ:face,id=1]", nil
	}
	//TODO: 设置距离限制，不然乱回答
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
