package wallet

import (
	"github.com/fanguanghui/coinapi/config"
)

type Factory struct {
	Environment string
}

func (this Factory) Get(coinID int) Wallet {

	config.Setup(this.Environment)

	if coinID == config.Coinid_USDT {
		return NewTether(&config.Usdtcfg, this.Environment)

	} else if coinID == config.Coinid_BTC {
		return NewBitcoin(&config.Btccfg)

	} else if coinID == config.Coinid_ETH {
		return NewEthereum(&config.Ethcfg)

	}
	return nil
}
