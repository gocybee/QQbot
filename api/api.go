package api

import (
	"QQbot/service"
	"github.com/gin-gonic/gin"
)

func StartEngines() {
	r := gin.Default()

	//反向http接口为跟路由，并且是POST请求
	r.POST("/", service.PostRespond)
	if err := r.Run(); err != nil {
		panic(err)
	}
}
