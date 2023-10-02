package token

import (
	"github.com/tencent-connect/botgo/log"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	AppID uint64 `yaml:"appid"`
	Token string `yaml:"token"`
}

var TConfig Config

func LoadTokenConfig() {
	content, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Errorf("read token from file failed, err: %v", err)
		os.Exit(1)
	}
	if err = yaml.Unmarshal(content, &TConfig); err != nil {
		log.Errorf("parse config failed, err: %v", err)
		os.Exit(1)
	}
}
