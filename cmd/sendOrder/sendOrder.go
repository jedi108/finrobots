package main

import (
	"finrgo/exhanges"
	"finrgo/exhanges/bittrex"
	"finrgo/exhanges/polo"
	"finrgo/exhanges/poloniex"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

const (
	BittrexExchange  string = "bittrex"
	PoloniexExchange string = "poloniex"
)

func main() {
	exchanges := exhanges.InitExhanges()
	bittrexSend(exchanges)
}

func bittrexSend(exchanges *exhanges.Exchanges)  {
	exchanges.AddExhange(BittrexExchange, &bittrex.Bittrex{})
	bittrexVal := exchanges.GetInstanceExchange(BittrexExchange)
	orders, err := exchanges.GetOrders(bittrexVal, "BTC-LTC", 1)
	if err != nil {
		fmt.Println(err)
	}
	spew.Dump(orders)
}

func poloSend(exchanges *exhanges.Exchanges)  {
	//exchanges.AddExhange(PoloniexExchange, &poloniex.Poloniex{})
	polo := exchanges.GetInstanceExchange(PoloniexExchange)
	orders, err := exchanges.GetOrders(polo, "BTC_LTC", 1)
	if err != nil {
		fmt.Println(err)
	}
	spew.Dump(orders)
}

func polo1Send(exchanges *exhanges.Exchanges)  {
	//exchanges.AddExhange(PoloniexExchange, &polo.Poloniex{})
	pollo := exchanges.GetInstanceExchange(PoloniexExchange)
	orders, err := exchanges.GetOrders(pollo, "BTC_LTC", 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(orders)
}