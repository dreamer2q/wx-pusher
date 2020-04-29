package config

import (
	"github.com/pelletier/go-toml"
)

type Config struct {
	Wechat struct {
		AppID        string
		AppSecret    string
		AppToken     string
		AesEncodeKey string
		Callback     string
		Timeout      int
		Debug        bool
	}
	Log struct {
		Path  string
		Level string
	}
	Database struct {
		Host   string
		User   string
		Pass   string
		DBName string
	}
}

var config Config
var Wechat = &config.Wechat
var Log = &config.Log
var DB = &config.Database

func Init() {
	loadConfig("config/config.toml")
}

func loadConfig(path string) {
	tree, err := toml.LoadFile(path)
	if err != nil {
		panic(err)
	}
	err = tree.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
}
