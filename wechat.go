package main

import (
	"encoding/base64"
	"fmt"
	wc "github.com/dreamer2q/go_wechat"
	"github.com/dreamer2q/go_wechat/menu"
	"time"
)

func (s *service) initWx() {
	c := wc.Config{
		AppID:        s.config.Wechat.AppID,
		AppSecret:    s.config.Wechat.AppSecret,
		AppToken:     s.config.Wechat.AppToken,
		AesEncodeKey: s.config.Wechat.AesEncodeKey,
		Callback:     "/wx",
		Timeout:      time.Duration(s.config.Wechat.Timeout) * time.Second,
		Debug:        s.config.Wechat.Debug,
	}
	s.wx = wc.New(&c)
	s.initMenu()
	s.wx.SetMessageHandler(s.MsgHandler())
	s.wx.SetEventHandler(s.EventHandler())
}

func (s *service) initMenu() {
	err := s.wx.Menu.Create(menu.RootMenu{
		Menus: []menu.Item{
			&menu.SubMenu{
				Name: "开始",
				Menus: []menu.Item{
					&menu.ClickMenu{
						Name: "获取token",
						Key:  "menu_click_get_token",
					},
					&menu.ClickMenu{
						Name: "注意事项",
						Key:  "menu_click_token_attention",
					},
				},
			},
			&menu.SubMenu{
				Name: "关于",
				Menus: []menu.Item{
					&menu.ClickMenu{
						Name: "关于我",
						Key:  "menu_click_get_about",
					},
				},
			},
		},
	})
	if err != nil {
		s.log.Errorf("wx create menu: %v", err)
	} else {
		s.log.Info("wx create menu success")
	}
}

func (s *service) EventHandler() wc.Handler {
	return func(msg wc.MessageReceive) wc.MessageReply {
		switch msg.Event {
		case wc.EvSubscribe:
			s.log.Infof("event: subscribe: %s", msg.FromUserName)
			if info, err := s.wx.User.GetUserInfo(msg.FromUserName); err == nil {
				return wc.Text{Content: fmt.Sprintf("欢迎关注: %s", info.Nickname)}
			}
			return wc.Text{Content: "欢迎关注"}
		case wc.EvUnsubscribe:
			s.log.Infof("event: unsubscribe: %s", msg.FromUserName)
			return nil
		case wc.EvClick:
			switch msg.EventKey {
			case "menu_click_get_token":
				return wc.Text{Content: base64.StdEncoding.EncodeToString([]byte(msg.FromUserName))}
			case "menu_click_token_attention":
				return wc.Text{Content: "非常尴尬，token与账号相关，一般情况是不会改变来的，请注意保管好token"}
			case "menu_click_get_about":
				return wc.Text{Content: "测试推送公众号"}
			}
		}
		return nil
	}
}

func (s *service) MsgHandler() wc.Handler {
	return func(msg wc.MessageReceive) wc.MessageReply {
		return nil
	}
}
