package wallet

import (
	"github.com/fanguanghui/coinapi/setting"
)

type Factory struct {
	Environment string
}

func (this Factory) Get(coinName string) Wallet {

	setting.Setup(this.Environment)

	if coinName == setting.CoinName_USDT {
		return NewTether(setting.UsdtCfg, this.Environment)

	} else if coinName == setting.CoinName_BTC {
		return NewBitcoin(setting.BtcCfg)

	} else if coinName == setting.CoinName_ETH {
		return NewEthereum(setting.EthCfg)

	}
	return nil
}
