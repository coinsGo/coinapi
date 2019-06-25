package wallet

import (
	"github.com/fanguanghui/coinapi/config"
)

type Factory struct {
	Environment string
}

func (this Factory) Get(coinId uint) Wallet {

	config.Setup(this.Environment)

	if coinId == config.CoinId_USDT {
		return NewTether(&config.Usdtcfg, this.Environment)

	} else if coinId == config.CoinId_BTC {
		return NewBitcoin(&config.Btccfg)

	} else if coinId == config.CoinId_ETH {
		return NewEthereum(&config.Ethcfg)

	}
	return nil
}
