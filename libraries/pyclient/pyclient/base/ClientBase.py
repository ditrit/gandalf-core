#! /usr/bin/env python3
# coding: utf-8

from typing import List

from ..grpc.connector_pb2 import *
from ..grpc.connector_pb2_grpc import *


class ClientBase:

    @property
    def ClientBaseConnection(self):
        return self._ClientBaseConnection

    @ClientBaseConnection.setter
    def ClientBaseConnection(self, value):
        self._ClientBaseConnection = value

    @property
    def Identity(self):
        return self._Identity

    @Identity.setter
    def Identity(self, value):
        self._Identity = value

    @property
    def client(self):
        return self._client

    @client.setter
    def client(self, value):
        self._client = value

    def __init__(self, identity: str, clientBaseConnection: str):
        self.Identity = identity
        self.ClientBaseConnection = clientBaseConnection

        conn = grpc.insecure_channel(clientBaseConnection)
        self.client = ConnectorStub(conn)

    def SendCommandList(self, major: int, commands: List[str]) -> Empty:
        commandList = CommandList()
        commandList.Major = major
        commandList.Commands.extend(commands)

        return self.client.SendCommandList(commandList)
