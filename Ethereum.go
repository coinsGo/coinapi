package coinapi

import (
	"github.com/fanguanghui/coinrpc/eth"
	"github.com/fanguanghui/coinrpc/rpc"
)

type Ethereum struct {
	client *eth.EthClient
}

func NewEthereum(host, user, pass string) *Ethereum {
	Ethereum := &Ethereum{}
	connCfg := rpc.NewConnConfig(host, user, pass)
	Ethereum.client = eth.NewEthClient(connCfg)
	return Ethereum
}

func (this Ethereum) NewAddress(account string) (address string) {
	return
}

func (this Ethereum) GetBalance(address string, propertyid uint32) (balance string) {
	return
}

func (this Ethereum) SendTransaction() (hash string) {
	return
}

func (this Ethereum) GetTransaction() (tx *Tx) {
	return
}

func (this Ethereum) GetBlockCount() (count uint32) {
	return
}

func (this Ethereum) GetBlockTransactions() (txids *[]string) {
	return
}
