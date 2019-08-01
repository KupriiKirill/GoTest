package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kirillkuprii/gotest/app/handler"
)

type Application struct {
	RequestHandler handler.HttpHandler
}

func (t *Application) Test() {
	fmt.Println("test")
}

func (t *Application) Run(hostaddr string) {
	log.Fatal(http.ListenAndServe(hostaddr, t.RequestHandler))
}
