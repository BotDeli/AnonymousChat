package main

import (
	"AnonimousChat/pkg/config"
	"log"
)

func main() {
	pathConfig := "config.env"

	err := config.LoadEnvArgsFromFile(pathConfig)
	if err != nil {
		log.Fatal(err)
	}
}
