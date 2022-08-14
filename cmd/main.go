package main

import (
	"QQbot/api"
	_ "QQbot/config"
)

func main() {
	//结束的时候导出所有的问答文件
	api.StartEngines()
}
