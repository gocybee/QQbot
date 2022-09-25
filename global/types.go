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
	Id       int64  // 信息来源的qq_id
	Msg      string // 问题
	Flag     string // group or private
	Repeated bool   // 是否触发了复读
	// 通过全局信息初始化
	Session   string // 此协程对应的rasa语境令牌
	RoutingID string // 目标协程的id--目标协程由发送者的qq号注册的
}

// RoutingMsg 描述用户和rasa的通讯方式
type RoutingMsg struct {
	Session string        // 此协程对应的rasa语境令牌
	C       chan *ChanMsg // 信息传输
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

// MysqlMsg 数据库链接的信息
type MysqlMsg struct {
	User     string
	Password string
	Address  string
	DbName   string // 数据库名称
}

//数据库相关结构体

// ChatWhiteListStruct 聊天白名单的基本信息描述
type ChatWhiteListStruct struct {
	Id  int    `gorm:"primaryKey"` //自增主键
	Uid string `gorm:"uid"`        // 允许聊天的对象（qq号或群号）
}

func (c *ChatWhiteListStruct) TableName() string {
	return "chat_ok_list"
}

// BannedAnswerListStruct 禁止这样的回答-需要关键句子触发
type BannedAnswerListStruct struct {
	Id    int    `gorm:"primaryKey"` //自增主键
	Baned string `gorm:"baned"`      // 被禁止的回答
}

func (b *BannedAnswerListStruct) TableName() string {
	return "ans_ban_list"
}

// AnswerAndIdStruct 存放回答的所有信息和对应的消息id
type AnswerAndIdStruct struct {
	Id      int    `gorm:"primaryKey"` //自增主键
	MsgId   string `gorm:"msg_id"`     // 信息的id
	Content string `gorm:"content"`    // 信息的内容
}

func (a *AnswerAndIdStruct) TableName() string {
	return "id_ans"
}
