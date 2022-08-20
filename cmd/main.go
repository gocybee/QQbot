package main

import (
	"QQbot/api"
	_ "QQbot/config"
	"QQbot/global"
	_ "QQbot/pool"
)

func main() {
	defer global.Pool.Release()

	api.StartEngines()
}
