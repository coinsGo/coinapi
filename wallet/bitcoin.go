package wallet

import (
	"github.com/fanguanghui/coinapi/setting"
	"github.com/fanguanghui/coinrpc/btc"
	"github.com/fanguanghui/coinrpc/rpc"
)

type Bitcoin struct {
	client *btc.BtcClient
}

func NewBitcoin(connCfg *rpc.ConnConfig) Wallet {
	Bitcoin := &Bitcoin{}
	Bitcoin.client = btc.NewBtcClient(connCfg)
	return Bitcoin
}

func (this Bitcoin) NewAddress(account string) (address string) {
	return
}

func (this Bitcoin) SendTransaction(fromaddress, toaddress string, amount, feeaddress string) (txid string) {
	return ""
}

func (this Bitcoin) CollectTransaction() (txid string) {
	return ""
}

func (this Bitcoin) GetBalance(address string) (balance string) {
	return
}

func (this Bitcoin) GetTransaction() (tx *setting.Tx) {
	return
}

func (this Bitcoin) GetBlockCount() (count uint32) {
	return
}

func (this Bitcoin) GetBlockTxids() (txids *[]string) {
	return
}
