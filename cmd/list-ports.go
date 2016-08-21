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
		}
	}
}
