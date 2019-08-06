package method

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/kirillkuprii/gotest/app/contract"
	"github.com/kirillkuprii/gotest/app/filters"
	"github.com/kirillkuprii/gotest/app/storage"
)

// GetAllCoupons returns list of all coupons
func GetAllCoupons(db storage.Storage, writer http.ResponseWriter, requestPtr *http.Request) {
	data, err := json.Marshal(db.GetAllItems())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(writer, "%s", data)
}

// GetCouponsCount returns count of all coupons
func GetCouponsCount(db storage.Storage, writer http.ResponseWriter, requestPtr *http.Request) {
	number := len(db.GetAllItems())
	writer.Header().Set("Content-Length", fmt.Sprintf("%v", number))
}

// GetFilteredCoupons aplies a filter to coupons
func GetFilteredCoupons(db storage.Storage, writer http.ResponseWriter, requestPtr *http.Request) {
	filters, err := getFilters(requestPtr.URL.String())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	items := make(map[int]*contract.Coupon)
	for key, item := range db.GetAllItems() {
		bFiltered := false
		for _, filter := range filters {
			if !filter.IsPassing(item) {
				bFiltered = true
				break
			}
		}
		if !bFiltered {
			items[key] = item
		}
	}
	data, err := json.Marshal(items)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(writer, "%s", data)
}

// DeleteCoupons deletes selected coupon(s)
func DeleteCoupons(db storage.Storage, writer http.ResponseWriter, requestPtr *http.Request) {
	items := []storage.UID{}

	body, err := ioutil.ReadAll(requestPtr.Body)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, &items); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	db.DeleteItems(items)
	writer.WriteHeader(http.StatusNoContent)
}

// PutCoupons adds new coupon(s) or modifies existing
func PutCoupons(db storage.Storage, writer http.ResponseWriter, requestPtr *http.Request) {
	var PutList struct {
		Update map[storage.UID]contract.Coupon `json:"update"`
		Add    []contract.Coupon               `json:"add"`
	}

	body, err := ioutil.ReadAll(requestPtr.Body)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, &PutList); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	db.UpdateItems(&PutList.Update)

	itemsToAdd := []*contract.Coupon{}
	for _, coupon := range PutList.Add {
		itemsToAdd = append(itemsToAdd, &coupon)
	}

	db.AddItems(itemsToAdd)
}

func getFilters(url string) ([]filters.Filter, error) {
	parsed := strings.Split(strings.Replace(url, "%20", " ", -1), "?")
	params := parsed[1:]
	filtersList := []filters.Filter{}
	for _, param := range params {
		filter, err := filters.CreateFilter(param)
		if err != nil {
			return nil, err
		}
		filtersList = append(filtersList, filter)
	}
	return filtersList, nil
}
