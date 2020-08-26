package wechat

import (
	"fmt"
	wc "github.com/dreamer2q/go_wechat"
	"github.com/dreamer2q/go_wechat/menu"
	"github.com/jinzhu/gorm"
	"wx-pusher/model"
	"wx-pusher/util"
)

func initMenu() {
	err := wx.Menu.Create(menu.RootMenu{
		Menus: []menu.Item{
			&menu.SubMenu{
				Name: "开始",
				Menus: []menu.Item{
					&menu.ClickMenu{
						Name: "获取token",
						Key:  "get_token",
					},
					&menu.ClickMenu{
						Name: "重新生成token",
						Key:  "regen_token",
					},
					&menu.ClickMenu{
						Name: "注意事项",
						Key:  "token_attention",
					},
				},
			},
			&menu.SubMenu{
				Name: "关于",
				Menus: []menu.Item{
					&menu.ClickMenu{
						Name: "关于我",
						Key:  "get_about",
					},
				},
			},
		},
	})
	if err != nil {
		l.Errorf("wx create menu: %v", err)
	} else {
		l.Info("wx create menu success")
	}
}

func initMenuEv() {
	wx.On("event.CLICK.get_token", func(msg wc.MessageReceive) wc.MessageReply {
		openID := msg.FromUserName
		tk := model.Token{OpenID: openID}
		if err := tk.Load(); err != nil {
			if err == gorm.ErrRecordNotFound {
				return wc.Text{Content: doRegenToken(openID)}
			}
			return wc.Text{Content: fmt.Sprintf("错误： %v", err)}
		}
		return wc.Text{Content: tk.Token}
	})
	wx.On("event.CLICK.regen_token", func(msg wc.MessageReceive) wc.MessageReply {
		return wc.Text{Content: doRegenToken(msg.FromUserName)}
	})
	wx.On("event.CLICK.token_attention", func(msg wc.MessageReceive) wc.MessageReply {
		return wc.Text{Content: tokenAttention}
	})
	wx.On("event.CLICK.get_about", func(msg wc.MessageReceive) wc.MessageReply {
		return wc.Text{Content: wxAbout}
	})
}

func doRegenToken(openID string) string {
	tk := model.Token{OpenID: openID}
	_ = tk.Load()
	tk.Token = util.GenUUID()
	if err := tk.Update(); err != nil {
		return fmt.Sprintf("RegenToken: %v", err)
	}
	return tk.Token
}
