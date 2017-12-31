package main

import (
	poloniex2 "finrgo/exhanges/poloniex"
	"fmt"
)

var (
	input string
	polo  = poloniex2.InitPoloniex()
)

func PrintBal() {
	polo.GetBalance()
	polo.PrintNotNull()
	polo.PrintInBTCValues()
	fmt.Println("All In USD", polo.BalanceInValues.BTCValues * polo.Pair.Close)
}

func main() {
	polo.GetChartCurrent("USDT_BTC")
	PrintBal()
	//fmt.Scan(&input)
}
