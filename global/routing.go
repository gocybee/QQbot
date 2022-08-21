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
