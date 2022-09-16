package routing_tool

import (
	"QQbot/global"
	"QQbot/tools/server_tool"
	"fmt"
	"sync"
	"time"
)

var maintainLock = sync.Mutex{}

// MaintainRouting 维护协程直到回话结束，第一次则注册此协程
func MaintainRouting(rmPtr *global.ReceivedMsg) error {
	// 询问全局是否注册-sender_id为key
	if _, ok := global.Routing[rmPtr.GetSenderIdStr()]; !ok {
		// 最大连接数量把控
		if int(global.PoolNumber.Load()) < global.MaxPoolNumber {
			maintainLock.Lock()
			defer maintainLock.Unlock()
			registerRouting(rmPtr.GetSenderIdStr())
		} else {
			server_tool.RespondWithText(rmPtr.GetOppositeIdInt64(),
				global.MyName+"不想和你说话",
				rmPtr.GetGlobalFlag(), false)
		}
	}
	return nil
}

// registerRouting 注册协程时不需要考虑是否为群聊
func registerRouting(senderIdStr string) {
	var x global.RoutingMsg
	// 一定要初始化通道
	x.C = make(chan *global.ChanMsg, 0)
	x.Session = fmt.Sprintf("%s%d", senderIdStr, time.Now().Unix())
	global.Routing[senderIdStr] = &x

	// 执行默认函数--会阻塞一定的时间
	err := global.Pool.Invoke(senderIdStr)
	if err != nil {
		return
	}
}
