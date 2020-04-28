package main

import "time"

func (s *service) initTpl() {
	s.g.FuncMap["time"] = func(t time.Time) string {
		return t.Format("2006 01-02 15:04:05")
	}
	s.g.LoadHTMLGlob("template/*")
}
