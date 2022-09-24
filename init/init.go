package init

import (
	"QQbot/config"
	"QQbot/dao"
	"QQbot/global"
	"QQbot/pool"
)

func init() {
	// config文件读取
	err := config.Init()
	if err != nil {
		panic(err)
	}

	//连接池的初始化
	err = pool.Init()
	if err != nil {
		panic(err)
	}

	//数据库初始化
	err = dao.Init()
	if err != nil {
		panic(err)
	}

	global.PrintVars()
}
