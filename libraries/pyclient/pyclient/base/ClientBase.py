#! /usr/bin/env python3
# coding: utf-8

from typing import List

from ..grpc.connector_pb2 import *
from ..grpc.connector_pb2_grpc import *


class ClientBase:

    clientBaseConnection: str
    identity: str
    client: ConnectorStub

    def __init__(self, identity: str, clientBaseConnection: str):
        self.identity = identity
        self.clientBaseConnection = clientBaseConnection

        conn = grpc.insecure_channel(clientBaseConnection)

        self.client = ConnectorStub(conn)
        print("clientBase connect : {}".format(clientBaseConnection))

    def SendCommandList(self, major: int, minor: int, commands: List[str]) -> Empty:
        commandList = CommandList()
        commandList.Major = major
        commandList.Minor = minor
        commandList.Commands.extend(commands)

        print("SEND COMMAND LIST LIB")
        return self.client.SendCommandList(commandList)

    def SendStop(self, major: int, minor: int):
        stop = Stop(major, minor)

        return self.client.SendStop(stop)
