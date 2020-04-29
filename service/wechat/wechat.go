package wechat

import (
	"fmt"
	wc "github.com/dreamer2q/go_wechat"
	"github.com/gin-gonic/gin"
	"time"
	"wxServ/config"
	"wxServ/service/log"
)

const (
	tokenAttention = `请保管好token，必要时可以重置token`
	wxAbout        = `关于： 微信推送服务`
)

var wx *wc.API
var l = log.Instance()

func Instance() *wc.API {
	return wx
}

func Init(g *gin.Engine) {
	l = log.Instance()
	r := wc.Config{
		AppID:        config.Wechat.AppID,
		AppSecret:    config.Wechat.AppSecret,
		AppToken:     config.Wechat.AppToken,
		AesEncodeKey: config.Wechat.AesEncodeKey,
		Callback:     config.Wechat.Callback,
		Timeout:      time.Duration(config.Wechat.Timeout) * time.Second,
		Debug:        config.Wechat.Debug,
	}
	wx = wc.New(&r)
	wx.AttachToGin(g)
	wx.SetEventHandler(eventHandler())
	wx.SetMessageHandler(messageHandler())

	initMenu()
	initMenuEv()
}

func messageHandler() wc.Handler {
	return func(msg wc.MessageReceive) wc.MessageReply {
		return nil
	}
}

func eventHandler() wc.Handler {
	return func(msg wc.MessageReceive) wc.MessageReply {
		switch msg.Event {
		case wc.EvSubscribe:
			l.Infof("event: subscribe: %s", msg.FromUserName)
			if info, err := wx.User.GetUserInfo(msg.FromUserName); err == nil {
				return wc.Text{Content: fmt.Sprintf("欢迎关注: %s", info.Nickname)}
			}
			return wc.Text{Content: "欢迎关注"}
		case wc.EvUnsubscribe:
			l.Infof("event: unsubscribe: %s", msg.FromUserName)
			return nil
		}
		return nil
	}
}
