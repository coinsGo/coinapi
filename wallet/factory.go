package wallet

import (
	"github.com/fanguanghui/coinapi/config"
)

type Factory struct {
	Environment string
}

func (this Factory) Get(coinName string) Wallet {

	config.Setup(this.Environment)

	if coinName == config.CoinName_USDT {
		return NewTether(config.UsdtCfg, this.Environment)

	} else if coinName == config.CoinName_BTC {
		return NewBitcoin(config.BtcCfg)

	} else if coinName == config.CoinName_ETH {
		return NewEthereum(config.EthCfg)

	}
	return nil
}
