package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"wxServ/config"
)

var log *logrus.Logger

func Instance() *logrus.Logger {
	return log
}

func Init() {
	log = logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		ForceQuote:             true,
		FullTimestamp:          true,
		TimestampFormat:        "01-02 15:04:05",
		DisableSorting:         true,
		DisableLevelTruncation: false,
		PadLevelText:           false,
		QuoteEmptyFields:       true,
	})
	log.Debugf("log init")
	if runtime.GOOS != "windows" {
		out, err := os.OpenFile(config.Log.Path, os.O_APPEND, 0)
		if err != nil {
			panic(err)
		}
		log.SetOutput(out)
	}
	if level, err := logrus.ParseLevel(config.Log.Level); err != nil {
		panic(err)
	} else {
		log.SetLevel(level)
	}
}
