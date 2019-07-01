package wallet

import (
	"github.com/go-develop/coinapi/config"
)

type Wallet interface {

	/*

	 */

	// 创建地址
	NewAddress(account string) (address string)

	GetBalance(address string) (balance string)

	/*

	 */

	// 发起提币
	SendTransaction(fromaddress, toaddress string, amount, feeaddress string) (txid string)

	CollectTransaction() (txid string)

	GetBlockCount() (index uint64)

	GetTxById(txid string) (tx config.Tx)

	GetBlockTxs(index uint64) (txids []config.Tx)

	GetPendingTxs(address string) (txs []config.Tx)
}
