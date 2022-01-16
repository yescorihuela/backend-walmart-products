package main

import (
	"log"

	"github.com/yescorihuela/walmart-products/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
