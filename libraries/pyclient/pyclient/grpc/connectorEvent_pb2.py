# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: connectorEvent.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import connector_pb2 as connector__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='connectorEvent.proto',
  package='grpc',
  syntax='proto3',
  serialized_options=b'\n\034com.ditrit.gandalf.java.grpcB\023ConnectorEventProtoP\001',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x14\x63onnectorEvent.proto\x12\x04grpc\x1a\x0f\x63onnector.proto\"\xa5\x01\n\x0c\x45ventMessage\x12\x0e\n\x06Tenant\x18\x01 \x01(\t\x12\r\n\x05Token\x18\x02 \x01(\t\x12\r\n\x05Topic\x18\x03 \x01(\t\x12\x0f\n\x07Timeout\x18\x04 \x01(\t\x12\x11\n\tTimestamp\x18\x05 \x01(\t\x12\x0c\n\x04UUID\x18\x06 \x01(\t\x12\r\n\x05\x45vent\x18\x07 \x01(\t\x12\x0f\n\x07Payload\x18\x08 \x01(\t\x12\x15\n\rReferenceUUID\x18\t \x01(\t\"q\n\x10\x45ventMessageWait\x12\x14\n\x0cWorkerSource\x18\x01 \x01(\t\x12\r\n\x05\x45vent\x18\x02 \x01(\t\x12\r\n\x05Topic\x18\x03 \x01(\t\x12\x12\n\nIteratorId\x18\x04 \x01(\t\x12\x15\n\rReferenceUUID\x18\x05 \x01(\t\"b\n\x10TopicMessageWait\x12\x14\n\x0cWorkerSource\x18\x01 \x01(\t\x12\r\n\x05Topic\x18\x02 \x01(\t\x12\x12\n\nIteratorId\x18\x03 \x01(\t\x12\x15\n\rReferenceUUID\x18\x04 \x01(\t2\x88\x02\n\x0e\x43onnectorEvent\x12\x35\n\x10SendEventMessage\x12\x12.grpc.EventMessage\x1a\x0b.grpc.Empty\"\x00\x12@\n\x10WaitEventMessage\x12\x16.grpc.EventMessageWait\x1a\x12.grpc.EventMessage\"\x00\x12@\n\x10WaitTopicMessage\x12\x16.grpc.TopicMessageWait\x1a\x12.grpc.EventMessage\"\x00\x12;\n\x13\x43reateIteratorEvent\x12\x0b.grpc.Empty\x1a\x15.grpc.IteratorMessage\"\x00\x42\x35\n\x1c\x63om.ditrit.gandalf.java.grpcB\x13\x43onnectorEventProtoP\x01\x62\x06proto3'
  ,
  dependencies=[connector__pb2.DESCRIPTOR,])




_EVENTMESSAGE = _descriptor.Descriptor(
  name='EventMessage',
  full_name='grpc.EventMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='Tenant', full_name='grpc.EventMessage.Tenant', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Token', full_name='grpc.EventMessage.Token', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Topic', full_name='grpc.EventMessage.Topic', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Timeout', full_name='grpc.EventMessage.Timeout', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Timestamp', full_name='grpc.EventMessage.Timestamp', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='UUID', full_name='grpc.EventMessage.UUID', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Event', full_name='grpc.EventMessage.Event', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Payload', full_name='grpc.EventMessage.Payload', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ReferenceUUID', full_name='grpc.EventMessage.ReferenceUUID', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=48,
  serialized_end=213,
)


_EVENTMESSAGEWAIT = _descriptor.Descriptor(
  name='EventMessageWait',
  full_name='grpc.EventMessageWait',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='WorkerSource', full_name='grpc.EventMessageWait.WorkerSource', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Event', full_name='grpc.EventMessageWait.Event', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Topic', full_name='grpc.EventMessageWait.Topic', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='IteratorId', full_name='grpc.EventMessageWait.IteratorId', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ReferenceUUID', full_name='grpc.EventMessageWait.ReferenceUUID', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=215,
  serialized_end=328,
)


_TOPICMESSAGEWAIT = _descriptor.Descriptor(
  name='TopicMessageWait',
  full_name='grpc.TopicMessageWait',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='WorkerSource', full_name='grpc.TopicMessageWait.WorkerSource', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Topic', full_name='grpc.TopicMessageWait.Topic', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='IteratorId', full_name='grpc.TopicMessageWait.IteratorId', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ReferenceUUID', full_name='grpc.TopicMessageWait.ReferenceUUID', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=330,
  serialized_end=428,
)

DESCRIPTOR.message_types_by_name['EventMessage'] = _EVENTMESSAGE
DESCRIPTOR.message_types_by_name['EventMessageWait'] = _EVENTMESSAGEWAIT
DESCRIPTOR.message_types_by_name['TopicMessageWait'] = _TOPICMESSAGEWAIT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

EventMessage = _reflection.GeneratedProtocolMessageType('EventMessage', (_message.Message,), {
  'DESCRIPTOR' : _EVENTMESSAGE,
  '__module__' : 'connectorEvent_pb2'
  # @@protoc_insertion_point(class_scope:grpc.EventMessage)
  })
_sym_db.RegisterMessage(EventMessage)

EventMessageWait = _reflection.GeneratedProtocolMessageType('EventMessageWait', (_message.Message,), {
  'DESCRIPTOR' : _EVENTMESSAGEWAIT,
  '__module__' : 'connectorEvent_pb2'
  # @@protoc_insertion_point(class_scope:grpc.EventMessageWait)
  })
_sym_db.RegisterMessage(EventMessageWait)

TopicMessageWait = _reflection.GeneratedProtocolMessageType('TopicMessageWait', (_message.Message,), {
  'DESCRIPTOR' : _TOPICMESSAGEWAIT,
  '__module__' : 'connectorEvent_pb2'
  # @@protoc_insertion_point(class_scope:grpc.TopicMessageWait)
  })
_sym_db.RegisterMessage(TopicMessageWait)


DESCRIPTOR._options = None

_CONNECTOREVENT = _descriptor.ServiceDescriptor(
  name='ConnectorEvent',
  full_name='grpc.ConnectorEvent',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=431,
  serialized_end=695,
  methods=[
  _descriptor.MethodDescriptor(
    name='SendEventMessage',
    full_name='grpc.ConnectorEvent.SendEventMessage',
    index=0,
    containing_service=None,
    input_type=_EVENTMESSAGE,
    output_type=connector__pb2._EMPTY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='WaitEventMessage',
    full_name='grpc.ConnectorEvent.WaitEventMessage',
    index=1,
    containing_service=None,
    input_type=_EVENTMESSAGEWAIT,
    output_type=_EVENTMESSAGE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='WaitTopicMessage',
    full_name='grpc.ConnectorEvent.WaitTopicMessage',
    index=2,
    containing_service=None,
    input_type=_TOPICMESSAGEWAIT,
    output_type=_EVENTMESSAGE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='CreateIteratorEvent',
    full_name='grpc.ConnectorEvent.CreateIteratorEvent',
    index=3,
    containing_service=None,
    input_type=connector__pb2._EMPTY,
    output_type=connector__pb2._ITERATORMESSAGE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_CONNECTOREVENT)

DESCRIPTOR.services_by_name['ConnectorEvent'] = _CONNECTOREVENT

# @@protoc_insertion_point(module_scope)