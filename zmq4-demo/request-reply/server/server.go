package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"log"
	"strings"
)

const endpoint = "tcp://127.0.0.1:5555"

func main() {
	server, err := zmq.NewSocket(zmq.REP)
	if err != nil {
		panic(fmt.Sprintf("fail to create a zmq socket: %v", err))
	}

	err = server.Bind(endpoint)
	if err != nil {
		panic(fmt.Sprintf("fail to bind a zmq socket: %v", err))
	}

	// receiveing requests
	for {
		// wait for the client request
		req, err := server.Recv(0)
		if err != nil {
			log.Fatalf("error while receiving a request: %v", err)
		}

		log.Printf("received message: %v", req)

		// sending back a response
		if _, err := server.Send(strings.ToUpper(req), 0); err != nil {
			log.Fatalf("error while sending a response: %v", err)
		}
	}

}
