package sendergrpc

import (
	"context"
	"fmt"
	pb "gandalf-go/grpc"
	"gandalf-go/message"

	"google.golang.org/grpc"
)

type SenderCommandGrpc struct {
	SenderCommandGrpcConnection string
	Identity                    string
	client                      pb.ConnectorCommandClient
}

func NewSenderCommandGrpc(identity, senderCommandGrpcConnection string) (senderCommandGrpc *SenderCommandGrpc) {
	senderCommandGrpc = new(SenderCommandGrpc)
	senderCommandGrpc.Identity = identity
	senderCommandGrpc.SenderCommandGrpcConnection = senderCommandGrpcConnection
	conn, err := grpc.Dial(senderCommandGrpc.SenderCommandGrpcConnection, grpc.WithInsecure())
	if err != nil {
	}
	senderCommandGrpc.client = pb.NewConnectorCommandClient(conn)
	fmt.Println("senderCommandGrpc connect : " + senderCommandGrpc.SenderCommandGrpcConnection)
	return
}

func (r SenderCommandGrpc) SendCommand(contextCommand, timeout, uuid, connectorType, commandType, command, payload string) (commandMessageUUID message.CommandMessageUUID) {
	commandMessage := new(pb.CommandMessage)
	commandMessage.Context = contextCommand
	commandMessage.Timeout = timeout
	commandMessage.Uuid = uuid
	commandMessage.ConnectorType = connectorType
	commandMessage.CommandType = command
	commandMessage.Command = command
	commandMessage.Payload = payload

	CommandMessageUUIDGrpc, _ := r.client.SendCommandMessage(context.Background(), commandMessage)
	commandMessageUUID = message.CommandMessageUUIDFromGrpc(CommandMessageUUIDGrpc)
	return

}

func (r SenderCommandGrpc) SendCommandReply(commandMessage message.CommandMessage, reply, payload string) *pb.Empty {
	commandMessageReply := new(pb.CommandMessageReply)
	commandMessageReply.SourceAggregator = commandMessage.SourceAggregator
	commandMessageReply.SourceConnector = commandMessage.SourceConnector
	commandMessageReply.SourceWorker = commandMessage.SourceWorker
	commandMessageReply.DestinationAggregator = commandMessage.DestinationAggregator
	commandMessageReply.DestinationConnector = commandMessage.DestinationConnector
	commandMessageReply.DestinationWorker = commandMessage.DestinationWorker
	commandMessageReply.Tenant = commandMessage.Tenant
	commandMessageReply.Token = commandMessage.Token
	commandMessageReply.Context = commandMessage.Context
	commandMessageReply.Timeout = commandMessage.Timeout
	commandMessageReply.Timestamp = commandMessage.Timestamp
	commandMessageReply.Uuid = commandMessage.Uuid
	commandMessageReply.Reply = reply
	commandMessageReply.Payload = payload

	empty, _ := r.client.SendCommandMessageReply(context.Background(), commandMessageReply)
	return empty
}
