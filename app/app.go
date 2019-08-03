package app

import (
	"log"
	"net/http"

	"github.com/kirillkuprii/gotest/app/handler"
	"github.com/kirillkuprii/gotest/app/method"
	"github.com/kirillkuprii/gotest/app/storage"
)

//Application hadles service flow
type Application struct {
}

//Run takes hoest address to serve on
func (t *Application) Run(hostaddr string) {

	router := new(handler.Router)
	_, err := storage.OpenStorage("coupons")
	if err != nil {
		return
	}
	router.HandleFunc("/coupons/", method.GetAllCoupons, handler.GET)
	router.HandleFunc("/coupons/filtered?.*", method.GetFileteredCoupons, handler.GET)
	router.HandleFunc("/coupons/delete", method.DeleteCoupons, handler.DELETE)
	router.HandleFunc("/coupons/put", method.PutCoupons, handler.PUT)

	log.Fatal(http.ListenAndServe(hostaddr, router))
}
