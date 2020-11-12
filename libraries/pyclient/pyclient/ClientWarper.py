#! /usr/bin/env python3
# coding: utf-8

from typing import List

from grpc import Channel, RpcError

from .grpc.connector_pb2 import *
from .grpc.connector_pb2_grpc import *


class ClientWarper:

    ClientConnection: str
    Identity: str
    Client: ConnectorStub

    def __init__(self, identity: str, clientConnection: str):
        self.Identity = identity
        self.ClientConnection = clientConnection

        try:
            conn = grpc.insecure_channel(clientConnection)

            self.loadStub(conn)
        except RpcError as err:
            print("{} failed dial : {}".format(type(self).__name__, err))

    def loadStub(self, conn: Channel):
        self.Client = ConnectorStub(conn)
        print("{} connect : {}".format(
            type(self).__name__, self.ClientConnection))
