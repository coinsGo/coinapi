package wallet

import (
	"github.com/fanguanghui/coinapi/config"
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

	GetTransaction(txid string) (tx config.Tx)

	GetBlockCount() (index uint64)

	GetBlockTxids(index uint64) (txids []string)

	GetPendingTxs(address string) (txs []config.Tx)
}
