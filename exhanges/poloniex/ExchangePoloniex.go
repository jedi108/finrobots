package poloniex

import (
	"github.com/pharrisee/poloniex-api"
	"finrgo/exhanges"
	"github.com/shopspring/decimal"
)

type (
	Poloniex struct {
		poloApi *poloniex.Poloniex
	}
)

func (p *Poloniex) InitExchange() {
	p.poloApi = poloniex.NewWithCredentials(API_POLONIEX_KEY, API_POLONIEX_SECRET)
}

func (p *Poloniex) GetBalance() {

}

func (b *Poloniex) GetOrdersByIndex(pair string, idx int) (*exhanges.Orders, error) {
	orders, err := b.poloApi.OrderBook(pair)
	if err != nil {
		return &exhanges.Orders{}, err
	}
	return &exhanges.Orders{
		Ask: exhanges.Order{
			Amount: decimal.NewFromFloat(orders.Asks[0].Rate),
			Rate:   decimal.NewFromFloat(orders.Asks[0].Amount),
		},
		Bid: exhanges.Order{
			Amount: decimal.NewFromFloat(orders.Bids[0].Rate),
			Rate:   decimal.NewFromFloat(orders.Bids[0].Amount),
		},
	}, err
}

//func (p *Poloniex) SendingSellOrder(pair string, rate float64, amount float64) {
//	p.poloApi.Buy(pair, rate, amount)
//}
//
//func (p *Poloniex) SendingBuyOrder(pair string, rate float64, amount float64) (string, error) {
//	return p.poloApi.Sell(pair, rate, amount)
//}
