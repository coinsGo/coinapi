package config

type Tx struct {
	Txid  string
	From  string
	Cfms  uint64
	Valid bool
	Time  uint32
	Sends *[][2]string
}
