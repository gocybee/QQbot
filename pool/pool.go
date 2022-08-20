package pool

import (
	"QQbot/global"
	"github.com/panjf2000/ants/v2"
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
	global.Pool, err = ants.NewPool(global.MaxPoolNumber)
	return err
}
