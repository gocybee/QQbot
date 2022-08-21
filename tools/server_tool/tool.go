package server_tool

import (
	"QQbot/global"
	"errors"
	"math/rand"
	"strconv"
	"unicode"
)

// GetIdAndMsg 从初始结构体中获取信息和发送者的id以及群号
func GetIdAndMsg(form map[string]interface{}, flag string) (int64, int64, string, error) {
	var id int64
	if form["raw_message"] == nil {
		return 0, 0, "", errors.New("其他操作")
	}
	//获取发送者的qq号
	senderId := (form["sender"].(map[string]interface{}))["user_id"].(float64)
	_senderId := int64(senderId)

	msg := form["raw_message"].(string)
	if flag == "group" {
		id = int64(form["group_id"].(float64)) // 获取群号
	} else if flag == "private" {
		id = int64(form["user_id"].(float64)) // 获取QQ号
	} else {
		return 0, 0, "", errors.New("flag error")
	}
	return _senderId, id, msg, nil
}

// DoOrNot 生成随机数换算为概率--输入小数,现两位
func DoOrNot(p float32) bool {
	i := rand.Int() % 100
	if i < int(p*100) {
		return true
	}
	return false
}

// GetUsefulMsg 删去@自己部分（CQcode部分），获取消息的可被分析部分
func GetUsefulMsg(msg string) string {
	var (
		x [2]int
		//记录事件发生
		x1Changed = false
		x2Changed = false
	)
	res := []rune(msg)

	//删除@以及表情
	for i := 0; i < len(res); i++ {

		if res[i] == '[' {
			x[0] = i
			x1Changed = true
		}
		if res[i] == ']' {
			x[1] = i
			x2Changed = true
		}
		if x1Changed && x2Changed {
			res = []rune(string(res[:x[0]]) + string(res[x[1]+1:]))
			i = -1 //res改变，再次遍历
			x1Changed = false
			x2Changed = false
		}
	}

	//特殊符号处理
	for i := 0; i < len(res); i++ {
		//无法识别的特殊符号和表情
		if !(unicode.IsLetter(res[i]) || unicode.IsNumber(res[i]) || unicode.Is(unicode.Han, res[i])) {
			res = []rune(string(res[:i]) + string(res[i+1:]))
			i -= 1
		}
	}

	return string(res)
}

// Beautify 为句子的尾部美化
func Beautify(ctx *string) {
	// 60%的概率做尾部美化
	if DoOrNot(0.6) {
		i := rand.Int()%221 + 1
		// 避开奇怪的表情
		if (i > 40 && i < 92) || (i > 111 && i < 172) || i > 183 {
			i = 179
		}
		*ctx += global.CodeCQFace(int64(i))
	}
}

// GetPossibleRepeatedMsg 获取可能重复的信息于全局
func GetPossibleRepeatedMsg(idPtr int64, msgPtr string, flag string, happened *bool) (int64, bool, string) {
	_idStr := strconv.FormatInt(idPtr, 10)
	// 找到此人
	if re, ok := global.Re[_idStr]; ok {
		re.Lock()
		defer re.Unlock()

		// 同样的消息
		if re.Content == msgPtr {
			re.Times++
			// 告诉外界重复信息但不构成复读
			*happened = true
			// 触发复读
			if re.Times > global.RepeatLimit {
				// 清除内存
				re.Times = 0
				return idPtr, true, flag
			}
		} else {
			// 消息更新
			re.Content = msgPtr
			re.Times = 1
			return 0, false, ""
		}
		return 0, false, ""
	}
	// 没有找到则创建消息记录
	var r = global.Repeat{Id: idPtr, Content: msgPtr, Flag: flag, Times: 1}
	global.Re[_idStr] = &r
	return 0, false, ""
}
