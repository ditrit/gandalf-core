#! /usr/bin/env python3
# coding: utf-8

import unittest
import uuid
import grpc
from concurrent import futures

from pyclient.command.ClientCommand import ClientCommand

from pyclient.grpc.connectorCommand_pb2_grpc import *
from pyclient.grpc.connectorCommand_pb2 import *

from threading import Thread
from time import sleep

DEFAULT_PORT = "50150"
SERVER_TEST_TIME = 1.0


class TestConnectorCommandServicer(ConnectorCommandServicer):
    """gRPC Server for unit tests"""

    def __init__(self):
        super().__init__()

    def SendCommandMessage(self, request, context) -> CommandMessageUUID:
        print('Handle send command')
        # request == CommandMessage

        cmu = CommandMessageUUID(UUID=request.UUID)

        return cmu


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_ConnectorCommandServicer_to_server(
        TestConnectorCommandServicer(), server)
    server.add_insecure_port('[::]:'+DEFAULT_PORT)
    server.start()
    server.wait_for_termination(timeout=SERVER_TEST_TIME)


class TestClientCommand(unittest.TestCase):

    def setUp(self):
        self.thread = Thread(target=serve)
        self.thread.start()

    def test_send_command(self):
        """
        Test Send Command
        """

        cc = ClientCommand('localhost:'+DEFAULT_PORT, 'test')
        print(cc.SendCommand("TypeBidon", "CommandeBidon", "10000", "PayloadBidon"))


if __name__ == '__main__':
    unittest.main()
