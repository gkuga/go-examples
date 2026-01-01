package main

import (
	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func main() {
	// Enable BLE interface.
	must("enable BLE stack", adapter.Enable())

	// Start scanning for devices.
	println("scanning...")
	err := adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		println("found device:", device.Address.String(), device.RSSI, device.LocalName())
		if device.LocalName() != "XXXX" {
			println("not the target device, continuing scan...")
			return
		}
		adapter.StopScan()
		println("connecting to device:", device.Address.String())
		peer, err := adapter.Connect(device.Address, bluetooth.ConnectionParams{})
		if err != nil {
			println("failed to connect to device:", err.Error())
			return
		}
		println("connected to device:", device.Address.String())

		// Perform operations with the connected peer here.
		// For example, you can discover services, read/write characteristics, etc.
		// Discover services and characteristics
		services, err := peer.DiscoverServices(nil)
		if err != nil {
			println("failed to discover services:", err.Error())
			return
		}
		for _, service := range services {
			println("discovered service:", service.UUID().String())
			characteristics, err := service.DiscoverCharacteristics(nil)
			if err != nil {
				println("failed to discover characteristics:", err.Error())
				continue
			}
			for _, char := range characteristics {
				println("discovered characteristic:", char.UUID().String())

				// Display properties of the characteristic
				props := char.Flags()
				println("  Readable:", props.Read)
				println("  Writable:", props.Write)
				println("  Notifiable:", props.Notify)
				println("  Indicatable:", props.Indicate)
			}
		}
	})
	must("start scan", err)
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
