package cluster

import (
	"fmt"
    "net/http"
	zmq "github.com/zeromq/goczmq"
)

type ClusterCaptureWorkerRoutine struct {
	workerCaptureCommandReceiveCL2W           zmq.Sock
	workerCaptureCommandReceiveCL2WConnection string
	workerCaptureEventReceiveCL2W             zmq.Sock
	workerCaptureEventReceiveCL2WConnection   string
	identity                                  string
}

func (r ClusterCaptureWorkerRoutine) New(identity, workerCaptureCommandReceiveCL2WConnection, workerCaptureEventReceiveC2WConnection string, topics *string) err error {
	r.identity = identity

	r.workerCaptureCommandReceiveCL2WConnection = workerCaptureCommandReceiveCL2WConnection
	r.workerCaptureCommandReceiveCL2W = zmq.NewDealer(workerCaptureCommandReceiveCL2WConnection)
	r.workerCaptureCommandReceiveCL2W.Identity(r.identity)
	fmt.Printf("workerCaptureCommandReceiveCL2W connect : " + workerCaptureCommandReceiveCL2WConnection)

	r.workerCaptureEventReceiveCL2WConnection = workerCaptureEventReceiveCL2WConnection
	r.workerCaptureEventReceiveC2W = zmq.NewSub(workerCaptureEventReceiveCL2WConnection)
	r.workerCaptureEventReceiveC2W.Identity(r.identity)
	fmt.Printf("workerCaptureEventReceiveC2W connect : " + workerCaptureEventReceiveCL2WConnection)
}

func (r ClusterCaptureWorkerRoutine) close() err error {
	r.workerCaptureCommandReceiveC2W.close()
	r.workerCaptureEventReceiveC2W.close()
	r.Context.close()
}

func (r ClusterCaptureWorkerRoutine) run() err error {
	pi := zmq.PollItems{
		zmq.PollItem{Socket: workerCaptureCommandReceiveCL2W, Events: zmq.POLLIN},
		zmq.PollItem{Socket: workerCaptureEventReceiveC2W, Events: zmq.POLLIN}}

	var command = [][]byte{}
	var event = [][]byte{}

	for {
		r.sendReadyCommand()

		_, _ = zmq.Poll(pi, -1)

		switch {
		case pi[0].REvents&zmq.POLLIN != 0:

			command, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommand(command)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq.POLLIN != 0:

			event, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEvent(event)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println("done")
}

func (r ClusterCaptureWorkerRoutine) processCommand(command [][]byte) err error {
	command = r.updateHeaderCommand(command)
    response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(command[1]))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    }
}

func (r ClusterCaptureWorkerRoutine) updateHeaderCommand(command [][]byte) err error {
}

func (r ClusterCaptureWorkerRoutine) processEvent(event [][]byte) err error {
	event = r.updateHeaderEvent(event)
    response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(event[0]))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    }
}

func (r ClusterCaptureWorkerRoutine) updateHeaderEvent(event [][]byte) err error {
}
