package exhanges

import (
	"github.com/shopspring/decimal"
)

type (
	Exchanges struct {
		Exchanges map[string]IExhanged
	}
	Exchange struct {
		IExhanged
	}
	Order struct {
		Rate   decimal.Decimal
		Amount decimal.Decimal
	}
	Orders struct {
		Ask Order
		Bid Order
	}
	OrderDataSend struct {
		Pair string
		Order
	}
	OrderDataReceive struct {
		UID string
		Err error
	}
)

type (
	IExhanged interface {
		InitExchange()
		IBalance
		IOrders
	}
	IBalance interface {
		GetBalance()
	}
	IOrders interface {
		GetOrdersByIndex(pair string, idx int) (*Orders, error)
		//SendingSellOrder(pair string, rate float64, amount float64) (string, error)
		//SendingBuyOrder(pair string, rate float64, amount float64) (string, error)
	}
)

func InitExhanges() *Exchanges {
	return &Exchanges{
		Exchanges: make(map[string]IExhanged),
	}
}

func (exs *Exchanges) AddExhange(nameOfExchange string, exchange IExhanged) {
	exs.Exchanges[nameOfExchange] = exchange
	exchange.InitExchange()
}

func (exs *Exchanges) GetInstanceExchange(exchangeName string) IExhanged {
	return exs.Exchanges[exchangeName]
}

func (exs *Exchanges) GetOrders(ex IOrders, pair string, idx int) (*Orders, error) {
	return ex.GetOrdersByIndex(pair, idx)
}

//func (exs *Exchanges) SendSellOrder(ex IOrders, pair string, rate float64, amount float64) (string, error) {
//	return ex.SendingSellOrder(pair, rate, amount)
//}
//
//func (exs *Exchanges) SendBuyOrder(ex IOrders, pair string, rate float64, amount float64) (string, error) {
//	return ex.SendingBuyOrder(pair, rate, amount)
//}
