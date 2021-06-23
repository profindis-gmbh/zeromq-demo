package main

import (
	"log"

	"github.com/zeromq/goczmq"
)

const endpoint = "tcp://127.0.0.1:5555"

func main() {
	dealer, err := goczmq.NewDealer(endpoint)
	if err != nil {
		log.Fatalf("fail to create a dealer with error: %v", err)
	}
	defer dealer.Destroy()

	log.Println("dealer created and bound")

	// need to be our custom message type
	msg := "hello"
	err = dealer.SendFrame([]byte(msg), goczmq.FlagNone)
	if err != nil {
		log.Fatalf("failed to send a message with error: %v", err)
	}

	log.Printf("dealer sent message: %v", msg)

	// Receive the response
	resp, err := dealer.RecvMessage()
	if err != nil {
		log.Fatalf("failed to receive the response with error: %v", err)
	}

	log.Printf("received a response (%v)", string(resp[0]))
}
