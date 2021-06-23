package main

import (
	"log"
	"time"

	"github.com/zeromq/goczmq"
)

const endpoint = "tcp://127.0.0.1:5555"

// simulate a message processor
func msgProcessor(cmd string) (string, error) {
	time.Sleep(500 * time.Millisecond)
	return cmd, nil
}

func main() {
	router, err := goczmq.NewRouter(endpoint)
	if err != nil {
		log.Fatalf("fail to create a zmq router: %v", err)
	}
	defer router.Destroy()

	log.Println("router created and bound")
	for {
		request, err := router.RecvMessage()
		if err != nil {
			log.Printf("failed while receiving a request with error: %v\n", err)
			continue
		}

		log.Printf("received a message (%s) from %v\n", request[1], request[0])
		// process the command
		resp, err := msgProcessor(string(request[1]))
		if err != nil {
			// should send the error over a custom message
			// for now just log it and abort
			log.Printf("fail to process the command with error: %v\n", err)
			continue
		}

		// Send a reply. First we send the routing frame, which
		// lets the dealer know which client to send the message to .
		// The FlagMore flag tells the router there will be more
		// frames in this message.
		err = router.SendFrame(request[0], goczmq.FlagMore)
		if err != nil {
			log.Printf("could not send the routing frame to the dealer with error: %v\n", err)
			continue
		}

		// send the response over a custom message back to the client
		err = router.SendFrame([]byte(resp), goczmq.FlagNone)
		if err != nil {
			log.Printf("could not send the response to the client with error: %v\n", err)
			continue
		}
	}

}
