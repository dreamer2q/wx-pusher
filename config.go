package main

import (
	"flag"
	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
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
}

func loadConfig(path string) (*Config, error) {
	ret := &Config{}
	tree, err := toml.LoadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "load")
	}
	err = tree.Unmarshal(ret)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}
	return ret, nil
}

func mustLoadConfig(path string) *Config {
	c, err := loadConfig(path)
	if err != nil {
		panic(err)
	}
	return c
}

func (s *service) initConfig() {
	path := flag.String("conf", "config/config.toml", "load config from file")
	flag.Parse()
	s.config = mustLoadConfig(*path)
}
