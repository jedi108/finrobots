package poloniex

import (
	"fmt"
)

func (c *ConfigPoloniex) GetBalance() {
	bl, err := c.Poloniex.Balances()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range bl {
		if v.BTCValue != 0 {
			c.BalanceInValues.BTCValues = c.BalanceInValues.BTCValues + v.BTCValue
		}
		//fmt.Println(k,v)
	}
	c.Balance = bl
}

func (c *ConfigPoloniex) PrintInBTCValues()  {
	fmt.Println("All in BTC", c.BalanceInValues.BTCValues)
}

func (c *ConfigPoloniex) PrintNotNull() {
	for k, v := range c.Balance {
		if v.Available != 0 || v.BTCValue != 0 || v.OnOrders != 0 {
			fmt.Println(k, v)
		}
	}
}

func (c *ConfigPoloniex) GetChartCurrent(pair string) {
	dt, err := c.Poloniex.ChartDataCurrent(pair)
	if err == nil {
		c.Pair.Close = dt[0].Close
		c.Pair.Open = dt[0].Open
		c.Pair.High = dt[0].High
		c.Pair.Low = dt[0].Low
	} else {
		fmt.Println(err)
	}
}
