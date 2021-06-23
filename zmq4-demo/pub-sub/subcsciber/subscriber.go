package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

const endpoint = "tcp://127.0.0.1:5555"
const topic = "topic-a"

func main() {
	zctx, err := zmq.NewContext()
	if err != nil {
		panic(fmt.Sprintf("fail to create a zmq context: %v", err))
	}

	subscriber, err := zctx.NewSocket(zmq.SUB)
	if err != nil {
		panic(fmt.Sprintf("fail to create a zmq socket: %v", err))
	}

	err = subscriber.Connect(endpoint)
	if err != nil {
		panic(fmt.Sprintf("fail to bind a zmq socket: %v", err))
	}

	err = subscriber.SetSubscribe(topic)
	if err != nil {
		panic(fmt.Sprintf("fail to set the subscribe topic: %v", err))
	}
	for {
		_, err := subscriber.Recv(0)
		if err != nil {
			fmt.Printf("fail to received the topic frame: %v", err)
		}

		msg, err := subscriber.Recv(0)
		fmt.Printf("error: %v, message: %v\n", err, msg)
	}

}
