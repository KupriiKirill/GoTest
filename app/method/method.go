package method

import (
	"fmt"
	"net/http"
)

// GetAllCoupons returns list of all coupons
func GetAllCoupons(writer http.ResponseWriter, requestPtr *http.Request) {
	fmt.Fprintln(writer, "hello")
}

// GetFileteredCoupons aplies a filter to coupons
func GetFileteredCoupons(writer http.ResponseWriter, requestPtr *http.Request) {
	fmt.Fprintln(writer, "filtered")
}

// DeleteCoupons deletes selected coupon(s)
func DeleteCoupons(writer http.ResponseWriter, requestPtr *http.Request) {
	fmt.Fprint(writer, "delete")
}

// PutCoupons adds new coupon(s) or modifies existing
func PutCoupons(writer http.ResponseWriter, requestPtr *http.Request) {
	fmt.Fprint(writer, "POST")
}
