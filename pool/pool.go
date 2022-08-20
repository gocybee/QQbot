package pool

import (
	"QQbot/global"
	"QQbot/tools/server_tool"
	"github.com/panjf2000/ants/v2"
	"time"
)

func init() {
	err := loadPool()
	if err != nil {
		panic(err)
	}
}

//loadPool 启动时加载协程池
func loadPool() error {
	var err error
	global.Pool, err = ants.NewPoolWithFunc(global.MaxPoolNumber, func(x interface{}) {
		str := x.(string)
		for {
			select {
			case t := <-global.Routing[str].C:
				err := server_tool.RespondLogic(t)
				if err != nil {
					return
				}
			case <-time.After(global.TimeLimit):
				//删除此协程记录
				delete(global.Routing, str)
				return
			}
		}
	})
	return err
}
