package config

import (
	"github.com/pelletier/go-toml"
)

type Config struct {
	Api struct {
		Domain string
	}
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
		//Path  string
		Level string
	}
	Database struct {
		Host   string
		User   string
		Pass   string
		DBName string
	}
	Redis struct {
		Host string
	}
}

var config Config
var Wechat = &config.Wechat
var Log = &config.Log
var DB = &config.Database
var Redis = &config.Redis
var Api = &config.Api

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
