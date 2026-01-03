package main

import (
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/godbus/dbus/v5"
)

func main() {
	targetMAC := "00:1B:DC:C0:67:AF"
	conn, err := dbus.SystemBus()
	if err != nil {
		log.Fatalf("Failed to connect to system bus: %v", err)
	}
	defer conn.Close()
	obj := conn.Object("org.bluez", "/")
	var managedObjects map[dbus.ObjectPath]map[string]map[string]dbus.Variant
	err = obj.Call("org.freedesktop.DBus.ObjectManager.GetManagedObjects", 0).Store(&managedObjects)
	if err != nil {
		log.Fatalf("Failed to call GetManagedObjects: %v", err)
	}
	found := false
	for objPath, interfaces := range managedObjects {
		if adapterProps, ok := interfaces["org.bluez.Adapter1"]; ok {
			if addrVariant, ok := adapterProps["Address"]; ok {
				addr := addrVariant.Value().(string)
				if strings.EqualFold(addr, targetMAC) {
					hciName := path.Base(string(objPath))
					fmt.Printf("Found adapter: %s (at %s)\n", hciName, objPath)
					found = true
					break
				}
			}
		}
	}

	if !found {
		fmt.Println("Adapter not found for address:", targetMAC)
	}
}
