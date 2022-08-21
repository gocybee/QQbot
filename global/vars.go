package global

import (
	"github.com/panjf2000/ants/v2"
	"time"
)

const (
	URLTOConfig = "D:/GithubLibrary/QQbot/config/config.yml" //到配置文件的绝对路径

	AT    = "at"
	FACE  = "face"
	SHARE = "share"
	IMAGE = "image"
	POKE  = "poke"

	RepeatLimit = 3
)

var (
	PostQuestionToRasaURL string //后端链接rasa机器人传输问题的接口
	QuestionAnalysisURl   string //后端语义分析接口
	GetRasaAnswerURL      string //后端从rasa获取答案的接口
	SendMsgURL            string //发送QQ信息的接口
	RefuseFileName        string //复读打断消息所需图片的文件名
	RefuseURL             string //复读打断消息所需图片的地址
	MyName                string //qq机器人的自称(名字)
	MYQQID                string //QQ机器人的qq号码
	MaxPoolNumber         int    //连接池最大数量

	PoolNumber int //记录已使用的连接池数量

	TimeLimit time.Duration //配置连接池的最大空闲时间

	Pool *ants.PoolWithFunc //协程池

	Routing = make(map[string]*RoutingMsg) //每一个用户的对话储存为一个Logic协程

	Repeated = make(map[string]*Repeat, 1) // 储存可能是复读的句子 索引为群号或QQ号
	// 其中表情的ID是1-221.
)
