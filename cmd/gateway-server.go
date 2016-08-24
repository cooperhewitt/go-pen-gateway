package main

import (
	"encoding/json"
	"fmt"
	"github.com/the-pen/go-pen-gateway"
	//"html"
	"log"
	"net/http"
	"os"
	"time"
)

type GatewayBoard struct {
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
}

type GatewayBoards []GatewayBoard

func main() {

	// pick a serial port
	port := os.Args[1]

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// generate the frame to request the gateway id
		frame := gateway.GenerateFrame(gateway.HOST_TO_GATEWAY, gateway.REQUEST_GATEWAY_ID)
		log.Printf("Sending REQUEST_GATEWAY_ID - %x\n", frame)

		// send the frame to the board
		rsp := gateway.SendFrame(frame, port)
		log.Printf("Received - %x\n", rsp)

		// decode the response
		id := gateway.DecodeD5Response(rsp)
		log.Printf("Gateway ID: %s\n", id)

		msg := fmt.Sprintf("Gateway ID: %s", id)

		data := GatewayBoards{
			GatewayBoard{Name: msg, Timestamp: time.Now()},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)

	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}

}
