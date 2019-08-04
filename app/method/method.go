package method

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kirillkuprii/gotest/app/contract"
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

// GetFileteredCoupons aplies a filter to coupons
func GetFileteredCoupons(db storage.Storage, writer http.ResponseWriter, requestPtr *http.Request) {
	fmt.Fprintln(writer, "filtered")
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
		Add    []contract.Coupon `json:"add"`
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
