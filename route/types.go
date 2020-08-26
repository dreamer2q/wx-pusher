package route

import (
	"github.com/russross/blackfriday/v2"
	"html/template"
	"time"
)

var funMap = template.FuncMap{
	"markdown": markdown,
	"time":     timefmt,
}

func timefmt(t time.Time) string {
	return t.Format("2006 01-02 15:04:05")
}

func markdown(s string) template.HTML {
	outHTML := blackfriday.Run([]byte(s))
	return template.HTML(outHTML)
}

