package global

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"go.uber.org/atomic"
)

const (
	URLTOOTHERConfig    = "./config/other_conf.yml"    // 到配置文件的路径
	URLTOANALYSISConfig = "./config/analysis_conf.yml" // 到意图配置的路径
	URLTOANSWERConfig   = "./config/story_conf.yml"    // 到回答配置文件的路径

	AT    = "at"
	FACE  = "face"
	SHARE = "share"
	IMAGE = "image"
	POKE  = "poke"

	RepeatLimit = 3

	// 意图分类

	STUDIO = "studio" // 其他工作室相关

	QFF         = "qff"      // 本工作室相关
	QffFreshmen = "freshmen" // 零基础
	QffStay     = "stay"     // 刷不刷人的问题
	QffRecruit  = "recruit"  // 涉及勤奋蜂招新问题
	QffSenior   = "senior"   // 涉及勤奋蜂关于学姐学长的问题
	QffExam     = "exam"     // 考核
	QffClass    = "class"    // 上课

	SCHOOL = "school" // 有关学校的问题

	LIKE = "like" // 贴贴

	CHAT = "chat" // 正常唠嗑的问题-交给rasa处理

	THREE = "3G" // 三G的故事
)

var (
	RasaURL       string // 后端链接rasa机器人传输问题的接口
	SendMsgURL    string // 发送QQ信息的接口
	MyName        string // qq机器人的自称(名字)
	MYQQID        string // QQ机器人的qq号码
	MaxPoolNumber int    // 连接池最大数量
	TimeLimit     int64  // 配置对话最长保持时间(s 单位)

	IntentionKey IntentionKeys                  // 意图关键词合集
	AnswerMap    = make(map[string][]string, 1) // 答案储存

	PoolNumber atomic.Int32 // 记录已使用的连接池数量

	Pool *ants.PoolWithFunc // 协程池

	Routing = make(map[string]*RoutingMsg) // 每一个用户的对话储存为一个Logic协程

	Repeated = make(map[string]*Repeat, 1) // 储存可能是复读的句子 索引为群号或QQ号

	// 其中表情的ID是1-221.
)

func PrintVars() {
	fmt.Println()
	fmt.Println("=====配置选项如下=====")
	fmt.Println("RasaURL:", RasaURL)
	fmt.Println("SendMsgURL:", SendMsgURL)
	fmt.Println("BotName:", MyName)
	fmt.Println("BotQQNumber:", MYQQID)
	fmt.Println("MaxPoolNumber:", MaxPoolNumber)
	fmt.Println("TimeLimit:", TimeLimit)
	fmt.Println()
}
