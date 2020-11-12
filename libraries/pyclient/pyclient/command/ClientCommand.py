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

    client: ConnectorCommandStub

    @property
    def clientCommandConnection(self) -> str:
        return self.clientConnection

    @clientCommandConnection.setter
    def clientCommandConnection(self, value: str):
        self.clientConnection = value

    def __init__(self, identity: str, clientCommandConnection: str):
        super().__init__(identity, clientCommandConnection)
    
    def loadStub(self, conn: Channel):
        self.client = ConnectorCommandStub(conn)

    def SendCommand(self, connectorType: str, command: str,  timeout: str, payload: str) -> CommandMessageUUID:
        commandMessage = CommandMessage()
        commandMessage.Timeout = timeout
        commandMessage.UUID = str(uuid.uuid4())
        commandMessage.ConnectorType = connectorType
        commandMessage.Command = command
        commandMessage.Payload = payload

        return self.client.SendCommandMessage(commandMessage)

    def WaitCommand(self, command: str, idIterator: str, major: int) -> CommandMessage:
        commandMessageWait = CommandMessageWait()
        commandMessageWait.WorkerSource = self.identity
        commandMessageWait.Value = command
        commandMessageWait.IteratorId = idIterator
        commandMessageWait.Major = major

        commandMessage = self.client.WaitCommandMessage(commandMessageWait)
        print(commandMessage)

        while commandMessage == None:
            time.sleep(1)

        return commandMessage

    def CreateIteratorCommand(self) -> IteratorMessage:
        return self.client.CreateIteratorCommand(Empty())
