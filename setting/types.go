package setting

type Send struct {
	Address string
	Amount  float32
}

type Tx struct {
	Time  string
	Txid  string
	Sends []Send
	Cfms  uint32
	Valid bool
}
