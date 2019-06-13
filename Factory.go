package coinapi

type Factory struct {
}

func (this Factory) Get(coin, host, user, pass string) Wallet {

	if coin == "USDT" {
		return NewTether(host, user, pass)
	} else if coin == "BTC" {
		return NewBitcoin(host, user, pass)
	} else if coin == "ETH" {
		return NewEthereum(host, user, pass)
	}
	return nil
}
