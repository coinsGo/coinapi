package wallet

import (
	"github.com/go-develop/coinapi/config"
	"github.com/go-develop/coinrpc/btc"
	"github.com/go-develop/coinrpc/rpc"
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

func (this Bitcoin) GetBlockCount() (index uint64) {
	return
}

func (this Bitcoin) GetTxById(txid string) (tx config.Tx) {
	return
}

func (this Bitcoin) GetBlockTxs(index uint64) (txs []config.Tx) {
	return
}

func (this Bitcoin) GetPendingTxs(address string) (txs []config.Tx) {
	return
}
