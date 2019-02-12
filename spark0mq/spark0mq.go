// Copyright © 2018 REV Robotics LLC (support@revrobotics.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spark0mq

import (
	"log"

	"github.com/REVrobotics/SPARK-MAX-Server/sparkmax"
	"github.com/golang/protobuf/proto"

	zmq "github.com/pebbe/zmq4"

	"fmt"
)

type Spark0mqServer struct {
	port      int
	verbosity int
	running   bool
	publisher chan []byte
}

func Spark0mqStart(port, verbosity int) (Spark0mqServer, error) {
	fmt.Printf("Starting smark0mq REQ on port: %d\r\n", port)

	var server Spark0mqServer

	server.port = port
	server.verbosity = verbosity
	server.running = true
	server.publisher = make(chan []byte, 32)

	go spark0mqREQ(port, verbosity)
	return server, nil
}

func (s *Spark0mqServer) Stop() {
	fmt.Println("Sending STOP command to server -- not implemented")
}

func (s *Spark0mqServer) IsRunning() bool {
	return s.running
}

func spark0mqREQ(port, verbosity int) {
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

		if verbosity >= 4 {
			log.Println("Received ", msg)
		}

		resp, err := parseMessage(msg, verbosity)
		if err != nil {
			fmt.Println(err.Error())
		}
		//  Send reply back to client
		responder.SendBytes(resp, 0)

		if verbosity >= 4 {
			log.Println("Sent ", resp)
		}
	}
}

func parseMessage(msg []byte, verbosity int) (rawBytes []byte, err error) {
	req := sparkmax.RequestWire{}
	root := &sparkmax.RootResponse{}
	var resp sparkmax.ResponseWire

	if err := proto.Unmarshal(msg, &req); err != nil {
		err = fmt.Errorf("Parse failure in server: %s", err.Error())
		root.Error = err.Error()
		resp.Resp = &sparkmax.ResponseWire_Root{Root: root}
	} else {
		if verbosity >= 3 {
			log.Println("ZeroMQ Recieved: ", req.String())
		}
		resp, err = sparkmax.RunCommand(req)
	}

	if verbosity >= 3 {
		log.Println("ZeroMQ Response: ", resp.String())
	}

	rawBytes, err = proto.Marshal(&resp)
	//fmt.Println(rawBytes)
	return
}

func (s *Spark0mqServer) PublishMessage(msg []byte) {
	s.publisher <- msg
}

func (s *Spark0mqServer) spark0mqPUB(port int) {
	//  Socket to talk to clients
	pubSocket, err := zmq.NewSocket(zmq.PUB)
	if err != nil {
		panic(err)
	}
	defer pubSocket.Close()
	defer fmt.Println("pubSocket.Close called")

	bindStr := fmt.Sprintf("tcp://*:%d", port)
	pubSocket.Bind(bindStr)

	for {
		for msg := range s.publisher {
			//  Wait for next request from client
			_, err := pubSocket.SendBytes(msg, 0)

			if err != nil {
				fmt.Println("Error when attemping to publish data, ", err.Error())
			}
		}
	}
}
