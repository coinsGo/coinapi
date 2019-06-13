package coinapi

import (
	"github.com/fanguanghui/coinrpc/rpc"
	"github.com/fanguanghui/coinrpc/usdt"
)

type Tether struct {
	omni *usdt.OmniClient
}

func NewTether(host, user, pass string) Wallet {
	Tether := &Tether{}
	connCfg := rpc.NewConnConfig(host, user, pass)
	Tether.omni = usdt.NewOmniClient(connCfg)
	return Tether
}

func (this Tether) NewAddress(account string) (address string) {
	a := this.omni.GetNewAddress(account)
	address = a
	return address
}

func (this Tether) GetBalance(address string, propertyid uint32) (balance string) {
	b, _ := this.omni.GetBalance(address, propertyid)
	balance = b
	return balance
}

func (this Tether) SendTransaction() (hash string) {
	return ""
}

func (this Tether) GetTransaction() (tx *Tx) {
	return
}

func (this Tether) GetBlockCount() (count uint32) {
	return
}

func (this Tether) GetBlockTransactions() (txids *[]string) {
	return
}
