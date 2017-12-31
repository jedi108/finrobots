package main

import (
	"finrgo/strategy"
	"finrgo/strategy/signals"
)

func main() {
	Trade()
}

func Trade()  {
	arbStrategy := strategy.NewStrategy()
	arbStrategy.Attach(&signals.SignalArb{})
	arbStrategy.Start()
}