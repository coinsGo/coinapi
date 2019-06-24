package config

type Vout struct {
	Address string
	Amount  string
}

type Tx struct {
	Txid  string
	From  string
	Cfms  uint64
	Valid bool
	Time  uint32
	Vouts map[int]Vout
}
