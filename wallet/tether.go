package wallet

import (
	"github.com/fanguanghui/coinapi/config"
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

	if environment == config.Env_AMAZON_AWS {
		Tether.propertyid = config.Propertyid_PROD
		Tether.ecosystem = config.Ecosystem_PROD
	} else {
		Tether.propertyid = config.Propertyid_TEST
		Tether.ecosystem = config.Ecosystem_TEST
	}
	return Tether
}

// 账户相关接口
func (this Tether) NewAddress(account string) (address string) {
	address = this.omni.GetNewAddress(account)
	return
}

func (this Tether) GetBalance(address string) (balance string) {
	balance, _ = this.omni.GetBalance(address, this.propertyid)
	return
}

// 交易相关接口
func (this Tether) SendTransaction(fromaddress, toaddress string, amount, feeaddress string) (txid string) {
	return ""
}

func (this Tether) CollectTransaction() (txid string) {
	return ""
}

func (this Tether) GetTransaction(txid string) (tx config.Tx) {
	row := this.omni.GetTransaction(txid)
	if row.Txid != "" && row.Propertyid == this.propertyid {
		sends := [][2]string{
			{row.Referenceaddress, row.Amount},
		}
		tx.Sends = &sends
		tx.Txid = row.Txid
		tx.From = row.Sendingaddress
		tx.Cfms = row.Confirmations
		tx.Valid = row.Valid
		tx.Time = row.Blocktime
	}
	return
}

func (this Tether) GetBlockCount() (count uint64) {
	count = this.omni.GetBlockCount()
	return
}

func (this Tether) GetBlockTxids(index uint64) (txids []string) {
	txids = this.omni.GetBlockTransactions(index)
	return
}

func (this Tether) GetPendingTxs(address string) (txs []config.Tx) {
	list := this.omni.GetPendingTransactions(address)
	for _, row := range list {
		if row.Txid != "" && row.Propertyid == this.propertyid {
			tx := config.Tx{}
			sends := [][2]string{
				{row.Referenceaddress, row.Amount},
			}
			tx.Sends = &sends
			tx.Txid = row.Txid
			tx.From = row.Sendingaddress
			tx.Cfms = row.Confirmations
			tx.Valid = row.Valid
			tx.Time = row.Blocktime
			txs = append(txs, tx)
		}
	}
	return
}
