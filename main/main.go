package main

import (
	"log"

	"github.com/fanguanghui/coinapi"
)

var (
	host = "192.168.199.141:20332"
	user = "niubit"
	pass = "niubit123"
)

func main() {
	factory := coinapi.Factory{}
	usdt := factory.Get("USDT", host, user, pass)
	balance := usdt.GetBalance("mt8Lvqmik6w4ZimnqDb8pKUupYyXFpjBQJ", 2)
	log.Printf("查询余额：%s\n", balance)

}
