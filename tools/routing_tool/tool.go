package routing_tool

import (
	"QQbot/global"
	"QQbot/tools/server_tool"
	"strconv"
)

// MaintainRouting 维护协程直到回话结束，第一次则注册此协程
func MaintainRouting(senderId int64, id int64, msg string, repeated bool, flag string) error {
	_senderIdStr := strconv.FormatInt(senderId, 10)
	// 信息包装
	var t = global.ChanMsg{
		Id:       id,
		Msg:      msg,
		Flag:     flag,
		Repeated: repeated,
	}

	// 逻辑处理
	if l, ok := global.Routing[_senderIdStr]; ok {
		l.Lock()
		defer l.Unlock()
		l.C <- &t
	} else {
		// 最大连接数量把控

		if global.PoolNumber < global.MaxPoolNumber {
			RegisterRouting(_senderIdStr, &t)
		} else {
			server_tool.ResPondWithText(id, "呀~"+global.MyName+"的脑袋要转不过来了，等一会再来找我聊天嘛", flag)
		}
		return nil
	}
	return nil
}

// RegisterRouting 注册协程时不需要考虑是否为群聊
func RegisterRouting(senderIdStr string, text *global.ChanMsg) {

	var x global.Logic
	global.Routing[senderIdStr] = &x

	// 守护进程
	global.Routing[senderIdStr].Lock()
	defer global.Routing[senderIdStr].Unlock()

	// 执行默认函数--会阻塞一定的时间
	err := global.Pool.Invoke(senderIdStr)
	if err != nil {
		return
	}

	// 信息发送
	x.C = make(chan *global.ChanMsg)
	x.C <- text
}
