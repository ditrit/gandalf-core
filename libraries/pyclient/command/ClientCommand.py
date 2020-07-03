#! /usr/bin/env python3
# coding: utf-8

import grpc
import uuid
import time

from ..grpc.connectorCommand_pb2 import *
from ..grpc.connectorCommand_pb2_grpc import *

class ClientCommand(self):
    @property
    def ClientCommandConnection(self):
        return self._ClientCommandConnection
    @ClientCommandConnection.setter
    def ClientCommandConnection(self, value):
        self._ClientCommandConnection = value

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

    def __init__(self, clientCommandConnection: str, identity: str):
        super().__init__()
        self.ClientCommandConnection = clientCommandConnection
        self.Identity = identity
        connections = self.ClientCommandConnection.split(":")
        conn = grpc.insecure_channel(connections[0]+":"+connections[1]) # grpc.insecure_channel(connections) ??? need to be checked
        self.client = ConnectorCommand(conn)

    def SendCommand(self, connectorType: str, command: str,  timeout: str, payload: str) -> CommandMessageUUID:
        commandMessage = CommandMessage()
        commandMessage.Timeout = timeout
        commandMessage.UUID = uuid.uuid4()
        commandMessage.ConnectorType = connectorType
        commandMessage.Command = command
        commandMessage.Payload = payload

        commandMessageUUID = self.client.SendCommandMessage(commandMessage)

        return commandMessageUUID

    def WaitCommand(self, command: str, idIterator: str, major: int) -> CommandMessage:
        commandMessageWait = CommandMessageWait()
        commandMessageWait.WorkerSource = self.Identity
        commandMessageWait.Value = command
        commandMessageWait.IteratorId = idIterator
        commandMessageWait.Major = major

        commandMessage = self.client.WaitCommandMessage(commandMessageWait)

        while commandMessage == None:
            time.sleep(1)

        return commandMessage

    def CreateIteratorCommand(self) -> IteratorMessage:
        iteratorMessage = self.client.CreateIteratorCommand()

        return iteratorMessage