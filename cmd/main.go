package main

import (
	"QQbot/api"
	_ "QQbot/config"
	_ "QQbot/pool"
)

func main() {
	api.StartEngines()
}
