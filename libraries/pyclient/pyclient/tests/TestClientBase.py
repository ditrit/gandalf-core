#! /usr/bin/env python3
# coding: utf-8

import unittest
import uuid
import grpc
from concurrent import futures

from pyclient.base.ClientBase import ClientBase

from pyclient.grpc.connector_pb2 import *
from pyclient.grpc.connector_pb2_grpc import *

from threading import Thread
from time import sleep

TEST_IDENTITY = "TestClientBase"
DEFAULT_PORT = "50150"
SERVER_STOP_TIME = 1.0

FIXED_TEST_COMMAND_LIST = ["DummyCommandA", "DummyCommandB", "DummyCommandC"]
FIXED_TEST_MAJOR = 42

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


class TestClientBase(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        add_ConnectorServicer_to_server(
            TestConnectorBaseServicer(), server)
        server.add_insecure_port('[::]:'+DEFAULT_PORT)
        thread = Thread(target=serve)
        thread.start()
        cls.client = ClientBase(TEST_IDENTITY, 'localhost:'+DEFAULT_PORT)

    @classmethod
    def tearDownClass(cls):
        server.stop(SERVER_STOP_TIME)

    def test_send_command_list(self):
        """
        Test Send Command List
        """
        print('TEST : ClientBase.SendCommandList(major, commandList)')

        result = self.client.SendCommandList(FIXED_TEST_MAJOR, FIXED_TEST_COMMAND_LIST)

if __name__ == '__main__':
    unittest.main()
