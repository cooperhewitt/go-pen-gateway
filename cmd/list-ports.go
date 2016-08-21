package main

import (
	"github.com/mikepb/go-serial"
	"log"
)

func main() {

	ports, err := serial.ListPorts()

	if err != nil {
		log.Panic(err)
	}

	log.Printf("\tFound %d ports. Here are the Smithsonian Gateways\n", len(ports))

	for _, info := range ports {
		if info.Description() == "Smithsonian Gateway" {
			log.Println("\tName:", info.Name())
			log.Println("\tDescription:", info.Description())
			log.Println("\tTransport:", info.Transport())

			if bus, addr, err := info.USBBusAddress(); err != nil {
				log.Println("\tbus:", bus, "\taddr:", addr)
			} else {
				log.Println(err)
			}

			if vid, pid, err := info.USBVIDPID(); err != nil {
				log.Println("\tvid:", vid, "\tpid:", pid)
			} else {
				log.Println(err)
			}

			log.Println("\tUSB Manufacturer:", info.USBManufacturer())
			log.Println("\tUSB Product:", info.USBProduct())
			log.Println("\tUSB Serial Number:", info.USBSerialNumber())
			log.Println("\tBluetooth Address:", info.BluetoothAddress())

		}
	}
}
