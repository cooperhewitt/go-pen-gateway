package main

import (
	"fmt"
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

	const ack_interrupt = "\x55\x55\x00\x00\x00\x00\x00\x00\x00\x00\xff\x00\xff\x00\x00\x00\xff\x00\xff\x00"
	fmt.Printf("ACK Interrupt: %x\n", []byte(ack_interrupt))

	n, err := s.Write([]byte(ack_interrupt))
	if err != nil {
		log.Fatal(err)
	}

	const ack_frame = "\x00\xFF\x00\xFF"
	fmt.Printf("ACK Frame: %x\n", []byte(ack_frame))

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ACK Response: %x\n", buf[:n])

}
