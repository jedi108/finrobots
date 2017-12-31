package services

import "finrgo/exhanges"

type (
	SendOrderService struct {
		exhanges.Exchange
	}
	SendOrderData struct {
		Rate float64
		Amount float64
	}

	Result struct {
		err error
		resultString string
	}

	OrderConnect struct {
		SendOrderData
		Result
	}
)

func OrderPool(connect chan OrderConnect)  {

}