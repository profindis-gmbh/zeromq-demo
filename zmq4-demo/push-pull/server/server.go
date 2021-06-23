package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"time"
)

const endpoint = "tcp://127.0.0.1:5555"

func main() {
	// Todo learn more about the ztx context; apparently it is used to define/read zmq configuration
	zctx, err := zmq.NewContext()
	if err != nil {
		panic(fmt.Sprintf("fail to create a zmq context: %v", err))
	}

	socket, err := zctx.NewSocket(zmq.PUSH)
	if err != nil {
		panic(fmt.Sprintf("fail to create a zmq socket: %v", err))
	}

	err = socket.Bind(endpoint)
	if err != nil {
		panic(fmt.Sprintf("fail to bind a zmq socket: %v", err))
	}
	// send jobs to the workers
	for i := 0; i < 10; i++ {
		time.Sleep(500 *time.Millisecond) // Todo: check why the first message fail to be deleved if this is placed at the end of the for-loop
		_, err := socket.Send(fmt.Sprintf("job %v", i), zmq.DONTWAIT)
		fmt.Printf("error: %v\n", err)
	}

}
