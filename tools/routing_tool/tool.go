package routing_tool

import (
	"QQbot/global"
	"QQbot/tools/server_tool"
	"strconv"
	"time"
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
	//守护进程
	var x global.Logic
	global.Routing[idStr] = &x
	x.Id = idStr
	global.Routing[idStr].Lock()
	defer global.Routing[idStr].Unlock()
	//信息发送
	x.C = make(chan *global.ChanMsg)
	x.C <- text
	//注册并维护协程
	g := generateRoutingToRun(idStr, &(x.C))
	//写入结构体
	x.HandleFunc = &g
	//执行
	err := global.Pool.Submit(*(global.Routing[idStr].HandleFunc))
	if err != nil {
		return
	}
	//调用一次
	global.Routing[idStr].C <- text

}

//generateRoutingToRun 描述逻辑协程
func generateRoutingToRun(idStr string, x *chan *global.ChanMsg) func() {
	return func() {
		var _x = x
		for {
			select {
			case t := <-*_x:
				err := server_tool.RespondLogic(t)
				if err != nil {
					return
				}
			case <-time.After(global.TimeLimit):
				//删除此协程记录
				delete(global.Routing, idStr)
				return
			}
		}
	}
}
