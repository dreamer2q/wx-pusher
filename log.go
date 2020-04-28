package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
)

func (s *service) initLog() {
	s.log = log.New()
	if runtime.GOOS != "windows" {
		file, err := os.OpenFile(s.config.Log.Path, os.O_APPEND, 0)
		if err != nil {
			log.Panic(err)
		}
		s.log.SetOutput(file)
	}
	if l, err := log.ParseLevel(s.config.Log.Level); err != nil {
		log.Panic(err)
	} else {
		s.log.SetLevel(l)
	}
	s.log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		ForceQuote:             true,
		FullTimestamp:          true,
		TimestampFormat:        "01-02 15:04:05",
		DisableSorting:         true,
		DisableLevelTruncation: false,
		PadLevelText:           false,
		QuoteEmptyFields:       true,
	})
	s.log.Debugf("log init")
}
