package pool

import (
	"QQbot/global"
	"QQbot/tools/server_tool"
	"time"
)

// RoutingRuntimeLogic 每一个协程的运行逻辑
// 传入的 x 必须为此信息的sender_id!!!!
func RoutingRuntimeLogic(x interface{}) {
	global.PoolNumber.Inc()
	str := x.(string)
	for {
		select {
		case t := <-global.Routing[str].C: // 协程接收到由handleFunction发来的目标信息
			// 接收到消息
			server_tool.RespondLogic(t)

		case <-time.After(time.Duration(global.TimeLimit * int64(time.Second))):
			// 删除此协程记录
			delete(global.Routing, str)
			global.PoolNumber.Dec()
			return
		}
	}
}
