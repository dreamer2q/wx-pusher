package main

import (
	wechat "github.com/dreamer2q/go_wechat"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type service struct {
	wx *wechat.API
	g  *gin.Engine

	log    *log.Logger
	config *Config
}

func (s *service) init() {
	s.initConfig()
	s.initLog()
	s.initWx()
	s.initRoute()
}

func (s *service) run() {
	s.init()

	log.Panic(s.g.Run(":80"))
}
