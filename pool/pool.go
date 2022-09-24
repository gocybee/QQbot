package pool

import (
	"QQbot/global"
	"github.com/panjf2000/ants/v2"
)

// Init 启动时加载协程池
func Init() error {
	var err error
	// 每一个协程invoke结束之后都会调用RoutingRuntimeLogic函数
	global.Pool, err = ants.NewPoolWithFunc(global.MaxPoolNumber, RoutingRuntimeLogic)
	return err
}
