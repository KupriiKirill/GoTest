package app

import (
	"time"

	"github.com/kirillkuprii/gotest/app/contract"
)

func (t *Application) addTestData() {
	time1, _ := time.Parse(contract.TimeLayout, "2012-12-21 10:03:23")
	time2, _ := time.Parse(contract.TimeLayout, "2012-11-21 10:03:23")
	coupons := []*contract.Coupon{
		&contract.Coupon{
			Name:        "coupon1",
			Brand:       "brand1",
			Value:       30,
			TimeCreated: time.Now(),
			TimeExpiry:  time.Now(),
		},
		&contract.Coupon{
			Name:        "coupon2",
			Brand:       "brand1",
			Value:       10,
			TimeCreated: time.Now(),
			TimeExpiry:  time.Now(),
		},
		&contract.Coupon{
			Name:        "coupon3",
			Brand:       "brand3",
			Value:       20,
			TimeCreated: time1,
			TimeExpiry:  time.Now(),
		},
		&contract.Coupon{
			Name:        "coupon4",
			Brand:       "brand4",
			Value:       50,
			TimeCreated: time2,
			TimeExpiry:  time.Now(),
		},
	}
	t.db.AddItems(coupons)
}
