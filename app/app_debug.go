package app

import (
	"time"

	"github.com/kirillkuprii/gotest/app/contract"
)

func (t *Application) addTestData() {
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
			Brand:       "brand1",
			Value:       20,
			TimeCreated: time.Now(),
			TimeExpiry:  time.Now(),
		},
		&contract.Coupon{
			Name:        "coupon4",
			Brand:       "brand1",
			Value:       50,
			TimeCreated: time.Now(),
			TimeExpiry:  time.Now(),
		},
	}
	t.db.AddItems(coupons)
}
