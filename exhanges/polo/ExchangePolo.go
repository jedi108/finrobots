package polo

import (
	"fmt"
	"finrgo/exhanges"
	"github.com/shopspring/decimal"
	polo "go-poloniex"
	"github.com/davecgh/go-spew/spew"
)


type (
	Poloniex struct {
		poloApi *polo.Poloniex
	}
)

func (p *Poloniex) InitExchange() {
	p.poloApi = polo.New(API_POLONIEX_KEY, API_POLONIEX_SECRET)
}

func (b *Poloniex) GetBalance() {
	balances, err := b.poloApi.GetBalances()
	fmt.Println(err, balances)
}

func (p *Poloniex) GetOrdersByIndex(pair string, idx int) (*exhanges.Orders, error) {
	result, err := p.poloApi.GetOrderBook(pair, "both", 1)
	if err != nil {
		return &exhanges.Orders{}, err
	}

	//{
		spew.Dump(result.Asks)
		fmt.Println(result.Asks)
		fmt.Println(result.Bids)
		return &exhanges.Orders{}, err
	//}

	rate, err := decimal.NewFromString(fmt.Sprintf("%s", result.Asks[idx]))
	if err != nil {
		return &exhanges.Orders{}, err
	}
	amount, err := decimal.NewFromString(fmt.Sprintf("%s", result.Asks[idx]))
	if err != nil {
		return &exhanges.Orders{}, err
	}
	return &exhanges.Orders{
		Ask: exhanges.Order{
			Rate:   rate,
			Amount: amount,
		},
		Bid: exhanges.Order{},
	}, err
}
