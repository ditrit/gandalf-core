package message

import (
	"fmt"
	"constant"
	msgpack "github.com/shamaton/msgpack"
)

type CommandMessage struct {
	sourceAggregator    string
	sourceConnector string
	sourceWorker   string
	destinationAggregator    string
    destinationConnector    string
    destinationWorker string
    tenant   string
    token    string
    context    string
    timeout string
    timestamp   string
    major    string
    minor    string
	uuid string
	connectorType string
    commandType   string
    command    string
    payload    string
}

func (c CommandMessage) New(context, timeout, uuid, connectorType, commandType, command, payload string) err error {
	c.context = context
	c.timeout = timeout
	c.uuid = uuid
	c.connectorType = connectorType
	c.commandType = commandType
	c.command = command
	c.payload = payload
	c.timestamp = time.Now()
}

func (c CommandMessage) sendWith(socket zmq.Sock, header string) isSend bool {
	validation = cr.sendHeaderWith(socket, header)
	validation += cr.sendCommandWith(socket)
	isSend := validation > 0 ? true : false
	return
}

func (c CommandMessage) sendHeaderWith(socket zmq.Sock, header string) isSend bool {
	validation = zmq_send(socket, header, ZMQ_SNDMORE);
	isSend := validation > 0 ? true : false
	return
}

func (c CommandMessage) sendCommandWith(socket zmq.Sock) isSend bool {
	validation = zmq_send(socket, encode(c), 0);
	isSend := validation > 0 ? true : false
	return
}

func (c CommandMessage) from(command []byte) err error {
	c.sourceAggregator = command[0]
	c.sourceConnector = command[1]
	c.sourceWorker = command[2]
	c.destinationAggregator = command[3]
    c.destinationConnector = command[4]
    c.destinationWorker = command[5]
    c.tenant = command[6]
    c.token = command[7]
    c.context = command[8]
    c.timeout = command[9]
    c.timestamp = command[10]
    c.major = command[11]
    c.minor = command[12]
	c.uuid = command[13]
	c.connectorType = command[14]
    c.commandType = command[15]
    c.command = command[16]
    c.payload = command[17]
}

//

type CommandReply struct {
	sourceAggregator    string
	sourceConnector string
	sourceWorker   string
	destinationAggregator    string
    destinationConnector    string
    destinationWorker string
    tenant   string
    token    string
    context    string
    timeout string
    timestamp   string
	uuid string
	reply    string
    payload    string
}

func (cr CommandReply) sendWith(socket zmq.Sock, header string) isSend bool {
	validation = cr.sendHeaderWith(socket, header)
	validation += cr.sendCommandReplyWith(socket)
	isSend := validation > 0 ? true : false
	return
}

func (cr CommandReply) sendHeaderCommandReplyWith(socket zmq.Sock, header string) isSend bool {
	validation = zmq_send(socket, header, ZMQ_SNDMORE);
	isSend := validation > 0 ? true : false
	return
}

func (cr CommandReply) sendCommandReplyWith(socket zmq.Sock) isSend bool {
	validation = zmq_send(socket, encode(cr), 0);
	isSend := validation > 0 ? true : false
	return
}

func (cr CommandReply) from(commandMessage CommandMessage, reply, payload string) {
	cr.sourceAggregator = commandMessage.sourceAggregator
	cr.sourceConnector = commandMessage.sourceConnector
	cr.sourceWorker = commandMessage.sourceWorker
	cr.destinationAggregator = commandMessage.destinationAggregator
    cr.destinationConnector = commandMessage.destinationConnector
    cr.destinationWorker = commandMessage.destinationWorker
    cr.tenant = commandMessage.tenant
    cr.token = commandMessage.token
    cr.context = commandMessage.context
    cr.timeout = commandMessage.timeout
    cr.timestamp = commandMessage.timestamp
	cr.uuid = commandMessage.uuid
	cr.reply = reply
    cr.payload = payload
}

//

type CommandFunction struct {
	functions    []string
}

func (cf CommandFunction) New(functions []string) err error {
	cf.functions = functions
}

func (cf CommandFunction) sendWith(socket zmq.Sock) isSend bool {
	validation = zmq_send(socket, constant.COMMAND_VALIDATION_FUNCTIONS, ZMQ_SNDMORE);
	validation += zmq_send(socket, encode(cf), 0);
	isSend := validation > 0 ? true : false
	return
}

//

type CommandFunctionReply struct {
	validation bool
}

func (cfr CommandFunctionReply) New(validation bool) err error {
	cfr.validation = validation
}

func (cfr CommandFunctionReply) sendWith(socket zmq.Sock, header string) isSend bool {
	validation = cfr.sendHeaderWith(socket, header)
	validation += cfr.sendCommandCommandsEventsReplyWith(socket)
	isSend := validation > 0 ? true : false
	return
}

func (cfr CommandFunctionReply) sendHeaderWith(socket zmq.Sock, header string) isSend bool {
	validation = zmq_send(socket, header, ZMQ_SNDMORE);
	isSend := validation > 0 ? true : false
	return
}

func (cfr CommandFunctionReply) sendCommandFunctionReplyWith(socket zmq.Sock) isSend bool {
	validation = zmq_send(socket, constant.COMMAND_VALIDATION_FUNCTIONS_REPLY, ZMQ_SNDMORE);
	validation += zmq_send(socket, encode(ccer), 0);
	isSend := validation > 0 ? true : false
	return
}

//

type CommandReady struct {
	// ???
}

func (cry CommandReady) New() err error {
}

func (cry CommandReady) sendWith(socket zmq.Sock) isSend bool {
	validation = zmq_send(socket, constant.COMMAND_READY, ZMQ_SNDMORE);
	validation += zmq_send(socket, encode(cry), 0);
	isSend := validation > 0 ? true : false
	return
}

//

func encode() (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(c)
	if err != nil {
		commandError = fmt.Errorf("Command %s", err)
		return
	}
	return
}

func decodeCommand(bytesContent []byte) (command Command, commandError error) {
	err := msgpack.Decode(bytesContent, command)
	if err != nil {
		commandError = fmt.Errorf("Command %s", err)
		return
	}
	return
}	

func decodeCommandReply(bytesContent []byte) (commandReply CommandReply, commandError error) {
	err := msgpack.Decode(bytesContent, commandReply)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}

func decodeCommandReady(bytesContent []byte) (commandReady CommandReady, commandError error) {
	err := msgpack.Decode(bytesContent, CommandReady)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}

func decodeCommandAck(bytesContent []byte) (commandAck CommandAck, commandError error) {
	err := msgpack.Decode(bytesContent, commandAck)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}

func decodeCommandFunction(bytesContent []byte) (commandFunctions CommandFunctions, commandError error) {
	err := msgpack.Decode(bytesContent, commandFunctions)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}

func decodeCommandFunctionReply(bytesContent []byte) (commandFunctionReply CommandFunctionReply, commandError error) {
	err := msgpack.Decode(bytesContent, commandFunctionReply)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}