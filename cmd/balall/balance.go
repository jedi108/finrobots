package main

import (
	"finrgo/exhanges"
	"finrgo/exhanges/poloniex"
	"finrgo/exhanges/bittrex"

)

const (
	BittrexExchange  string = "bittrex"
	PoloniexExchange string = "poloniex"
)

func main() {
	exchanges := exhanges.InitExhanges()
	exchanges.AddExhange(BittrexExchange, &bittrex.Bittrex{})
	exchanges.AddExhange(PoloniexExchange, &poloniex.Poloniex{})

	bittrexVal := exchanges.GetInstanceExchange(BittrexExchange)
	bittrexVal.GetBalance()

	poloniexVal := exchanges.GetInstanceExchange(PoloniexExchange)
	poloniexVal.GetBalance()
}
