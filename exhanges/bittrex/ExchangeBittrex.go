package bittrex

import (
	"fmt"
	"finrgo/exhanges"
	"github.com/toorop/go-bittrex"
)

type (
	Bittrex struct {
		bittrexApi *bittrex.Bittrex
	}
)

func (b *Bittrex) InitExchange() {
	b.bittrexApi = bittrex.New(API_BITTREX_KEY, API_BITTREX_SECRET)
}

func (b *Bittrex) GetBalance() {
	balances, err := b.bittrexApi.GetBalances()
	fmt.Println(err, balances)
	if err != nil {

	}
}

func (b *Bittrex) GetOrdersByIndex(pair string, idx int) (*exhanges.Orders, error) {
	resutl, err := b.bittrexApi.GetOrderBook(pair, "both")
	if err != nil {
		return &exhanges.Orders{}, err
	}
	return &exhanges.Orders{
		Ask: exhanges.Order{
			Rate:   resutl.Buy[idx].Rate,
			Amount: resutl.Buy[idx].Quantity,
		},
		Bid: exhanges.Order{
			Rate:   resutl.Sell[idx].Rate,
			Amount: resutl.Sell[idx].Quantity,
		},
	}, err
}

func (b *Bittrex) SendingSellOrder(pair string, rate float64, amount float64) (string, error) {
	return b.bittrexApi.SellLimit(pair, rate, amount)
}

func (b *Bittrex) SendingBuyOrder(pair string, rate float64, amount float64) (string, error) {
	return b.bittrexApi.BuyLimit(pair, rate, amount)
}