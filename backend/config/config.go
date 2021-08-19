package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	SessionsSecret []byte
	Redis          redis
	Mysql          mysql
}

type mysql struct {
	Ip   string
	Port int16
	Db   string
}

type redis struct {
	Ip   string
	Port int16
}

var GlobalConfig Config

func ParseConfig() {
	ctx, err := ioutil.ReadFile("config.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(ctx, &GlobalConfig)
	if err != nil {
		return
	}
}
