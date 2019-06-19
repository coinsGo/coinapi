package config

import (
	"log"

	"github.com/fanguanghui/coinrpc/rpc"
	"github.com/go-ini/ini"
)

var UsdtCfg = &rpc.ConnConfig{}

var BtcCfg = &rpc.ConnConfig{}

var EthCfg = &rpc.ConnConfig{}

var cfg *ini.File

func Setup(environment string) {
	var err error
	cfg, err = ini.Load("./config/ini/" + environment + ".ini")
	if err != nil {
		log.Fatalf("config.Setup, fail to parse 'config/%v.ini': %v", environment, err)
	}

	mapTo("USDT", UsdtCfg)
	mapTo("BTC", BtcCfg)
	mapTo("ETH", EthCfg)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)

	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
