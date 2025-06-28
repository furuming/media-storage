package main

import (
	"storage/cmd/config"
	"storage/cmd/router"
)

func main() {

	env := config.New()

	router := router.New(env)

	router.Run(":8000")
}
