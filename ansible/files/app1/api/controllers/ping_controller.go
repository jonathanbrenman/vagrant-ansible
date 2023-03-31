package controllers

import (
	"github.com/gin-gonic/gin"
)

type PingController interface {
	Ping(context *gin.Context)
}

type pingControllerImpl struct{}

var singletonPingCtrl *pingControllerImpl

func NewPingController() PingController {
	if singletonPingCtrl != nil {
		return singletonPingCtrl
	}

	singletonPingCtrl = &pingControllerImpl{}
	
	return singletonPingCtrl
}

func (p *pingControllerImpl) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":         200,
		"message":      "pong",
	})
}
