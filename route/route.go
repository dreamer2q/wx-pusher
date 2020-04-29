package route

import (
	"github.com/gin-gonic/gin"
	"wxServ/route/api"
	"wxServ/route/page"
)

func Init(g *gin.Engine) {
	initTpl(g)
	group := g.Group("")
	page.Init(group)
	v1 := group.Group("/api")
	api.Init(v1)
}

func initTpl(g *gin.Engine) {
	g.SetFuncMap(funMap)
	g.LoadHTMLGlob("template/*.gohtml")
}
