#! /usr/bin/env python3
# coding: utf-8

from grpc import Channel
from ..ClientWarper import ClientWarper
import uuid
import time

from ..grpc.connectorCommand_pb2 import *
from ..grpc.connectorCommand_pb2_grpc import *

from ..grpc.connector_pb2 import IteratorMessage
from ..grpc.connector_pb2 import Empty


class ClientCommand(ClientWarper):    

    Client: ConnectorCommandStub

    @property
    def ClientCommandConnection(self) -> str:
        return self.ClientConnection

    @ClientCommandConnection.setter
    def ClientCommandConnection(self, value: str):
        self.ClientConnection = value

    def __init__(self, identity: str, clientCommandConnection: str):
        super().__init__(identity, clientCommandConnection)
    
    def loadStub(self, conn: Channel):
        self.Client = ConnectorCommandStub(conn)

    def SendCommand(self, connectorType: str, command: str,  timeout: str, payload: str) -> CommandMessageUUID:
        commandMessage = CommandMessage()
        commandMessage.Timeout = timeout
        commandMessage.UUID = str(uuid.uuid4())
        commandMessage.ConnectorType = connectorType
        commandMessage.Command = command
        commandMessage.Payload = payload

        return self.Client.SendCommandMessage(commandMessage)

    def WaitCommand(self, command: str, idIterator: str, major: int) -> CommandMessage:
        commandMessageWait = CommandMessageWait()
        commandMessageWait.WorkerSource = self.Identity
        commandMessageWait.Value = command
        commandMessageWait.IteratorId = idIterator
        commandMessageWait.Major = major

        commandMessage = self.Client.WaitCommandMessage(commandMessageWait)
        print(commandMessage)

        while commandMessage == None:
            time.sleep(1)

        return commandMessage

    def CreateIteratorCommand(self) -> IteratorMessage:
        return self.Client.CreateIteratorCommand(Empty())
