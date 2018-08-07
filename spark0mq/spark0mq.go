package spark0mq

import (
	"github.com/golang/protobuf/proto"
	sparkusb "github.com/willtoth/USB-BLDC-TOOL/sparkusb"

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
	fmt.Println("Sending STOP command to server -- not implemented")
	/*

		//  Socket to talk to server
		requester, _ := zmq.NewSocket(zmq.REQ)
		defer requester.Close()
		bindStr := fmt.Sprintf("tcp://localhost:%d", s.port)
		requester.Connect(bindStr)

		ctrlMsg := sparkusb.ParameterRequest{Parameter: 0}
		msg := sparkusb.RequestWire{Req: &sparkusb.RequestWire_Parameter{&ctrlMsg}}

		val, _ := proto.Marshal(&msg)

		// send quit
		requester.SendBytes(val, 0)

		bytes, _ := requester.RecvBytes(0)

		resp := sparkusb.ResponseWire{}
		proto.Unmarshal(bytes, &resp)
		fmt.Println("Resp: ", resp.GetParameter().Value)
	*/
}

func (s *Spark0mqServer) IsRunning() bool {
	return s.running
}

func spark0mqREQ(port int) {
	//  Socket to talk to clients
	responder, _ := zmq.NewSocket(zmq.REP)
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
			fmt.Println(err)
		}
		//  Send reply back to client
		responder.SendBytes(resp, 0)
		fmt.Println("Sent ", resp)
	}
}

func parseMessage(msg []byte) (rawBytes []byte, err error) {
	req := &sparkusb.RequestWire{}
	root := &sparkusb.RootResponse{}
	var resp sparkusb.ResponseWire

	if err := proto.Unmarshal(msg, req); err != nil {
		err = fmt.Errorf("Parse failure in server: %s", err.Error())
		root.Error = err.Error()
		resp.Resp = &sparkusb.ResponseWire_Root{Root: root}
	} else {
		switch x := req.Req.(type) {
		//Get or Set parameter
		case *sparkusb.RequestWire_Parameter:
			param := x.Parameter.Parameter
			val := x.Parameter.Value
			paramResp := sparkusb.ParameterResponse{}

			//Get
			if val == "" {
				r, err := sparkusb.GetParameter(&sparkusb.ParameterRequest{Parameter: param})
				if err != nil {
					tmp := sparkusb.RootResponse{Error: err.Error()}
					paramResp.Root = &tmp
				}
				paramResp.Value = r.Value
			} else {
				_, err := sparkusb.SetParameter(&sparkusb.ParameterRequest{Parameter: param, Value: val})
				if err != nil {
					paramResp.Root.Error = err.Error()
				}
			}

			resp.Resp = &sparkusb.ResponseWire_Parameter{Parameter: &paramResp}

		case *sparkusb.RequestWire_Setpoint:
			setpoint := x.Setpoint.Setpoint
			enabled := x.Setpoint.Enable
			req := sparkusb.SetpointRequest{Setpoint: setpoint, Enable: enabled}
			r, err := sparkusb.Setpoint(&req)
			if err != nil {
				r.Root.Error = err.Error()
			}
			resp.Resp = &sparkusb.ResponseWire_Setpoint{Setpoint: r}

		//Invalid message
		default:
			root.Error = fmt.Sprintf("Message has unexpected type %T", x)
			err = fmt.Errorf(root.Error)
			resp.Resp = &sparkusb.ResponseWire_Root{Root: root}
		}
	}

	rawBytes, err = proto.Marshal(&resp)
	return
}
