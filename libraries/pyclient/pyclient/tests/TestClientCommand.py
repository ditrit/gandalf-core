#! /usr/bin/env python3
# coding: utf-8

import unittest
import uuid
import grpc
from concurrent import futures

from pyclient.command.ClientCommand import ClientCommand

from pyclient.grpc.connectorCommand_pb2_grpc import *
from pyclient.grpc.connectorCommand_pb2 import *
from pyclient.grpc.connector_pb2 import IteratorMessage

from threading import Thread
from time import sleep

TEST_IDENTITY = "TestClientCommand"
DEFAULT_PORT = "50150"
SERVER_STOP_TIME = 1.0

FIXED_TEST_UUID = "3dc28f74-9ad3-4fc7-8b2b-03b7aab660c9"
FIXED_TEST_COMMAND = "DummyCommand"
FIXED_TEST_COMMAND_TYPE = "DummyCommandType"
FIXED_TEST_ITERATOR_ID = "DummyIteratorId"
FIXED_TEST_PAYLOAD = "DummyPayload"
FIXED_TEST_TIMEOUT = "10000"
FIXED_TEST_MAJOR = 42

server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))


def serve():
    server.start()
    server.wait_for_termination()


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


class TestClientCommand(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        add_ConnectorCommandServicer_to_server(
            TestConnectorCommandServicer(), server)
        server.add_insecure_port('[::]:'+DEFAULT_PORT)
        thread = Thread(target=serve)
        thread.start()
        cls.client = ClientCommand(TEST_IDENTITY, 'localhost:'+DEFAULT_PORT)

    @classmethod
    def tearDownClass(cls):
        server.stop(SERVER_STOP_TIME)

    def test_send_command(self):
        """
        Test Send Command
        """
        print('TEST : ClientCommand.SendCommand(commandType, command, timeoutn, payload)')

        result = self.client.SendCommand(
            FIXED_TEST_COMMAND_TYPE, FIXED_TEST_COMMAND, FIXED_TEST_TIMEOUT, FIXED_TEST_PAYLOAD)

        self.assertEqual(result.UUID, FIXED_TEST_UUID)

    def test_wait_command(self):
        """
        Test Wait Command
        """
        print('TEST : ClientCommand.WaitCommand(command, iteratorId, major)')

        result = self.client.WaitCommand(
            FIXED_TEST_COMMAND, FIXED_TEST_ITERATOR_ID, FIXED_TEST_MAJOR)

        self.assertEqual(result.Command, FIXED_TEST_COMMAND)
        self.assertEqual(result.Major, FIXED_TEST_MAJOR)

    def test_create_iterator_command(self):
        """
        Test Create Iterator Command
        """
        print('TEST : ClientCommand.CreateIteratorCommand()')

        result = self.client.CreateIteratorCommand()

        self.assertEqual(result.Id, FIXED_TEST_ITERATOR_ID)


if __name__ == '__main__':
    unittest.main()
