package wallet

import "github.com/fanguanghui/coinapi/setting"

type Wallet interface {

	// 创建地址
	NewAddress(account string) (address string)

	// 发起提币
	SendTransaction(fromaddress, toaddress string, amount, feeaddress string) (txid string)

	CollectTransaction() (txid string)

	GetBalance(address string) (balance string)

	GetTransaction() (tx *setting.Tx)

	GetBlockCount() (count uint32)

	GetBlockTxids() (txids *[]string)
}
