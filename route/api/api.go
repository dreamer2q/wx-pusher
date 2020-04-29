package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"wxServ/model"
	"wxServ/service/wechat"
	"wxServ/util"
)

func Init(g gin.IRouter) {
	g.POST("/push", push())
	g.GET("/push", push())
}

type pushParam struct {
	Token string `form:"token" json:"token" binding:"required"`
	Msg   string `form:"msg" json:"msg" binding:"required"`
}

func push() gin.HandlerFunc {
	return func(c *gin.Context) {
		pp := pushParam{}
		err := c.Bind(&pp)
		if err != nil {
			util.MakeFailure(c, util.ErrBadPayload, err)
			return
		}
		tk := model.Token{
			Token: pp.Token,
		}
		if err = tk.Load(); err != nil {
			util.MakeFailure(c, util.ErrTokenInvalid, err)
			return
		}
		if msgId, err := wechat.PushMsg(tk.OpenID, &model.PushMsg{
			CreateTime: time.Now(),
			Content:    pp.Msg,
		}); err != nil {
			util.MakeFailure(c, util.ErrInternal, err)
		} else {
			util.MakeSuccess(c, gin.H{
				"msgId": msgId,
				"url":   fmt.Sprintf("http://mjj.dreamer2q.wang/show?id=%v", msgId),
			})
		}
	}
}
