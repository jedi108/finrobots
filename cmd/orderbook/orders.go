package main

import (
	"finrgo/exhanges"
	"finrgo/exhanges/polo"
	"fmt"
	"finrgo/exhanges/poloniex"
	"finrgo/exhanges/bittrex"
	"github.com/davecgh/go-spew/spew"
)

const (
	BittrexExchange  string = "bittrex"
	PoloniexExchange string = "poloniex"
)

func main() {
	exchanges := exhanges.InitExhanges()
	bittrex0(exchanges)
}

func bittrex0(exchanges *exhanges.Exchanges)  {
	exchanges.AddExhange(BittrexExchange, &bittrex.Bittrex{})
	bittrexVal := exchanges.GetInstanceExchange(BittrexExchange)
	orders, err := exchanges.GetOrders(bittrexVal, "BTC-LTC", 1)
	if err != nil {
		fmt.Println(err)
	}
	spew.Dump(orders)
}

func polo0(exchanges *exhanges.Exchanges)  {
	exchanges.AddExhange(PoloniexExchange, &poloniex.Poloniex{})
	polo := exchanges.GetInstanceExchange(PoloniexExchange)
	orders, err := exchanges.GetOrders(polo, "BTC_LTC", 1)
	if err != nil {
		fmt.Println(err)
	}
	spew.Dump(orders)
}

func polo1(exchanges *exhanges.Exchanges)  {
	exchanges.AddExhange(PoloniexExchange, &polo.Poloniex{})
	pollo := exchanges.GetInstanceExchange(PoloniexExchange)
	orders, err := exchanges.GetOrders(pollo, "BTC_LTC", 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(orders)
}