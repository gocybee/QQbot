package routing_tool

import (
	"QQbot/global"
	"strconv"
)

//MaintainRouting 维护协程直到回话结束，第一次则注册此协程
func MaintainRouting(id int64, msg string, repeated bool, flag string) error {
	_idStr := strconv.FormatInt(id, 10)
	//信息包装
	var t = global.ChanMsg{
		Id:       id,
		Msg:      msg,
		Flag:     flag,
		Repeated: repeated,
	}

	//逻辑处理
	if l, ok := global.Routing[_idStr]; ok {
		l.Lock()
		defer l.Unlock()
		l.C <- &t
	} else {
		RegisterRouting(_idStr, &t)
	}
	return nil
}

func RegisterRouting(idStr string, text *global.ChanMsg) {

	var x global.Logic
	global.Routing[idStr] = &x

	//守护进程
	global.Routing[idStr].Lock()
	defer global.Routing[idStr].Unlock()

	//执行默认函数--会阻塞一定的时间
	err := global.Pool.Invoke(idStr)
	if err != nil {
		return
	}

	//信息发送
	x.C = make(chan *global.ChanMsg)
	x.C <- text
}
