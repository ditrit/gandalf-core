#! /usr/bin/env python3
# coding: utf-8

import unittest
import grpc
from concurrent import futures
from threading import Thread

from pyclient.ClientGandalf import ClientGandalf
from pyclient.grpc.connector_pb2 import *
from pyclient.grpc.connector_pb2_grpc import *
from pyclient.grpc.connectorCommand_pb2_grpc import *
from pyclient.grpc.connectorCommand_pb2 import *
from pyclient.grpc.connectorEvent_pb2_grpc import *
from pyclient.grpc.connectorEvent_pb2 import *

DEFAULT_PORT = "50150"
SERVER_STOP_TIME = 1.0
FIXED_TEST_IDENTITY = "TestClientGandalfIdentity"
FIXED_TEST_TIMEOUT = "10000"
FIXED_TEST_CONNECTIONS = ["[::]:50150"]
FIXED_TEST_UUID = "3dc28f74-9ad3-4fc7-8b2b-03b7aab660c9"
FIXED_TEST_ITERATOR_ID = "DummyIteratorId"

server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))


def serve():
    server.start()
    server.wait_for_termination()


class TestConnectorBaseServicer(ConnectorServicer):
    """gRPC Server for unit tests"""

    def __init__(self):
        super().__init__()

    def SendCommandList(self, request, context) -> Empty:
        print('Handle Send Command List')

        return Empty()


class TestConnectorCommandServicer(ConnectorCommandServicer):
    """gRPC Server for unit tests"""

    def __init__(self):
        super().__init__()

    def SendCommandMessage(self, request, context) -> CommandMessageUUID:
        print('Handle send command')

        return CommandMessageUUID(UUID=FIXED_TEST_UUID)

    def WaitCommandMessage(self, request, context) -> CommandMessage:
        print('Handle Wait CommandMessage')

        return CommandMessage(SourceWorker=request.WorkerSource, Major=request.Major, Command=request.Value)

    def CreateIteratorCommand(self, request, context) -> IteratorMessage:
        print('Handle Create Iterator Command')

        return IteratorMessage(Id=FIXED_TEST_ITERATOR_ID)


class TestConnectorEventServicer(ConnectorEventServicer):
    """gRPC Server for unit tests"""

    def __init__(self):
        super().__init__()

    def SendEventMessage(self, request, context) -> Empty:
        print('Handle Send Event Message')

        return Empty()

    def WaitEventMessage(self, request, context) -> EventMessage:
        print('Handle Wait Event Message')

        return EventMessage(Topic=request.Topic, ReferenceUUID=request.ReferenceUUID)

    def WaitTopicMessage(self, request, context) -> EventMessage:
        print('Handle Wait Topic Message')

        return EventMessage(Topic=request.Topic, ReferenceUUID=request.ReferenceUUID)

    def CreateIteratorEvent(self, request, context) -> IteratorMessage:
        print('Handle Create Iterator Event')

        return IteratorMessage(Id=FIXED_TEST_ITERATOR_ID)


class TestClientGandalf(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        add_ConnectorServicer_to_server(TestConnectorBaseServicer(), server)
        add_ConnectorCommandServicer_to_server(
            TestConnectorCommandServicer(), server)
        add_ConnectorEventServicer_to_server(
            TestConnectorEventServicer(), server)

        server.add_insecure_port('[::]:'+DEFAULT_PORT)
        thread = Thread(target=serve)
        thread.start()

        cls.client = ClientGandalf(
            FIXED_TEST_IDENTITY, FIXED_TEST_TIMEOUT, FIXED_TEST_CONNECTIONS)

    @classmethod
    def tearDownClass(cls):
        server.stop(SERVER_STOP_TIME)


if __name__ == '__main__':
    unittest.main()
