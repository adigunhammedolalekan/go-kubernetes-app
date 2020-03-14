package main

import (
	"github.com/adigunhammedolalekan/go-app-kubernetes"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
