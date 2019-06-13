package coinapi

import (
	"github.com/fanguanghui/coinrpc/btc"
	"github.com/fanguanghui/coinrpc/rpc"
)

type Bitcoin struct {
	client *btc.BtcClient
}

func NewBitcoin(host, user, pass string) *Bitcoin {
	Bitcoin := &Bitcoin{}
	connCfg := rpc.NewConnConfig(host, user, pass)
	Bitcoin.client = btc.NewBtcClient(connCfg)
	return Bitcoin
}

func (this Bitcoin) NewAddress(account string) (address string) {
	return
}

func (this Bitcoin) GetBalance(address string, propertyid uint32) (balance string) {
	return
}

func (this Bitcoin) SendTransaction() (hash string) {
	return
}

func (this Bitcoin) GetTransaction() (tx *Tx) {
	return
}

func (this Bitcoin) GetBlockCount() (count uint32) {
	return
}

func (this Bitcoin) GetBlockTransactions() (txids *[]string) {
	return
}
