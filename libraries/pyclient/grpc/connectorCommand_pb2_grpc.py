# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import connectorCommand_pb2 as connectorCommand__pb2
import connector_pb2 as connector__pb2


class ConnectorCommandStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SendCommandMessage = channel.unary_unary(
                '/grpc.ConnectorCommand/SendCommandMessage',
                request_serializer=connectorCommand__pb2.CommandMessage.SerializeToString,
                response_deserializer=connectorCommand__pb2.CommandMessageUUID.FromString,
                )
        self.WaitCommandMessage = channel.unary_unary(
                '/grpc.ConnectorCommand/WaitCommandMessage',
                request_serializer=connectorCommand__pb2.CommandMessageWait.SerializeToString,
                response_deserializer=connectorCommand__pb2.CommandMessage.FromString,
                )
        self.CreateIteratorCommand = channel.unary_unary(
                '/grpc.ConnectorCommand/CreateIteratorCommand',
                request_serializer=connector__pb2.Empty.SerializeToString,
                response_deserializer=connector__pb2.IteratorMessage.FromString,
                )


class ConnectorCommandServicer(object):
    """Missing associated documentation comment in .proto file."""

    def SendCommandMessage(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def WaitCommandMessage(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateIteratorCommand(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ConnectorCommandServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'SendCommandMessage': grpc.unary_unary_rpc_method_handler(
                    servicer.SendCommandMessage,
                    request_deserializer=connectorCommand__pb2.CommandMessage.FromString,
                    response_serializer=connectorCommand__pb2.CommandMessageUUID.SerializeToString,
            ),
            'WaitCommandMessage': grpc.unary_unary_rpc_method_handler(
                    servicer.WaitCommandMessage,
                    request_deserializer=connectorCommand__pb2.CommandMessageWait.FromString,
                    response_serializer=connectorCommand__pb2.CommandMessage.SerializeToString,
            ),
            'CreateIteratorCommand': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateIteratorCommand,
                    request_deserializer=connector__pb2.Empty.FromString,
                    response_serializer=connector__pb2.IteratorMessage.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'grpc.ConnectorCommand', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ConnectorCommand(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def SendCommandMessage(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/grpc.ConnectorCommand/SendCommandMessage',
            connectorCommand__pb2.CommandMessage.SerializeToString,
            connectorCommand__pb2.CommandMessageUUID.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def WaitCommandMessage(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/grpc.ConnectorCommand/WaitCommandMessage',
            connectorCommand__pb2.CommandMessageWait.SerializeToString,
            connectorCommand__pb2.CommandMessage.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateIteratorCommand(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/grpc.ConnectorCommand/CreateIteratorCommand',
            connector__pb2.Empty.SerializeToString,
            connector__pb2.IteratorMessage.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)
