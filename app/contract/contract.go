package contract

import (
	"encoding/json"
	"time"
)

// Coupon represents data about coupons
type Coupon struct {
	Name        string    `json:"name"`
	Brand       string    `json:"brand"`
	Value       int       `json:"value"`
	TimeCreated time.Time `json:"createdAt"`
	TimeExpiry  time.Time `json:"expiry"`
}

//TimeLayout of the contract's dates
var TimeLayout = "2006-01-02 15:04:05"

// UnmarshalJSON overrides json unmarshal logic
func (t *Coupon) UnmarshalJSON(data []byte) error {
	type Alias Coupon
	aux := &struct {
		*Alias
		TimeCreated string `json:"createdAt"`
		TimeExpiry  string `json:"expiry"`
	}{
		Alias: (*Alias)(t),
	}
	var err error
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if t.TimeCreated, err = time.Parse(TimeLayout, aux.TimeCreated); err != nil {
		return err
	}

	if t.TimeExpiry, err = time.Parse(TimeLayout, aux.TimeExpiry); err != nil {
		return err
	}
	return nil
}

// MarshalJSON overrides json marshal logic
func (t *Coupon) MarshalJSON() ([]byte, error) {
	type Alias Coupon
	return json.Marshal(&struct {
		*Alias
		TimeCreated string `json:"createdAt"`
		TimeExpiry  string `json:"expiry"`
	}{
		Alias:       (*Alias)(t),
		TimeCreated: t.TimeCreated.Format(TimeLayout),
		TimeExpiry:  t.TimeExpiry.Format(TimeLayout),
	})
}

func formatTime(time time.Time) string {
	return time.Format(TimeLayout)
}
