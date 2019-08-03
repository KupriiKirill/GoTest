package contract

import "time"

type Coupon struct {
	Name        string
	Brand       string
	Value       int
	TimeCreated time.Time
	TimeExpiry  time.Time
}
