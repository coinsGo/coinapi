package config

type Vout struct {
	ToAddress string
	TxAmount  string
}

type Tx struct {
	Txid        string
	FromAddress string
	Confirms    uint64
	Valid       bool
	Time        uint32
	Vouts       map[int]Vout
}
