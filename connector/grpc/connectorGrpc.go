//Package grpc :
package grpc

import (
	"context"
	pb "core/grpc"
	"errors"
	"log"
	"net"
	"shoset/msg"
	sn "shoset/net"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var sendIndex = 0

// ConnectorGrpc : ConnectorGrpc struct.
type ConnectorGrpc struct {
	GrpcConnection string
	Shoset         sn.Shoset
	//MapWorkerIterators map[string][]*msg.Iterator
	MapIterators      map[string]*msg.Iterator
	MapCommandChannel map[string]chan msg.Message
	EventChannel      chan msg.Message
	ValidationChannel chan msg.Message
	timeoutMax        int64
}

// NewConnectorGrpc : ConnectorGrpc constructor.
func NewConnectorGrpc(grpcConnection string, timeoutMax int64, shoset *sn.Shoset) (connectorGrpc ConnectorGrpc, err error) {
	connectorGrpc.Shoset = *shoset
	connectorGrpc.GrpcConnection = grpcConnection
	connectorGrpc.timeoutMax = timeoutMax
	//connectorGrpc.MapWorkerIterators = make(map[string][]*msg.Iterator)
	connectorGrpc.MapIterators = make(map[string]*msg.Iterator)
	connectorGrpc.MapCommandChannel = make(map[string]chan msg.Message)
	connectorGrpc.EventChannel = make(chan msg.Message)
	connectorGrpc.ValidationChannel = make(chan msg.Message)

	return
}

// StartGrpcServer : ConnectorGrpc start.
func (r ConnectorGrpc) StartGrpcServer() {
	lis, err := net.Listen("tcp", r.GrpcConnection)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	connectorGrpcServer := grpc.NewServer()

	pb.RegisterConnectorCommandServer(connectorGrpcServer, &r)
	pb.RegisterConnectorEventServer(connectorGrpcServer, &r)
	connectorGrpcServer.Serve(lis)
}

//SendCommandMessage : Connector send command function.
func (r ConnectorGrpc) SendCommandMessage(ctx context.Context, in *pb.CommandMessage) (commandMessageUUID *pb.CommandMessageUUID, err error) {
	log.Println("Handle send command")

	cmd := pb.CommandFromGrpc(in)
	cmd.Tenant = r.Shoset.Context["tenant"].(string)
	shosets := sn.GetByType(r.Shoset.ConnsByAddr, "a")

	if len(shosets) != 0 {
		if cmd.GetTimeout() > r.timeoutMax {
			cmd.Timeout = r.timeoutMax
		}

		iteratorMessage, _ := r.CreateIteratorEvent(ctx, new(pb.Empty))
		iterator := r.MapIterators[iteratorMessage.GetId()]

		go r.runIterator(cmd.GetUUID(), "validation", iterator, r.ValidationChannel)

		notSend := true
		for notSend {
			index := getSendIndex(shosets)
			shosets[index].SendMessage(cmd)
			log.Printf("%s : send command %s to %s\n", r.Shoset.GetBindAddr(), cmd.GetCommand(), shosets[index])

			timeoutSend := time.Duration((int(cmd.GetTimeout()) / len(shosets)))

			messageChannel := <-r.ValidationChannel
			log.Printf("%s : receive validation event for command %s to %s\n", r.Shoset.GetBindAddr(), cmd.GetCommand(), shosets[index])

			if messageChannel != nil {
				notSend = false
				break
			}

			time.Sleep(timeoutSend)
		}

		if notSend {
			return nil, nil
		}

		commandMessageUUID = &pb.CommandMessageUUID{UUID: cmd.UUID}
	} else {
		log.Println("can't find aggregators to send")
		err = errors.New("can't find aggregators to send")
	}

	return commandMessageUUID, err
}

//WaitCommandMessage : Connector wait command function.
func (r ConnectorGrpc) WaitCommandMessage(ctx context.Context, in *pb.CommandMessageWait) (commandMessage *pb.CommandMessage, err error) {
	log.Println("Handle wait command")

	iterator := r.MapIterators[in.GetIteratorId()]

	go r.runIterator(in.GetValue(), "cmd", iterator, r.MapCommandChannel[in.GetIteratorId()])

	messageChannel := <-r.MapCommandChannel[in.GetIteratorId()]
	commandMessage = pb.CommandToGrpc(messageChannel.(msg.Command))

	return
}

//SendEventMessage : Connector send event function.
func (r ConnectorGrpc) SendEventMessage(ctx context.Context, in *pb.EventMessage) (*pb.Empty, error) {
	log.Println("Handle send event")

	evt := pb.EventFromGrpc(in)
	evt.Tenant = r.Shoset.Context["tenant"].(string)
	thisOne := r.Shoset.GetBindAddr()

	r.Shoset.ConnsByAddr.Iterate(
		func(key string, val *sn.ShosetConn) {
			if key != r.Shoset.GetBindAddr() && key != thisOne && val.ShosetType == "a" {
				val.SendMessage(evt)
				log.Printf("%s : send event %s to %s\n", thisOne, evt.GetEvent(), val)
			}
		},
	)

	return &pb.Empty{}, nil
}

//WaitEventMessage : Connector wait event function.
func (r ConnectorGrpc) WaitEventMessage(ctx context.Context, in *pb.EventMessageWait) (messageEvent *pb.EventMessage, err error) {
	log.Println("Handle wait event")

	iterator := r.MapIterators[in.GetIteratorId()]

	go r.runIterator(in.GetEvent(), "evt", iterator, r.EventChannel)

	messageChannel := <-r.EventChannel
	messageEvent = pb.EventToGrpc(messageChannel.(msg.Event))

	return
}

//WaitTopicMessage : Connector wait event by topic function.
func (r ConnectorGrpc) WaitTopicMessage(ctx context.Context, in *pb.TopicMessageWait) (messageEvent *pb.EventMessage, err error) {
	log.Println("Handle wait event by topic")

	iterator := r.MapIterators[in.GetIteratorId()]

	go r.runIterator(in.GetTopic(), "topic", iterator, r.EventChannel)

	messageChannel := <-r.EventChannel
	messageEvent = pb.EventToGrpc(messageChannel.(msg.Event))

	return
}

//TODO REFACTORING

//CreateIteratorCommand : Connector create command iterator function.
func (r ConnectorGrpc) CreateIteratorCommand(ctx context.Context, in *pb.Empty) (iteratorMessage *pb.IteratorMessage, err error) {
	log.Println("Handle create iterator command")

	iterator := msg.NewIterator(r.Shoset.Queue["cmd"])
	index := uuid.New()
	log.Printf("Create new iterator command: %s", index)

	//r.MapWorkerIterators[index.String()] = append(r.MapWorkerIterators[index.String()], iterator)
	r.MapIterators[index.String()] = iterator
	r.MapCommandChannel[index.String()] = make(chan msg.Message)
	iteratorMessage = &pb.IteratorMessage{Id: index.String()}

	return
}

//CreateIteratorEvent : Connector create event iterator function.
func (r ConnectorGrpc) CreateIteratorEvent(ctx context.Context, in *pb.Empty) (iteratorMessage *pb.IteratorMessage, err error) {
	log.Println("Handle create iterator event")

	iterator := msg.NewIterator(r.Shoset.Queue["evt"])
	index := uuid.New()
	log.Printf("Create new iterator event: %s", index)

	//r.MapWorkerIterators[index.String()] = append(r.MapWorkerIterators[index.String()], iterator)
	r.MapIterators[index.String()] = iterator

	iteratorMessage = &pb.IteratorMessage{Id: index.String()}

	return
}

// runIterator : Iterator run function.
func (r ConnectorGrpc) runIterator(value, msgtype string, iterator *msg.Iterator, channel chan msg.Message) {
	log.Printf("Run iterator %s on value %s", msgtype, value)

	for {
		messageIterator := iterator.Get()

		if messageIterator != nil {
			if msgtype == "cmd" {
				message := (messageIterator.GetMessage()).(msg.Command)

				if value == message.GetCommand() {
					log.Println("Get iterator command")
					log.Println(message)

					channel <- message

					break
				}
			} else if msgtype == "evt" {
				message := (messageIterator.GetMessage()).(msg.Event)

				if value == message.Event {
					log.Println("Get iterator event")
					log.Println(message)

					channel <- message

					break
				}
			} else if msgtype == "topic" {
				message := (messageIterator.GetMessage()).(msg.Event)

				if value == message.Topic {
					log.Println("Get iterator event by topic")
					log.Println(message)

					channel <- message

					break
				}
			} else if msgtype == "validation" {
				message := (messageIterator.GetMessage()).(msg.Event)

				if value == message.ReferencesUUID && message.Event == "TAKEN" {
					log.Println("Get iterator event validation")
					log.Println(message)

					channel <- message

					break
				}
			}
		}
		time.Sleep(time.Duration(2000) * time.Millisecond)
	}
	//delete(r.MapIterators, iteratorId)
}

// getSendIndex : Connector getSendIndex function.
func getSendIndex(conns []*sn.ShosetConn) int {
	aux := sendIndex
	sendIndex++

	if sendIndex >= len(conns) {
		sendIndex = 0
	}

	return aux
}
