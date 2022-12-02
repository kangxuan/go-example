package main

import "fmt"

type valuable interface {
	getValue() float32
}

type stockPosition struct {
	ticker     string  "股票代码"
	sharePrice float32 "股票单价"
	count      float32 "股票份数"
}

func (s *stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c *car) getValue() float32 {
	return c.price
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

func main() {
	var o valuable = &stockPosition{"GOOG", 5878.22, 5}
	showValue(o)

	o = &car{"BMW", "M3", 654433}
	showValue(o)
}

//Value of the asset is 29391.101562
//Value of the asset is 654433.000000
