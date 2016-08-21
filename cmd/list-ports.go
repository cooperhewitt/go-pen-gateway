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

	log.Printf("Found %d ports:\n", len(ports))

	for _, info := range ports {
		log.Println(info.Name())
		log.Println("\tName:", info.Name())
		log.Println("\tDescription:", info.Description())
		log.Println("\tTransport:", info.Transport())
	}
}
