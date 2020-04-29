package page

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wxServ/model"
	"wxServ/service/redis"
)

func Init(g gin.IRouter) {
	g.GET("/show", show())
}

func show() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawId := c.Query("id")
		if rawId == "" {
			c.HTML(http.StatusNotFound, "notfound", nil)
			return
		}
		_, err := strconv.ParseInt(rawId, 10, 64)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error", err)
			return
		}
		msg := model.PushMsg{}
		err = redis.Load(rawId, &msg)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error", err)
			return
		}
		c.HTML(http.StatusOK, "show", msg)
	}
}
