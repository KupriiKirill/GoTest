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
	db storage.Storage
}

//Run takes host address to serve on
func (t *Application) Run(hostaddr string) {

	router := new(handler.Router)
	var err error
	t.db, err = storage.OpenStorage("coupons")
	if err != nil {
		log.Fatal("Cannot open database")
		return
	}

	t.addTestData()

	router.HandleFunc("/coupons/$", t.handleDBRequest(method.GetAllCoupons), handler.GET)
	router.HandleFunc("/coupons/$", t.handleDBRequest(method.GetCouponsCount), handler.HEAD)
	router.HandleFunc("/coupons/filtered?.*", t.handleDBRequest(method.GetFileteredCoupons), handler.GET)
	router.HandleFunc("/coupons/delete/$", t.handleDBRequest(method.DeleteCoupons), handler.DELETE)
	router.HandleFunc("/coupons/put/$", t.handleDBRequest(method.PutCoupons), handler.PUT)

	log.Fatal(http.ListenAndServe(hostaddr, router))
}

type dbRequestHandlerFunc func(storage storage.Storage, writer http.ResponseWriter, requestPtr *http.Request)

func (t *Application) handleDBRequest(handler dbRequestHandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, requestPtr *http.Request) {
		handler(t.db, writer, requestPtr)
	}
}
