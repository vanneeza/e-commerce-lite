package main

import (
	"github.com/vanneeza/e-commerce-lite/server"
)

func main() {
	if err := server.Run(); err != nil {
		panic(err)
	}
}
