package main

import (
	"log"

	"github.com/brettlazarine/cccs/server"
)

func main() {
	if err := server.RunServer(); err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
