// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connectorEvent.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EventMessage struct {
	Tenant               string   `protobuf:"bytes,1,opt,name=Tenant,proto3" json:"Tenant,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	Topic                string   `protobuf:"bytes,3,opt,name=Topic,proto3" json:"Topic,omitempty"`
	Timeout              string   `protobuf:"bytes,4,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
	Timestamp            string   `protobuf:"bytes,5,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	UUID                 string   `protobuf:"bytes,6,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Event                string   `protobuf:"bytes,7,opt,name=Event,proto3" json:"Event,omitempty"`
	Payload              string   `protobuf:"bytes,8,opt,name=Payload,proto3" json:"Payload,omitempty"`
	ReferenceUUID        string   `protobuf:"bytes,9,opt,name=ReferenceUUID,proto3" json:"ReferenceUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventMessage) Reset()         { *m = EventMessage{} }
func (m *EventMessage) String() string { return proto.CompactTextString(m) }
func (*EventMessage) ProtoMessage()    {}
func (*EventMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a15ac99c650ec24e, []int{0}
}

func (m *EventMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventMessage.Unmarshal(m, b)
}
func (m *EventMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventMessage.Marshal(b, m, deterministic)
}
func (m *EventMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMessage.Merge(m, src)
}
func (m *EventMessage) XXX_Size() int {
	return xxx_messageInfo_EventMessage.Size(m)
}
func (m *EventMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EventMessage proto.InternalMessageInfo

func (m *EventMessage) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *EventMessage) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *EventMessage) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *EventMessage) GetTimeout() string {
	if m != nil {
		return m.Timeout
	}
	return ""
}

func (m *EventMessage) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *EventMessage) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *EventMessage) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *EventMessage) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *EventMessage) GetReferenceUUID() string {
	if m != nil {
		return m.ReferenceUUID
	}
	return ""
}

type EventMessageWait struct {
	WorkerSource         string   `protobuf:"bytes,1,opt,name=WorkerSource,proto3" json:"WorkerSource,omitempty"`
	Event                string   `protobuf:"bytes,2,opt,name=Event,proto3" json:"Event,omitempty"`
	Topic                string   `protobuf:"bytes,3,opt,name=Topic,proto3" json:"Topic,omitempty"`
	IteratorId           string   `protobuf:"bytes,4,opt,name=IteratorId,proto3" json:"IteratorId,omitempty"`
	ReferenceUUID        string   `protobuf:"bytes,5,opt,name=ReferenceUUID,proto3" json:"ReferenceUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventMessageWait) Reset()         { *m = EventMessageWait{} }
func (m *EventMessageWait) String() string { return proto.CompactTextString(m) }
func (*EventMessageWait) ProtoMessage()    {}
func (*EventMessageWait) Descriptor() ([]byte, []int) {
	return fileDescriptor_a15ac99c650ec24e, []int{1}
}

func (m *EventMessageWait) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventMessageWait.Unmarshal(m, b)
}
func (m *EventMessageWait) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventMessageWait.Marshal(b, m, deterministic)
}
func (m *EventMessageWait) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMessageWait.Merge(m, src)
}
func (m *EventMessageWait) XXX_Size() int {
	return xxx_messageInfo_EventMessageWait.Size(m)
}
func (m *EventMessageWait) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMessageWait.DiscardUnknown(m)
}

var xxx_messageInfo_EventMessageWait proto.InternalMessageInfo

func (m *EventMessageWait) GetWorkerSource() string {
	if m != nil {
		return m.WorkerSource
	}
	return ""
}

func (m *EventMessageWait) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *EventMessageWait) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *EventMessageWait) GetIteratorId() string {
	if m != nil {
		return m.IteratorId
	}
	return ""
}

func (m *EventMessageWait) GetReferenceUUID() string {
	if m != nil {
		return m.ReferenceUUID
	}
	return ""
}

type TopicMessageWait struct {
	WorkerSource         string   `protobuf:"bytes,1,opt,name=WorkerSource,proto3" json:"WorkerSource,omitempty"`
	Topic                string   `protobuf:"bytes,2,opt,name=Topic,proto3" json:"Topic,omitempty"`
	IteratorId           string   `protobuf:"bytes,3,opt,name=IteratorId,proto3" json:"IteratorId,omitempty"`
	ReferenceUUID        string   `protobuf:"bytes,4,opt,name=ReferenceUUID,proto3" json:"ReferenceUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicMessageWait) Reset()         { *m = TopicMessageWait{} }
func (m *TopicMessageWait) String() string { return proto.CompactTextString(m) }
func (*TopicMessageWait) ProtoMessage()    {}
func (*TopicMessageWait) Descriptor() ([]byte, []int) {
	return fileDescriptor_a15ac99c650ec24e, []int{2}
}

func (m *TopicMessageWait) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicMessageWait.Unmarshal(m, b)
}
func (m *TopicMessageWait) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicMessageWait.Marshal(b, m, deterministic)
}
func (m *TopicMessageWait) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicMessageWait.Merge(m, src)
}
func (m *TopicMessageWait) XXX_Size() int {
	return xxx_messageInfo_TopicMessageWait.Size(m)
}
func (m *TopicMessageWait) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicMessageWait.DiscardUnknown(m)
}

var xxx_messageInfo_TopicMessageWait proto.InternalMessageInfo

func (m *TopicMessageWait) GetWorkerSource() string {
	if m != nil {
		return m.WorkerSource
	}
	return ""
}

func (m *TopicMessageWait) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *TopicMessageWait) GetIteratorId() string {
	if m != nil {
		return m.IteratorId
	}
	return ""
}

func (m *TopicMessageWait) GetReferenceUUID() string {
	if m != nil {
		return m.ReferenceUUID
	}
	return ""
}

func init() {
	proto.RegisterType((*EventMessage)(nil), "grpc.EventMessage")
	proto.RegisterType((*EventMessageWait)(nil), "grpc.EventMessageWait")
	proto.RegisterType((*TopicMessageWait)(nil), "grpc.TopicMessageWait")
}

func init() {
	proto.RegisterFile("connectorEvent.proto", fileDescriptor_a15ac99c650ec24e)
}

var fileDescriptor_a15ac99c650ec24e = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0xce, 0xd2, 0x40,
	0x14, 0xb5, 0xfd, 0x0a, 0xc8, 0x15, 0xb5, 0x19, 0xd0, 0x4c, 0x1a, 0x62, 0x4c, 0xe3, 0xc2, 0x55,
	0x17, 0x1a, 0x56, 0x6e, 0x0c, 0xe8, 0x82, 0x85, 0x09, 0xe1, 0x27, 0xac, 0xc7, 0xe9, 0x85, 0x54,
	0xe8, 0x4c, 0x33, 0x0c, 0x24, 0xbc, 0x81, 0x6b, 0x9f, 0xc2, 0x07, 0xf4, 0x01, 0x4c, 0x67, 0x5a,
	0x68, 0xa5, 0xc4, 0xe4, 0xdb, 0xcd, 0x39, 0x77, 0xe6, 0xdc, 0x73, 0x4f, 0xee, 0xc0, 0x80, 0x4b,
	0x21, 0x90, 0x6b, 0xa9, 0xbe, 0x9e, 0x50, 0xe8, 0x28, 0x53, 0x52, 0x4b, 0xe2, 0x6d, 0x55, 0xc6,
	0x83, 0x97, 0x97, 0x9a, 0xa5, 0xc3, 0x3f, 0x0e, 0xf4, 0xcc, 0xb5, 0x6f, 0x78, 0x38, 0xb0, 0x2d,
	0x92, 0xd7, 0xd0, 0x5e, 0xa2, 0x60, 0x42, 0x53, 0xe7, 0xad, 0xf3, 0xbe, 0x3b, 0x2f, 0x10, 0x19,
	0x40, 0x6b, 0x29, 0x77, 0x28, 0xa8, 0x6b, 0x68, 0x0b, 0x2c, 0x9b, 0x25, 0x9c, 0x3e, 0x94, 0x6c,
	0x96, 0x70, 0x42, 0xa1, 0xb3, 0x4c, 0x52, 0x94, 0x47, 0x4d, 0x3d, 0xc3, 0x97, 0x90, 0x0c, 0xa1,
	0x9b, 0x1f, 0x0f, 0x9a, 0xa5, 0x19, 0x6d, 0x99, 0xda, 0x95, 0x20, 0x04, 0xbc, 0xd5, 0x6a, 0xfa,
	0x85, 0xb6, 0x4d, 0xc1, 0x9c, 0xf3, 0x0e, 0xc6, 0x1f, 0xed, 0xd8, 0x0e, 0x06, 0xe4, 0x1d, 0x66,
	0xec, 0xbc, 0x97, 0x2c, 0xa6, 0x4f, 0x6d, 0x87, 0x02, 0x92, 0x77, 0xf0, 0x7c, 0x8e, 0x1b, 0x54,
	0x28, 0x38, 0x1a, 0xb1, 0xae, 0xa9, 0xd7, 0xc9, 0xf0, 0xb7, 0x03, 0x7e, 0x75, 0xec, 0x35, 0x4b,
	0x34, 0x09, 0xa1, 0xb7, 0x96, 0x6a, 0x87, 0x6a, 0x21, 0x8f, 0x8a, 0x63, 0x11, 0x40, 0x8d, 0xbb,
	0xda, 0x71, 0xab, 0x76, 0x9a, 0x63, 0x78, 0x03, 0x30, 0xd5, 0xa8, 0x98, 0x96, 0x6a, 0x1a, 0x17,
	0x49, 0x54, 0x98, 0x5b, 0xab, 0xad, 0x26, 0xab, 0xbf, 0x1c, 0xf0, 0x8d, 0xde, 0x23, 0xac, 0x5a,
	0x53, 0xee, 0x7d, 0x53, 0x0f, 0xff, 0x37, 0xe5, 0x35, 0x98, 0xfa, 0xf0, 0xd3, 0x85, 0x17, 0x93,
	0xda, 0x9a, 0x91, 0x11, 0xf8, 0x0b, 0x14, 0x71, 0x6d, 0x99, 0x48, 0x94, 0x6f, 0x5d, 0x54, 0xe5,
	0x82, 0x67, 0x05, 0x97, 0x66, 0xfa, 0x1c, 0x3e, 0x21, 0x9f, 0xc1, 0xcf, 0x27, 0xaa, 0xef, 0xe0,
	0xed, 0xb3, 0xfc, 0x4e, 0xd0, 0x20, 0x77, 0x55, 0xa8, 0x66, 0x54, 0x2a, 0xfc, 0x9b, 0xdb, 0x1d,
	0x85, 0x4f, 0xd0, 0x9f, 0x28, 0x64, 0x1a, 0xcb, 0x1c, 0xec, 0x44, 0x55, 0xa7, 0xc1, 0x2b, 0x0b,
	0xca, 0x1b, 0x97, 0xc7, 0xe3, 0x11, 0x0c, 0xb9, 0x4c, 0xa3, 0x38, 0xd1, 0x2a, 0xd1, 0xd1, 0x96,
	0x89, 0x98, 0xed, 0x37, 0xd1, 0x0f, 0x76, 0x62, 0xe6, 0xc5, 0xb8, 0x5f, 0xcf, 0x69, 0x96, 0x7f,
	0xbb, 0x99, 0xf3, 0xbd, 0x6d, 0xfe, 0xdf, 0xc7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x41, 0x1c,
	0x24, 0x94, 0xae, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConnectorEventClient is the client API for ConnectorEvent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConnectorEventClient interface {
	SendEventMessage(ctx context.Context, in *EventMessage, opts ...grpc.CallOption) (*Empty, error)
	WaitEventMessage(ctx context.Context, in *EventMessageWait, opts ...grpc.CallOption) (*EventMessage, error)
	WaitTopicMessage(ctx context.Context, in *TopicMessageWait, opts ...grpc.CallOption) (*EventMessage, error)
	CreateIteratorEvent(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*IteratorMessage, error)
}

type connectorEventClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectorEventClient(cc grpc.ClientConnInterface) ConnectorEventClient {
	return &connectorEventClient{cc}
}

func (c *connectorEventClient) SendEventMessage(ctx context.Context, in *EventMessage, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/grpc.ConnectorEvent/SendEventMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorEventClient) WaitEventMessage(ctx context.Context, in *EventMessageWait, opts ...grpc.CallOption) (*EventMessage, error) {
	out := new(EventMessage)
	err := c.cc.Invoke(ctx, "/grpc.ConnectorEvent/WaitEventMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorEventClient) WaitTopicMessage(ctx context.Context, in *TopicMessageWait, opts ...grpc.CallOption) (*EventMessage, error) {
	out := new(EventMessage)
	err := c.cc.Invoke(ctx, "/grpc.ConnectorEvent/WaitTopicMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorEventClient) CreateIteratorEvent(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*IteratorMessage, error) {
	out := new(IteratorMessage)
	err := c.cc.Invoke(ctx, "/grpc.ConnectorEvent/CreateIteratorEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectorEventServer is the server API for ConnectorEvent service.
type ConnectorEventServer interface {
	SendEventMessage(context.Context, *EventMessage) (*Empty, error)
	WaitEventMessage(context.Context, *EventMessageWait) (*EventMessage, error)
	WaitTopicMessage(context.Context, *TopicMessageWait) (*EventMessage, error)
	CreateIteratorEvent(context.Context, *Empty) (*IteratorMessage, error)
}

// UnimplementedConnectorEventServer can be embedded to have forward compatible implementations.
type UnimplementedConnectorEventServer struct {
}

func (*UnimplementedConnectorEventServer) SendEventMessage(ctx context.Context, req *EventMessage) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEventMessage not implemented")
}
func (*UnimplementedConnectorEventServer) WaitEventMessage(ctx context.Context, req *EventMessageWait) (*EventMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitEventMessage not implemented")
}
func (*UnimplementedConnectorEventServer) WaitTopicMessage(ctx context.Context, req *TopicMessageWait) (*EventMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitTopicMessage not implemented")
}
func (*UnimplementedConnectorEventServer) CreateIteratorEvent(ctx context.Context, req *Empty) (*IteratorMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIteratorEvent not implemented")
}

func RegisterConnectorEventServer(s *grpc.Server, srv ConnectorEventServer) {
	s.RegisterService(&_ConnectorEvent_serviceDesc, srv)
}

func _ConnectorEvent_SendEventMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorEventServer).SendEventMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ConnectorEvent/SendEventMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorEventServer).SendEventMessage(ctx, req.(*EventMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorEvent_WaitEventMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventMessageWait)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorEventServer).WaitEventMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ConnectorEvent/WaitEventMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorEventServer).WaitEventMessage(ctx, req.(*EventMessageWait))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorEvent_WaitTopicMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicMessageWait)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorEventServer).WaitTopicMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ConnectorEvent/WaitTopicMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorEventServer).WaitTopicMessage(ctx, req.(*TopicMessageWait))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorEvent_CreateIteratorEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorEventServer).CreateIteratorEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ConnectorEvent/CreateIteratorEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorEventServer).CreateIteratorEvent(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConnectorEvent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.ConnectorEvent",
	HandlerType: (*ConnectorEventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEventMessage",
			Handler:    _ConnectorEvent_SendEventMessage_Handler,
		},
		{
			MethodName: "WaitEventMessage",
			Handler:    _ConnectorEvent_WaitEventMessage_Handler,
		},
		{
			MethodName: "WaitTopicMessage",
			Handler:    _ConnectorEvent_WaitTopicMessage_Handler,
		},
		{
			MethodName: "CreateIteratorEvent",
			Handler:    _ConnectorEvent_CreateIteratorEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connectorEvent.proto",
}
