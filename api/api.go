package api

import (
	"QQbot/service"
	"github.com/gin-gonic/gin"
)

func StartEngines() {
	r := gin.Default()

	r.Any("/", service.PostRespond)
	if err := r.Run(); err != nil {
		panic(err)
	}
}
