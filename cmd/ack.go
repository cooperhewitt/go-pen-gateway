package main

import (
	"github.com/tarm/serial"
	"github.com/the-pen/go-pen-gateway"
	"log"
	"os"
)

func main() {

	port := os.Args[1]

	c := &serial.Config{Name: port, Baud: 9600}
	s, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ACK Interrupt: %x\n", []byte(gateway.ACK_INTERRUPT))

	n, err := s.Write([]byte(gateway.ACK_INTERRUPT))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ACK Frame: %x\n", []byte(gateway.ACK_FRAME))

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ACK Response: %x\n", buf[:n])

}
