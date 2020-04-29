package wechat

import (
	"fmt"
	"github.com/dreamer2q/go_wechat/message"
	"github.com/gin-gonic/gin"
	"wxServ/model"
	"wxServ/service/redis"
)

type tplVal struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

func PushMsg(openID string, msg *model.PushMsg) (string, error) {
	key, err := redis.Store(msg)
	if err != nil {
		return "", err
	}
	_, err = wx.Template.Send(
		&message.TemplateMsg{
			ToUser:     openID,
			TemplateID: "HVDIV2B3Z5HFxVwiQekfSOnMz3Yte02VMYYJdMl7mMA",
			URL:        fmt.Sprintf("mjj.dreamer2q.wang/show?id=%s", key),
			Data: gin.H{
				"time": tplVal{
					Value: msg.CreateTime.Format("2006 01-02 15:04:05"),
					Color: "#173177",
				},
				"msg": tplVal{
					Value: msg.Content,
				},
			},
		})
	return key, err
}