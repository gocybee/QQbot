package global

import (
	"sync"
)

// Repeat 描述复读信息
type Repeat struct {
	Flag    string // group or private
	Content string
	Id      int64 // 指向群号
	Times   int   // 重复次数
	sync.Mutex
}

// ChanMsg 被维护的协程和主程序的通信
type ChanMsg struct {
	Id       int64  //信息来源的qq_id
	Msg      string //问题
	Flag     string //group or private
	Repeated bool   //是否触发了复读
	//通过全局信息初始化
	Session   string //此协程对应的rasa语境令牌
	RoutingID string //目标协程的id--目标协程由发送者的qq号注册的
}

// RoutingMsg 描述用户和rasa的通讯方式
type RoutingMsg struct {
	EnvironmentKey string        //对应问答目的（这次问话处于什么环境）-前面起到qff且后不被覆盖则默认问qff相关问题 默认勤奋蜂
	OldQueId       string        //储存上一个问题的rasaID,防止语义重复
	Session        string        //此协程对应的rasa语境令牌
	C              chan *ChanMsg //信息传输
	sync.Mutex
}

// RasaPost 向Rasa发送的消息
type RasaPost struct {
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

// RasaRec 接收的消息
type RasaRec struct {
	RecipientId string `json:"recipient_id"`
	Text        string `json:"text"`
}

// IntentionKeys 意图分类关键词合集
type IntentionKeys struct {
	StudioKey       []string //其他工作室相关
	QffKey          []string //勤奋蜂相关
	QffFreshmenKey  []string //勤奋蜂-零基础相关
	QffStayKey      []string //勤奋蜂-刷人相关
	QffRecruitKey   []string //勤奋蜂-招新相关
	QffSeniorStuKey []string //勤奋蜂-学长学姐相关
	QffExam         []string //勤奋蜂-考核相关
	QffClass        []string //勤奋蜂-上课相关
	SchoolKey       []string //学校相关
	LikeKey         []string //“喜欢”情感倾向相关
}
