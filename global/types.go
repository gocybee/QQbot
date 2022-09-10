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
	OldQueId  string //储存上一个问题的rasaID,防止语义重复
	RoutingID string //目标协程的id
}

// RoutingMsg 描述用户和rasa的通讯方式
type RoutingMsg struct {
	OldQueId string        //储存上一个问题的rasaID,防止语义重复
	Session  string        //此协程对应的rasa语境令牌
	C        chan *ChanMsg //信息传输
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
	StudioKey       []string `yaml:"studio_key"`         //其他工作室相关
	QffKey          []string `yaml:"qff_key"`            //勤奋蜂相关
	QffFreshmenKey  []string `yaml:"qff_freshmen_key"`   //勤奋蜂-零基础相关
	QffStayKey      []string `yaml:"qff_stay_key"`       //勤奋蜂-刷人相关
	QffRecruitKey   []string `yaml:"qff_recruit_key"`    //勤奋蜂-招新相关
	QffSeniorStuKey []string `yaml:"qff_senior_stu_key"` //勤奋蜂-学长学姐相关
	SchoolKey       []string `yaml:"school_key"`         //学校相关
	LikeKey         []string `yaml:"like_key"`           //“喜欢”情感倾向相关
}
