package main

import (
	"github.com/the-pen/go-pen-gateway"
	"log"
	"os"
)

func main() {

	// pick a serial port
	port := os.Args[1]

	// generate the frame to request the gateway id
	frame := gateway.GenerateFrame(gateway.HOST_TO_GATEWAY, gateway.REQUEST_GATEWAY_ID)
	log.Printf("Sending REQUEST_GATEWAY_ID - %x\n", frame)

	// send the frame to the board
	rsp := gateway.SendFrame(frame, port)
	log.Printf("Received - %x\n", rsp)

	// decode the response
	id := gateway.DecodeD5Response(rsp)
	log.Printf("Gateway ID: %s\n", id)

}
