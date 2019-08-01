package main

import "github.com/kirillkuprii/gotest/app"

func main() {
	application := app.Application{}
	application.Run("localhost:500")
}
