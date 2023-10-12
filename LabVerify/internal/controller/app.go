package controller

import (
	"github.com/gin-gonic/gin"
)

type app struct {
	reports map[string][]byte
}

func NewApp() *app {
	return &app{make(map[string][]byte)}
}

func (a app) Routes(r *gin.Engine) {
	r.POST("/check/:userID", a.RunCheck)
	r.GET("/check/:userID", a.DownloadReport)
}
