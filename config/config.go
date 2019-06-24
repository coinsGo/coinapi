package config

import (
	"errors"
	"log"
	"path/filepath"
	"runtime"

	"github.com/fanguanghui/coinrpc/rpc"
	"github.com/go-ini/ini"
)

type envdata struct {
	UsdtHost string
	UsdtUser string
	UsdtPass string

	BtcHost string
	BtcUser string
	BtcPass string

	EthHost string
	EthUser string
	EthPass string
}

var Envcfg = &envdata{}

var Usdtcfg rpc.ConnConfig
var Btccfg rpc.ConnConfig
var Ethcfg rpc.ConnConfig

func Setup(env string) {

	file := currentDir() + "/config.ini"
	cfg, err := ini.Load(file)
	if err != nil {
		log.Fatalf("config.Setup, fail to parse '"+file+"': %v", err)
	}
	mapTo(cfg, env, Envcfg)

	Usdtcfg = rpc.ConnConfig{
		Host: Envcfg.UsdtHost,
		User: Envcfg.UsdtUser,
		Pass: Envcfg.UsdtPass,
	}

	Btccfg = rpc.ConnConfig{
		Host: Envcfg.BtcHost,
		User: Envcfg.BtcUser,
		Pass: Envcfg.BtcPass,
	}

	Ethcfg = rpc.ConnConfig{
		Host: Envcfg.EthHost,
		User: Envcfg.EthUser,
		Pass: Envcfg.EthPass,
	}

}

func mapTo(cfg *ini.File, section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}

func currentDir() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	return filepath.Dir(file)
}
