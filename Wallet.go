package coinapi

type Wallet interface {
	NewAddress(account string) (address string)

	GetBalance(address string, propertyid uint32) (balance string)

	SendTransaction() (hash string)

	GetTransaction() (tx *Tx)

	GetBlockCount() (count uint32)

	GetBlockTransactions() (txids *[]string)
}
