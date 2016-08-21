package main

import (
	"github.com/tarm/serial"
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

	n, err := s.Write([]byte("P2P_INTERRUPT_ACK"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", buf[:n])

}
