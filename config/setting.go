package config

import (
	"github.com/go-ini/ini"
	"log"
)

var (
	Cfg *ini.File

	HTTPPort     int

)

func init() {

	var err error
	var err1 error

	Cfg, err = ini.Load("config/server.conf")
	if err != nil {
		//如果是test修改测试路径
		Cfg, err1 = ini.Load("../config/server.conf")
		if err1 != nil {
			log.Fatalf("Fail to parse 'config/server.conf': %v", err)
		}
	}

	LoadServer()
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
}
