package bluetooth

import (
	"time"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

// Maybe wifi cred go here

func RunBLEServer() error {
	// Enable the BLE interface
	must("enable BLE adapter", adapter.Enable())

	// Define peripheral info
	adv := adapter.DefaultAdvertisement()
	must("config adv", adv.Configure(bluetooth.AdvertisementOptions{
		LocalName: "CCCS BLE",
	}))

	must("start adv", adv.Start())
	println("Advertising CCCS BLE...")

	for {
		time.Sleep(time.Hour)
	}
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
