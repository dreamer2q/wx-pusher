package main

import (
	"encoding/base64"
	"fmt"
	"github.com/dreamer2q/go_wechat/message"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type PushMsg struct {
	CreateTime time.Time
	Content    string
}

var MsgStore map[int32]*PushMsg

func (s *service) initRoute() {
	MsgStore = make(map[int32]*PushMsg)

	s.g = gin.New()
	s.g.Use(gin.Logger(), gin.Recovery())
	s.wx.AttachToGin(s.g)
	s.initTpl()

	g := s.g.Group("")

	//page group
	{
		g.GET("/show", s.Show())
	}
	//api group
	v1 := g.Group("/api")
	{
		v1.GET("/send", s.Send())
		v1.POST("/send", s.Send())
	}
}

func (s *service) Show() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("msg")
		if id == "" {
			c.HTML(http.StatusNotFound, "notfound", nil)
			return
		}
		storedId, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error", err)
			return
		}
		if msg, ok := MsgStore[int32(storedId)]; !ok {
			c.HTML(http.StatusNotFound, "notfound", nil)
		} else {
			c.HTML(http.StatusOK, "show", msg)
		}
		return
	}
}

type tplVal struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

func (s *service) Send() gin.HandlerFunc {
	return func(c *gin.Context) {
		qr := struct {
			Token string `form:"token" json:"token" binding:"required"`
			Msg   string `form:"msg" json:"msg" binding:"required"`
		}{}
		if err := c.ShouldBind(&qr); err != nil {
			MakeFailure(c, ErrBadPayload, err)
			return
		}
		if tb, err := base64.StdEncoding.DecodeString(qr.Token); err != nil {
			MakeFailure(c, ErrBadPayload)
			return
		} else {
			qr.Token = string(tb)
		}
		userInfo, err := s.wx.User.GetUserInfo(qr.Token)
		if err != nil {
			MakeFailure(c, ErrTokenInvalid)
			return
		}
		var storeId = rand.Int31()
		if msgId, err := s.wx.Template.Send(
			&message.TemplateMsg{
				ToUser:     qr.Token,
				TemplateID: "HVDIV2B3Z5HFxVwiQekfSOnMz3Yte02VMYYJdMl7mMA",
				URL:        fmt.Sprintf("mjj.dreamer2q.wang/show?msg=%d", storeId),
				Data: gin.H{
					"time": tplVal{
						Value: time.Now().Format("2006-01-02 15:04:05"),
						Color: "#173177",
					},
					"msg": tplVal{
						Value: qr.Msg,
					},
				},
			}); err != nil {
			MakeFailure(c, ErrInternal, err)
			return
		} else {
			MsgStore[storeId] = &PushMsg{
				CreateTime: time.Now(),
				Content:    qr.Msg,
			}
			MakeSuccess(c, gin.H{
				"msgId": msgId,
				"user":  userInfo.Nickname,
			})
		}
		return
	}
}
