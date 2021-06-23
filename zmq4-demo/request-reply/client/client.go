package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"log"
)

const endpoint = "tcp://127.0.0.1:5555"

func main() {
	client, err := zmq.NewSocket(zmq.REQ)
	if err != nil {
		panic(fmt.Sprintf("fail to create a zmq socket: %v", err))
	}

	err = client.Connect(endpoint)
	if err != nil {
		panic(fmt.Sprintf("fail to bind a zmq socket: %v", err))
	}


	// sending a request
	req := "hello"
	if _, err := client.Send(req, 0); err != nil {
		log.Fatalf("error while sending a request: %v", err)
	}
	log.Printf("client send the request: %v", req)

	res, err := client.Recv(0)
	if err != nil {
		log.Fatalf("fail to receive respone")
	}

	log.Printf("response: %v", res)
}

