package app

import (
	"github.com/gin-gonic/gin"
	"wx-pusher/config"
	"wx-pusher/model"
	"wx-pusher/route"
	"wx-pusher/service/db"
	"wx-pusher/service/log"
	"wx-pusher/service/redis"
	"wx-pusher/service/wechat"
)

var g *gin.Engine

func Init() {
	g = gin.New()
	g.Use(gin.Logger(), gin.Recovery())

	config.Init()
	log.Init()
	db.Init()
	redis.Init()
	model.Init()
	wechat.Init(g)
	route.Init(g)
}

func Run() {
	Init()
	log.Instance().Info("starting wx server")
	panic(g.Run(":8080"))
}
