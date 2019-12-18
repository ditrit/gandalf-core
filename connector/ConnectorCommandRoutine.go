package connector

import (
	"fmt"
	"errors"
	"strconv"
	"time"
	"gandalfgo/message"
	"gandalfgo/constant"
	"github.com/pebbe/zmq4"
)

type ConnectorCommandRoutine struct {
	Context												*zmq4.Context		
	ConnectorCommandSendToWorker              			*zmq4.Socket
	ConnectorCommandSendToWorkerConnection    			string
	ConnectorCommandReceiveFromAggregator           	*zmq4.Socket
	ConnectorCommandReceiveFromAggregatorConnections 	[]string
	ConnectorCommandSendToAggregator              		*zmq4.Socket
	ConnectorCommandSendToAggregatorConnections    		[]string
	ConnectorCommandReceiveFromWorker           		*zmq4.Socket
	ConnectorCommandReceiveFromWorkerConnection 		string
	Identity                              				string
	ConnectorMapUUIDCommandMessage		 				map[string][]message.CommandMessage					
	ConnectorMapWorkerCommands 			 				map[string][]string		
}

func NewConnectorCommandRoutine(identity, connectorCommandSendToWorkerConnection, connectorCommandReceiveFromWorkerConnection string, connectorCommandReceiveFromAggregatorConnections, connectorCommandSendToAggregatorConnections []string) (connectorCommandRoutine *ConnectorCommandRoutine) {
	connectorCommandRoutine = new(ConnectorCommandRoutine)
	
	connectorCommandRoutine.Identity = identity

	connectorCommandRoutine.Context, _ = zmq4.NewContext()
	connectorCommandRoutine.ConnectorCommandSendToWorkerConnection = connectorCommandSendToWorkerConnection
	connectorCommandRoutine.ConnectorCommandSendToWorker, _ = connectorCommandRoutine.Context.NewSocket(zmq4.ROUTER)
	connectorCommandRoutine.ConnectorCommandSendToWorker.SetIdentity(connectorCommandRoutine.Identity)
	connectorCommandRoutine.ConnectorCommandSendToWorker.Bind(connectorCommandRoutine.ConnectorCommandSendToWorkerConnection)
	fmt.Printf("connectorCommandSendToWorker bind : " + connectorCommandSendToWorkerConnection)

	connectorCommandRoutine.ConnectorCommandReceiveFromAggregatorConnections = connectorCommandReceiveFromAggregatorConnections
	connectorCommandRoutine.ConnectorCommandReceiveFromAggregator, _ = connectorCommandRoutine.Context.NewSocket(zmq4.DEALER)
	connectorCommandRoutine.ConnectorCommandReceiveFromAggregator.SetIdentity(connectorCommandRoutine.Identity)
	for _, connection := range connectorCommandRoutine.ConnectorCommandReceiveFromAggregatorConnections {
		connectorCommandRoutine.ConnectorCommandReceiveFromAggregator.Connect(connection)
		fmt.Printf("connectorCommandReceiveFromAggregator connect : " + connection)
	}

	connectorCommandRoutine.ConnectorCommandSendToAggregatorConnections = connectorCommandSendToAggregatorConnections
	connectorCommandRoutine.ConnectorCommandSendToAggregator, _ = connectorCommandRoutine.Context.NewSocket(zmq4.ROUTER)
	connectorCommandRoutine.ConnectorCommandSendToAggregator.SetIdentity(connectorCommandRoutine.Identity)
	for _, connection := range connectorCommandRoutine.ConnectorCommandSendToAggregatorConnections {
		connectorCommandRoutine.ConnectorCommandSendToAggregator.Connect(connection)
		fmt.Printf("connectorCommandSendToAggregator connect : " + connection)
	}


	connectorCommandRoutine.ConnectorCommandReceiveFromWorkerConnection = connectorCommandReceiveFromWorkerConnection
	connectorCommandRoutine.ConnectorCommandReceiveFromWorker, _ = connectorCommandRoutine.Context.NewSocket(zmq4.DEALER)
	connectorCommandRoutine.ConnectorCommandReceiveFromWorker.SetIdentity(connectorCommandRoutine.Identity)
	connectorCommandRoutine.ConnectorCommandReceiveFromWorker.Bind(connectorCommandRoutine.ConnectorCommandReceiveFromWorkerConnection)
	fmt.Printf("connectorCommandReceiveFromWorker bind : " + connectorCommandReceiveFromWorkerConnection)

	return
}

func (r ConnectorCommandRoutine) close() {	
	r.ConnectorCommandSendToWorker.Close()
	r.ConnectorCommandReceiveFromAggregator.Close()
	r.ConnectorCommandSendToAggregator.Close()
	r.ConnectorCommandReceiveFromWorker.Close()
	r.Context.Term()
}

func (r ConnectorCommandRoutine) reconnectToProxy() {

}

func (r ConnectorCommandRoutine) run() {
	go r.cleanCommandsByTimeout()

	poller := zmq4.NewPoller()
	poller.Add(r.ConnectorCommandSendToWorker, zmq4.POLLIN)
	poller.Add(r.ConnectorCommandReceiveFromAggregator, zmq4.POLLIN)
	poller.Add(r.ConnectorCommandSendToAggregator, zmq4.POLLIN)
	poller.Add(r.ConnectorCommandReceiveFromWorker, zmq4.POLLIN)

	command := [][]byte{}
	err := errors.New("")

	for {
	
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.ConnectorCommandSendToWorker:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}

				err = r.processCommandSendToWorker(command)
				if err != nil {
					panic(err)
				}

			case r.ConnectorCommandReceiveFromAggregator:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processCommandReceiveFromAggregator(command)
				if err != nil {
					panic(err)
				}

			case r.ConnectorCommandSendToAggregator:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processCommandSendAggregator(command)
				if err != nil {
					panic(err)
				}

			case r.ConnectorCommandReceiveFromWorker:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processCommandReceiveFromWorker(command)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

func (r ConnectorCommandRoutine) processCommandSendToWorker(command [][]byte) (err error) {
	commandMessage, err := message.DecodeCommandMessage(command[1])
	r.addCommands(commandMessage)
	go commandMessage.SendWith(r.ConnectorCommandReceiveFromAggregator, commandMessage.SourceConnector)
	return
}

func (r ConnectorCommandRoutine) processCommandReceiveFromAggregator(command [][]byte) (err error) {
	commandMessage, err := message.DecodeCommandMessage(command[1])
	r.ConnectorMapUUIDCommandMessage[commandMessage.Command] = append(r.ConnectorMapUUIDCommandMessage[commandMessage.Command], commandMessage)
	return
}

func (r ConnectorCommandRoutine) processCommandSendAggregator(command [][]byte) (err error) {
	commandMessage, err := message.DecodeCommandMessage(command[1])
	go commandMessage.SendCommandWith(r.ConnectorCommandReceiveFromWorker)
	return
}

func (r ConnectorCommandRoutine) processCommandReceiveFromWorker(command [][]byte) (err error) {
	workerSource := string(command[0])
	commandHeader := string(command[1])
    if commandHeader == constant.COMMAND_READY {
        //commandReady := decodeCommandReady(command[2])
        commandMessage, err := r.getCommandByWorkerCommands(workerSource)
        if err != nil {
        }
		go commandMessage.SendWith(r.ConnectorCommandSendToWorker, workerSource)
	} else if commandHeader == constant.COMMAND_VALIDATION_FUNCTIONS {
		commandFunction, _ := message.DecodeCommandFunction(command[2])
		result := r.validationCommands(workerSource, commandFunction.Functions)
        if result {
			r.ConnectorMapWorkerCommands[workerSource] = commandFunction.Functions
			commandFunctionReply := message.NewCommandFunctionReply(result)
			go commandFunctionReply.SendCommandFunctionReplyWith(r.ConnectorCommandSendToWorker)
        }
	} else {
		commandMessage, _ := message.DecodeCommandMessage(command[1])
		commandMessage.SourceWorker = workerSource
		go commandMessage.SendWith(r.ConnectorCommandSendToAggregator, workerSource)
	}
	return
}

func (r ConnectorCommandRoutine) getCommandByWorkerCommands(worker string) (commandMessage message.CommandMessage, err error) {
	
	var maxCommand string
	maxTimestamp := -1
	currentTimestamp := -1
	commandsWorker := r.ConnectorMapWorkerCommands[worker]
	var commands []message.CommandMessage
	
	for i, commandWorker := range commandsWorker {
		if currentCommandWorker, ok := r.ConnectorMapUUIDCommandMessage[commandWorker]; ok {
			commands[i] = currentCommandWorker[0]
		}
	}
	
	for _, command := range commands {
		currentTimestamp, _ = strconv.Atoi(command.Timestamp)
		if currentTimestamp >= maxTimestamp {
			maxTimestamp, _ = strconv.Atoi(command.Timestamp)
			maxCommand = command.Command
		}
	}
	
	commandMessage = r.ConnectorMapUUIDCommandMessage[maxCommand][0]
	delete(r.ConnectorMapUUIDCommandMessage, maxCommand);

	return 
}

func (r ConnectorCommandRoutine) getConnectorMapUUIDCommandMessage(command string) (commandMessage []message.CommandMessage) {
    if commandMessage, ok := r.ConnectorMapUUIDCommandMessage[command]; ok {
		if ok {
			return commandMessage
		}
	}
	return
}

func (r ConnectorCommandRoutine) validationCommands(workerSource string, commands []string) (result bool) {
	//TODO
	result = true

	return
}

func (r ConnectorCommandRoutine) addCommands(commandMessage message.CommandMessage) {
	if _, ok := r.ConnectorMapUUIDCommandMessage[commandMessage.Uuid]; ok {
		if !ok {
			r.ConnectorMapUUIDCommandMessage[commandMessage.Uuid] = append(r.ConnectorMapUUIDCommandMessage[commandMessage.Uuid], commandMessage)
		}
	}
}

func (r ConnectorCommandRoutine) cleanCommandsByTimeout() {
	maxTimeout := 0
	currentTimestamp := -1
	currentTimeout := -1
	for {
		for uuid, commandMessageSlice := range r.ConnectorMapUUIDCommandMessage { 
			for _, commandMessage := range commandMessageSlice {
				currentTimestamp, _ = strconv.Atoi(commandMessage.Timestamp)
				currentTimeout, _ = strconv.Atoi(commandMessage.Timeout)

				if currentTimestamp - currentTimeout == 0 {
					delete(r.ConnectorMapUUIDCommandMessage, uuid) 	
				} else {
					if currentTimeout >= maxTimeout {
						maxTimeout = currentTimeout
					}
				}
			}
		}
		time.Sleep(time.Duration(maxTimeout) * time.Millisecond)
	}
}