#! /usr/bin/env python3
# coding: utf-8

import unittest
import grpc
from concurrent import futures
from threading import Thread

from pyclient.ClientGandalf import ClientGandalf
from pyclient.models.Options import Options
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
FIXED_TEST_CONNECTIONS = ["localhost:50150"]
FIXED_TEST_UUID = "3dc28f74-9ad3-4fc7-8b2b-03b7aab660c9"
FIXED_TEST_ITERATOR_ID = "DummyIteratorId"
FIXED_TEST_PAYLOAD = "DummyPayload"
FIXED_TEST_EVENT = "DummyEvent"
FIXED_TEST_TOPIC = "DummyTopic"
FIXED_TEST_COMMAND = "DummyCommand"
FIXED_TEST_COMMAND_TYPE = "DummyCommandType"
FIXED_TEST_MAJOR = 3
FIXED_TEST_MINOR = 42
FIXED_TEST_COMMAND_LIST = ["DummyCommandA", "DummyCommandB", "DummyCommandC"]


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

        return EventMessage(Topic=request.Topic, ReferenceUUID=request.ReferenceUUID, Event="SUCCES")

    def WaitTopicMessage(self, request, context) -> EventMessage:
        print('Handle Wait Topic Message')

        return EventMessage(Topic=request.Topic, ReferenceUUID=request.ReferenceUUID, Event="SUCCES")

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

    def test_send_command(self):
        self.client.SendCommand(FIXED_TEST_COMMAND_TYPE+"."+FIXED_TEST_COMMAND, Options(
            FIXED_TEST_TIMEOUT, FIXED_TEST_PAYLOAD))

    def test_send_event(self):
        self.client.SendEvent(FIXED_TEST_TOPIC, FIXED_TEST_EVENT, Options(
            FIXED_TEST_TIMEOUT, FIXED_TEST_PAYLOAD))

    def test_send_reply(self):
        self.client.SendReply(FIXED_TEST_TOPIC, FIXED_TEST_EVENT, FIXED_TEST_UUID, Options(
            FIXED_TEST_TIMEOUT, FIXED_TEST_PAYLOAD))

    def test_send_command_list(self):
        self.client.SendCommandList(FIXED_TEST_MAJOR, FIXED_TEST_MINOR, FIXED_TEST_COMMAND_LIST)

    def test_wait_command(self):
        self.client.WaitCommand(
            FIXED_TEST_COMMAND, FIXED_TEST_ITERATOR_ID, FIXED_TEST_MAJOR)

    def test_wait_event(self):
        self.client.WaitEvent(
            FIXED_TEST_TOPIC, FIXED_TEST_EVENT, FIXED_TEST_ITERATOR_ID)

    def test_wait_reply_by_event(self):
        self.client.WaitReplyByEvent(
            FIXED_TEST_TOPIC, FIXED_TEST_EVENT, FIXED_TEST_UUID, FIXED_TEST_ITERATOR_ID)

    def test_wait_topic(self):
        self.client.WaitTopic(FIXED_TEST_TOPIC, FIXED_TEST_ITERATOR_ID)
    
    def test_wait_reply_by_topic(self):
        self.client.WaitReplyByTopic(FIXED_TEST_TOPIC, FIXED_TEST_UUID, FIXED_TEST_ITERATOR_ID)
    
    def test_wait_all_reply_by_topic(self):
        self.client.WaitAllReplyByTopic(FIXED_TEST_TOPIC, FIXED_TEST_UUID, FIXED_TEST_ITERATOR_ID, FIXED_TEST_MAJOR)

    def test_create_iterator_command(self):
        self.client.CreateIteratorCommand()
    
    def test_create_iterator_event(self):
        self.client.CreateIteratorEvent()

    def test_get_client_index(self):
        self.client.getClientIndex(FIXED_TEST_COMMAND_LIST, False)

if __name__ == '__main__':
    unittest.main()
