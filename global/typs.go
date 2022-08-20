package global

import (
	"sync"
)

//Repeat 描述复读信息
type Repeat struct {
	Flag    string // group or private
	Content string
	Id      int64 // 指向群号
	Times   int   // 重复次数
	sync.Mutex
}

//ChanMsg 被维护的协程和主程序的通信
type ChanMsg struct {
	Id       int64  //发送方的id
	Msg      string //问题
	Flag     string //group or private
	Repeated bool   //是否触发了复读
}

//Logic 描述用户的通讯方式
type Logic struct {
	C chan *ChanMsg //信息传输
	sync.Mutex
}
