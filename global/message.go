package global

import (
	"errors"
	"strconv"
	"unicode"
)

// ReceivedMsg 描述截取到的有用信息
type ReceivedMsg struct {
	flag       string //群聊或私聊的global标志
	oppositeId int64  //信息来源的qq号或群号
	senderId   int64  //信息发送方的qq号
	msg        string //接收到的信息或者经过提取后的信息
	repeated   bool   //消息和此人上一个问题是否相同
}

// GetSentenceStruct 从原始信息中获取信息结构体
// 参数 原始信息
// 返回值 信息综合结构体 错误信息
func GetSentenceStruct(form map[string]interface{}) (*ReceivedMsg, error) {
	if form["post_type"] != "message" {
		return &ReceivedMsg{}, errors.New("信息符合要求")
	}
	//获取发送者的qq号
	senderId := int64((form["sender"].(map[string]interface{}))["user_id"].(float64))

	msg := form["raw_message"].(string)

	var flag string    //区分群聊，私聊
	var opposite int64 //信息来源的qq_id
	if opp, ok := form["group_id"]; ok {
		flag = "group"
		opposite = int64(opp.(float64))
	} else if opp, ok := form["user_id"]; ok {
		flag = "private"
		opposite = int64(opp.(float64))
	} else {
		return &ReceivedMsg{}, errors.New("消息错误")
	}

	return &ReceivedMsg{
		flag:       flag,
		oppositeId: opposite,
		senderId:   senderId,
		msg:        msg,
		repeated:   false,
	}, nil
}

// GetGlobalFlag 获取消息来源-群聊或者私聊的flag
func (s *ReceivedMsg) GetGlobalFlag() string {
	return s.flag
}

// GetOppositeIdInt64 获取信息来源的qq号或群号的int64类型
func (s *ReceivedMsg) GetOppositeIdInt64() int64 {
	return s.oppositeId
}

// GetOppositeIdStr 获取信息来源的qq号或群号的string类型
func (s *ReceivedMsg) GetOppositeIdStr() string {
	return strconv.FormatInt(s.oppositeId, 10)
}

// GetSenderIdInt64 获取发送方的qq号的int64类型
func (s *ReceivedMsg) GetSenderIdInt64() int64 {
	return s.oppositeId
}

// GetSenderIdStr 获取发送方的qq号的string类型
func (s *ReceivedMsg) GetSenderIdStr() string {
	return strconv.FormatInt(s.senderId, 10)
}

// GetMsg 获取原始信息
func (s *ReceivedMsg) GetMsg() string {
	return s.msg
}

// ExtractRawMsg 将信息变为rasa机器人可读取的信息
func (s *ReceivedMsg) ExtractRawMsg() {
	s.msg = GetUsefulMsg(s.msg)
}

// IsRepeated 消息是否相同
func (s *ReceivedMsg) IsRepeated() bool {
	return s.repeated
}

// Repeated 标记消息相同
func (s *ReceivedMsg) Repeated() {
	s.repeated = true
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
