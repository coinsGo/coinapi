package wallet

import (
	"github.com/fanguanghui/coinapi/setting"
	"github.com/fanguanghui/coinrpc/eth"
	"github.com/fanguanghui/coinrpc/rpc"
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

func (this Ethereum) GetTransaction() (tx *setting.Tx) {
	return
}

func (this Ethereum) GetBlockCount() (count uint32) {
	return
}

func (this Ethereum) GetBlockTxids() (txids *[]string) {
	return
}
