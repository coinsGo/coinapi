package wallet

import (
	"github.com/fanguanghui/coinapi/setting"
	"github.com/fanguanghui/coinrpc/rpc"
	"github.com/fanguanghui/coinrpc/usdt"
)

type Tether struct {
	omni       *usdt.OmniClient
	propertyid uint32
	ecosystem  uint32
}

func NewTether(connCfg *rpc.ConnConfig, environment string) Wallet {
	Tether := &Tether{}
	Tether.omni = usdt.NewOmniClient(connCfg)

	if environment == setting.Env_AMAZON_AWS {
		Tether.propertyid = setting.Propertyid_PROD
		Tether.ecosystem = setting.Ecosystem_PROD
	} else {
		Tether.propertyid = setting.Propertyid_TEST
		Tether.ecosystem = setting.Ecosystem_TEST
	}
	return Tether
}

func (this Tether) NewAddress(account string) (address string) {
	a := this.omni.GetNewAddress(account)
	address = a
	return address
}

func (this Tether) SendTransaction(fromaddress, toaddress string, amount, feeaddress string) (txid string) {
	return ""
}

func (this Tether) CollectTransaction() (txid string) {
	return ""
}

func (this Tether) GetBalance(address string) (balance string) {
	b, _ := this.omni.GetBalance(address, this.propertyid)
	balance = b
	return balance
}

func (this Tether) GetTransaction() (tx *setting.Tx) {
	return
}

func (this Tether) GetBlockCount() (count uint32) {
	return
}

func (this Tether) GetBlockTxids() (txids *[]string) {
	return
}
