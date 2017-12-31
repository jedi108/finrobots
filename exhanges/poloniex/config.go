package poloniex

import (
	"github.com/pharrisee/poloniex-api"
	"sync"
	"github.com/k0kubun/pp"
)

type BalanceInValues struct {
	BTCValues  float64
	USDTValues float64
	Balance    poloniex.Balances
}

type ConfigPoloniex struct {
	BalanceInValues
	Pair
	Poloniex *poloniex.Poloniex
}

type Pair struct {
	sync.RWMutex
	Open  float64
	Close float64
	High  float64
	Low   float64
}

func InitPoloniex() *ConfigPoloniex {
	init := poloniex.NewWithCredentials("TF1Y8W86-BNOJN7GI-014QHRVF-B67G9C4T", "8ebee83bf20bfe63cf384bc8c9978859f3ca364bdf825e32011c9d1bd7482a49eaaf19feb1dc0203662ce452f3eb0a7eb5639c7ed24c9120fb4f9e29b704538f")
	init.Subscribe("ticker")
	init.Subscribe("USDT_BTC")
	init.Subscribe("BTC_NXT")

	init.On("ticker", func(m poloniex.WSTicker) {
		pp.Println(m)
	}).On("USDT_BTC-trade", func(m poloniex.WSOrderbook) {
		pp.Println(m)
	})
	return &ConfigPoloniex{
		Poloniex: init}
}
