package spark0mq

import (
	"time"

	"github.com/willtoth/USB-BLDC-TOOL/sparkusb"

	zmq "github.com/pebbe/zmq4"

	"fmt"
)

type Spark0mqServer struct {
	port      int
	verbosity int
	running   bool
}

func Spark0mqStart(port int) (Spark0mqServer, error) {
	fmt.Println("Starting smark0mq REQ")
	go spark0mqREQ(port)
	return Spark0mqServer{port, 0, true}, nil
}

func (s *Spark0mqServer) Stop() {
	fmt.Println("Sending STOP command to server")

	//  Socket to talk to server
	requester, _ := zmq.NewSocket(zmq.REQ)
	defer requester.Close()
	bindStr := fmt.Sprintf("tcp://localhost:%d", s.port)
	requester.Connect(bindStr)

	ctrlMsg := sparkusb.ControlRequest{Req: ""}
	msg := sparkusb.RequestWire{Req: &sparkusb.RequestWire_Control{&ctrlMsg}}

	// send quit
	requester.Send(msg.String(), 0)
}

func (s *Spark0mqServer) IsRunning() bool {
	return s.running
}

func spark0mqREQ(port int) {
	//  Socket to talk to clients
	responder, _ := zmq.NewSocket(zmq.REP)
	defer responder.Close()

	bindStr := fmt.Sprintf("tcp://*:%d", port)
	responder.Bind(bindStr)

	for {
		//  Wait for next request from client
		msg, _ := responder.Recv(0)

		fmt.Println("Received ", msg)

		time.Sleep(100 * time.Millisecond)

		//  Send reply back to client
		reply := "World"
		responder.Send(reply, 0)
		fmt.Println("Sent ", reply)
	}
}
