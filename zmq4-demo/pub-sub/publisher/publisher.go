package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"time"
)

const endpoint = "tcp://127.0.0.1:5555"
const topic = "topic-a"

func main() {
	zctx, err := zmq.NewContext()
	if err != nil {
		panic(fmt.Sprintf("fail to create a zmq context: %v", err))
	}

	socket, err := zctx.NewSocket(zmq.PUB)
	if err != nil {
		panic(fmt.Sprintf("fail to create a zmq socket: %v", err))
	}

	err = socket.Bind(endpoint)
	if err != nil {
		panic(fmt.Sprintf("fail to bind a zmq socket: %v", err))
	}
	// send jobs to the subscribers
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		if _, err := socket.Send(topic, zmq.SNDMORE); err != nil {
			fmt.Printf("topic frame error: %v\n", err)
		}
		_, err = socket.Send(fmt.Sprintf("data %v", i), zmq.DONTWAIT)
		fmt.Printf("error: %v\n", err)
	}

}
