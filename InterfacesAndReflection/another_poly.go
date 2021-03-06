package main

import f "fmt"

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

/* method to determine the value of a stock position */
func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

/* method to determine the value of a car */
func (c car) getValue() float32 {
	return c.price
}

/* contract that defines different things that have value */
type valuable interface {
	getValue() float32
}

/* anything that satisfies the "valuable" interface is accepted */
func showValue(asset valuable) {
	f.Printf("Value of the asset is %f\n", asset.getValue())
}

func main() {
	//
	var o valuable = stockPosition{"GOOG", 577.20, 4}
	showValue(o)

	// Generalizing and reusing code using interface.
	o = car{"BMW", "M3", 66500}
	showValue(o)
}
