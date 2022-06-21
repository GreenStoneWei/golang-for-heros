package main

import "fmt"

type couponType int

const (
	MINUS_50 couponType = iota
	MINUS_100
	PERCENT_OFF_90
	PERCENT_OFF_80
	REBATE_150_FOR_1000
)

func calculate(price int, coupon couponType) int {
	if coupon == MINUS_50 {
		price -= 50
	} else if coupon == MINUS_100 {
		price -= 100
	} else if coupon == PERCENT_OFF_90 {
		price = int(float64(price) * 0.9)
	} else if coupon == PERCENT_OFF_80 {
		price = int(float64(price) * 0.8)
	} else if coupon == REBATE_150_FOR_1000 {
		if price > 1000 {
			price -= 150
		}
	}
	if price < 0 {
		return 0
	} else {
		return int(price)
	}

}

type coupon interface {
	calculate(price int) int
}

type cashier struct {
	coupon coupon
}

func NewCashier() *cashier {
	return &cashier{}
}
func (c *cashier) setCoupon(coupon coupon) {
	c.coupon = coupon
}
func (c *cashier) calculate(price int) int {
	return c.coupon.calculate(price)
}

type minus_50_coupon struct{}

func (mc minus_50_coupon) calculate(price int) int {
	return price - 50
}

type percent_off_90 struct{}

func (mc percent_off_90) calculate(price int) int {
	return int(float64(price) * 0.9)
}

func main() {

	// if else pattern
	fmt.Println(calculate(1150, MINUS_50))

	// strategy pattern
	cashier := NewCashier()
	cashier.setCoupon(minus_50_coupon{})
	fmt.Println(cashier.calculate(1000))

	cashier.setCoupon(percent_off_90{})
	fmt.Println(cashier.calculate(1000))
}
