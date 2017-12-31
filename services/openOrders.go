package services

import "finrgo/exhanges"

type (
	Orders []Order
	Order struct {
		exhanges.Exchange
	}
)

