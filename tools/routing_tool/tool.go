package routing_tool

import (
	"QQbot/global"
	"QQbot/tools/server_tool"
)

//MaintainRouting 维护协程直到回话结束，第一次则注册此协程
func MaintainRouting(rmPtr *global.ReceivedMsg) error {

	//询问全局是否注册-sender_id为key
	if _, ok := global.Routing[rmPtr.GetSenderIdStr()]; !ok {
		//最大连接数量把控
		if global.PoolNumber < global.MaxPoolNumber {
			registerRouting(rmPtr.GetSenderIdStr())

		} else {
			server_tool.ResPondWithText(rmPtr.GetOppositeIdInt64(), "呀~"+global.MyName+"的脑袋要转不过来了，等一会再来找我聊天嘛", rmPtr.GetGlobalFlag())
		}
	}
	return nil
}

//registerRouting 注册协程时不需要考虑是否为群聊
func registerRouting(senderIdStr string) {
	var x global.RoutingMsg
	x.C = make(chan *global.ChanMsg, 0)

	global.Routing[senderIdStr] = &x

	//执行默认函数--会阻塞一定的时间
	err := global.Pool.Invoke(senderIdStr)
	if err != nil {
		return
	}
}
