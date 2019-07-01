package wallet

import (
	"github.com/go-develop/coinapi/config"
	"github.com/go-develop/coinrpc/eth"
	"github.com/go-develop/coinrpc/rpc"
)

type Ethereum struct {
	client *eth.EthClient
}

func NewEthereum(connCfg *rpc.ConnConfig) Wallet {
	Ethereum := &Ethereum{}
	Ethereum.client = eth.NewEthClient(connCfg)
	return Ethereum
}

func (this Ethereum) NewAddress(account string) (address string) {
	return
}

func (this Ethereum) SendTransaction(fromaddress, toaddress string, amount, feeaddress string) (txid string) {
	return ""
}

func (this Ethereum) CollectTransaction() (txid string) {
	return ""
}

func (this Ethereum) GetBalance(address string) (balance string) {
	return
}

func (this Ethereum) GetBlockCount() (index uint64) {
	return
}

func (this Ethereum) GetTxById(txid string) (tx config.Tx) {
	return
}

func (this Ethereum) GetBlockTxs(index uint64) (txs []config.Tx) {
	return
}

func (this Ethereum) GetPendingTxs(address string) (txs []config.Tx) {
	return
}
