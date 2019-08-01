package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kirillkuprii/gotest/app/handler"
)

type Application struct {
}

func (t *Application) Test() {
	fmt.Println("test")
}

func (t *Application) Run(hostaddr string) {

	router := new(handler.Router)

	router.HandleFunc("/getall/[0-9]", handler.GetAllCoupons)

	log.Fatal(http.ListenAndServe(hostaddr, router))
}
