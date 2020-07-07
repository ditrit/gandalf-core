#! /usr/bin/env python3
# coding: utf-8

import unittest
import grpc
from concurrent import futures

from pyclient.command.ClientCommand import ClientCommand

from pyclient.grpc.connectorCommand_pb2_grpc import *
from pyclient.grpc.connectorCommand_pb2 import *

class TestConnectorCommandServicer(ConnectorCommandServicer):
    """gRPC Server for unit tests"""

    def __init__(self):
        super().__init__()
    
    def SendCommandMessage(self, request, context) -> CommandMessageUUID:
        print('Handle send command')
        print(request)
        print(context)

        return CommandMessageUUID(name="yolo")
    


class TestClientCommand(unittest.TestCase):        

    def test_send_command(self):
        """
        Test Send Command
        """

        cc = ClientCommand('localhost:50051', 'test')
        cc.SendCommand("TypeBidon", "CommandeBidon", "10000", "PayloadBidon")

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_ConnectorCommandServicer_to_server(TestConnectorCommandServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()      

if __name__ == '__main__':
    serve()
    unittest.main()