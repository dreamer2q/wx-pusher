package app

import (
	"github.com/gin-gonic/gin"
	"wxServ/config"
	"wxServ/model"
	"wxServ/route"
	"wxServ/service/db"
	"wxServ/service/log"
	"wxServ/service/redis"
	"wxServ/service/wechat"
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
	panic(g.Run(":80"))
}
