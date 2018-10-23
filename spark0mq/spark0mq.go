package spark0mq

import (
	"github.com/golang/protobuf/proto"
	"github.com/willtoth/USB-BLDC-TOOL/sparkmax"

	zmq "github.com/pebbe/zmq4"

	"fmt"
)

type Spark0mqServer struct {
	port      int
	verbosity int
	running   bool
}

func Spark0mqStart(port int) (Spark0mqServer, error) {
	fmt.Printf("Starting smark0mq REQ on port: %d\r\n", port)
	go spark0mqREQ(port)
	return Spark0mqServer{port, 0, true}, nil
}

func (s *Spark0mqServer) Stop() {
	fmt.Println("Sending STOP command to server -- not implemented")
}

func (s *Spark0mqServer) IsRunning() bool {
	return s.running
}

func spark0mqREQ(port int) {
	//  Socket to talk to clients
	responder, err := zmq.NewSocket(zmq.REP)
	if err != nil {
		panic(err)
	}
	defer responder.Close()
	defer fmt.Println("Responder.Close called")

	bindStr := fmt.Sprintf("tcp://*:%d", port)
	responder.Bind(bindStr)

	for {
		//  Wait for next request from client
		msg, _ := responder.RecvBytes(0)

		fmt.Println("Received ", msg)

		resp, err := parseMessage(msg)
		if err != nil {
			fmt.Println(err.Error())
		}
		//  Send reply back to client
		responder.SendBytes(resp, 0)
		fmt.Println("Sent ", resp)
	}
}

func parseMessage(msg []byte) (rawBytes []byte, err error) {
	req := sparkmax.RequestWire{}
	root := &sparkmax.RootResponse{}
	var resp sparkmax.ResponseWire

	if err := proto.Unmarshal(msg, &req); err != nil {
		err = fmt.Errorf("Parse failure in server: %s", err.Error())
		root.Error = err.Error()
		resp.Resp = &sparkmax.ResponseWire_Root{Root: root}
	} else {
		resp, err = sparkmax.RunCommand(req)
	}

	rawBytes, err = proto.Marshal(&resp)
	fmt.Println(rawBytes)
	return
}
