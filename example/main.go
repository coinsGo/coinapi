package main

import (
	"log"

	"github.com/fanguanghui/coinapi/setting"
	"github.com/fanguanghui/coinapi/wallet"
)

func main() {
	factory := wallet.Factory{Environment: setting.Env_DEVELOPMENT}
	usdt := factory.Get("USDT")

	balance := usdt.GetBalance("mt8Lvqmik6w4ZimnqDb8pKUupYyXFpjBQJ")
	log.Printf("查询余额：%s\n", balance)

	//address := usdt.NewAddress("")
	//log.Printf("钱包地址：%s\n", address)

}
