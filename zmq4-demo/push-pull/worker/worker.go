package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)


func main() {
	endpoint := "tcp://127.0.0.1:5555"
	zctx, err := zmq.NewContext()
	if err != nil {
		panic(fmt.Sprintf("fail to create a zmq context: %v", err))
	}

	socket, err := zctx.NewSocket(zmq.PULL)
	if err != nil {
		panic(fmt.Sprintf("fail to create a zmq socket: %v", err))
	}

	err = socket.Connect(endpoint)
	if err != nil {
		panic(fmt.Sprintf("fail to bind a zmq socket: %v", err))
	}

	for {
		msg, err := socket.Recv(0)
		fmt.Printf("error: %v, message: %v\n", err, msg)
	}


}
