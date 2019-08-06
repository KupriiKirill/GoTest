package filters

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kirillkuprii/gotest/app/contract"
)

//Filter is a basic interface for filters
type Filter interface {
	IsPassing(coupon *contract.Coupon) bool
}

//CreateFilter creates a filter from description
func CreateFilter(description string) (Filter, error) {
	parsed := strings.Split(description, "+")
	if len(parsed) < 3 {
		return nil, fmt.Errorf("cannot create filter, description is incomplete: %s", description)
	}
	switch parsed[0] {
	case "name":
		if parsed[1] != equal {
			return nil, fmt.Errorf("cannot create filter, description is incomplete: %s", description)
		}
		return &nameFilter{parsed[2]}, nil
	case "brand":
		if parsed[1] != equal {
			return nil, fmt.Errorf("cannot create filter, description is incomplete: %s", description)
		}
		return &brandFilter{parsed[2]}, nil
	case "value":
		value, err := strconv.Atoi(parsed[2])
		if err != nil {
			return nil, err
		}
		return &valueFilter{parsed[1], value}, nil
	case "createdAt":
		timeStamp, err := time.Parse(contract.TimeLayout, parsed[2])
		if err != nil {
			return nil, err
		}
		return &createdTimeFilter{parsed[1], timeStamp}, nil
	case "expiry":
		timeStamp, err := time.Parse(contract.TimeLayout, parsed[2])
		if err != nil {
			return nil, err
		}
		return &expliryTimeFilter{parsed[1], timeStamp}, nil
	default:
	}
	return nil, nil
}

const (
	more  = "mt"
	less  = "lt"
	equal = "eq"
)

type nameFilter struct {
	name string
}

type brandFilter struct {
	brand string
}

type valueFilter struct {
	more  string
	value int
}

type createdTimeFilter struct {
	more  string
	value time.Time
}

type expliryTimeFilter struct {
	more  string
	value time.Time
}

func (t *nameFilter) IsPassing(coupon *contract.Coupon) bool {
	return strings.Contains(coupon.Name, t.name)
}

func (t *brandFilter) IsPassing(coupon *contract.Coupon) bool {
	return strings.Contains(coupon.Brand, t.brand)
}

func (t *valueFilter) IsPassing(coupon *contract.Coupon) bool {
	switch t.more {
	case more:
		return coupon.Value > t.value
	case less:
		return coupon.Value < t.value
	case equal:
		return coupon.Value == t.value
	}
	return false
}

func (t *createdTimeFilter) IsPassing(coupon *contract.Coupon) bool {
	switch t.more {
	case more:
		return coupon.TimeCreated.Unix() > t.value.Unix()
	case less:
		return coupon.TimeCreated.Unix() < t.value.Unix()
	case equal:
		return coupon.TimeCreated.Unix() == t.value.Unix()
	}
	return false
}

func (t *expliryTimeFilter) IsPassing(coupon *contract.Coupon) bool {
	switch t.more {
	case more:
		return coupon.TimeExpiry.Unix() > t.value.Unix()
	case less:
		return coupon.TimeExpiry.Unix() < t.value.Unix()
	case equal:
		return coupon.TimeExpiry.Unix() == t.value.Unix()
	}
	return false
}
