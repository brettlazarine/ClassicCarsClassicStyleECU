package main

import (
	"log"

	"github.com/brettlazarine/cccs/bluetooth"
	"github.com/brettlazarine/cccs/server"
)

func main() {
	done := make(chan bool)

	go func() {
		if err := server.RunServer(); err != nil {
			log.Fatalf("Error running server: %v", err)
		}
		done <- true
	}()

	go func() {
		if err := bluetooth.RunBLEServer(); err != nil {
			log.Fatalf("Error running BLE server: %v", err)
		}
		done <- true
	}()

	<-done

	// if err := server.RunServer(); err != nil {
	// 	log.Fatalf("Error running server: %v", err)
	// }
	// if err := bluetooth.RunBLEServer(); err != nil {
	// 	log.Fatalf("Error running BLE server: %v", err)
	// }
}
