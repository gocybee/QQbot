package global

import (
	"github.com/panjf2000/ants/v2"
	"time"
)

const (
	URLTOOTHERConfig    = "D:/GithubLibrary/gocybee/QQbot/config/source/other_conf.yml"    //到配置文件的绝对路径
	URLTOANALYSISConfig = "D:/GithubLibrary/gocybee/QQbot/config/source/analysis_conf.yml" //到意图配置的路径

	AT    = "at"
	FACE  = "face"
	SHARE = "share"
	IMAGE = "image"
	POKE  = "poke"

	RepeatLimit = 3

	//意图分类

	STUDIO = "studio" //其他工作室相关

	QFF          = "qff"      //本工作室相关
	QFF_FRESHMEN = "freshmen" //零基础
	QFF_Stay     = "stay"     //刷不刷人的问题
	QFF_RECRUIT  = "recruit"  //涉及勤奋蜂招新问题
	QFF_SENIOR   = "senior"   //涉及勤奋蜂关于学姐学长的问题

	SCHOOL = "school" //有关学校的问题

	LIKE = "like" //贴贴

	CHAT = "chat" //正常唠嗑的问题-交给rasa处理

	UNKNOWN = "u" //无法分析的类
)

var (
	RasaURL        string //后端链接rasa机器人传输问题的接口
	SendMsgURL     string //发送QQ信息的接口
	RefuseFileName string //复读打断消息所需图片的文件名
	RefuseURL      string //复读打断消息所需图片的地址
	MyName         string //qq机器人的自称(名字)
	MYQQID         string //QQ机器人的qq号码
	MaxPoolNumber  int    //连接池最大数量

	IntentionKey IntentionKeys //意图关键词合集

	PoolNumber int //记录已使用的连接池数量

	TimeLimit time.Duration //配置连接池的最大空闲时间

	Pool *ants.PoolWithFunc //协程池

	Routing = make(map[string]*RoutingMsg) //每一个用户的对话储存为一个Logic协程

	Repeated = make(map[string]*Repeat, 1) // 储存可能是复读的句子 索引为群号或QQ号

	// 其中表情的ID是1-221.

)
