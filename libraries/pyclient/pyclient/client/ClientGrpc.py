#! /usr/bin/env python3
# coding: utf-8

from typing import List

from grpc import Channel

from ..command.ClientCommand import ClientCommand
from ..event.ClientEvent import ClientEvent
from ..base.ClientBase import ClientBase
from ..grpc.connector_pb2 import *
from ..grpc.connectorCommand_pb2 import *
from ..grpc.connectorEvent_pb2 import *
from ..ClientWarper import ClientWarper

class ClientGrpc(ClientWarper):
    
    clientBase: ClientBase
    clientCommand: ClientCommand
    clientEvent: ClientEvent

    def __init__(self, identity: str, clientConnection: str):
        super().__init__(identity, clientConnection)

        self.clientBase = ClientBase(identity, clientConnection)
        self.clientCommand = ClientCommand(identity, clientConnection)
        self.clientEvent = ClientEvent(identity, clientConnection)

    def loadStub(self, conn: Channel):
        # loadStub Override so that we don't call the ClientWarper one
        pass

    def SendCommandList(self, major: int, minor: int, commands: List[str]) -> Validate:
        return self.clientBase.SendCommandList(major, minor, commands)

    def SendStop(self, major: int, minor:int) -> Validate:
        return self.clientBase.SendStop(major, minor)

    def SendCommand(self, connectorType, command, timeout, payload: str) -> CommandMessageUUID:
        return self.clientCommand.SendCommand(connectorType, command, timeout, payload)

    def SendEvent(self, topic, event, referenceUUID, timeout, payload: str) -> Empty:
        return self.clientEvent.SendEvent(topic, event, referenceUUID, timeout, payload)

    def WaitCommand(self, command, idIterator: str, version: int) -> CommandMessage:
        return self.clientCommand.WaitCommand(command, idIterator, version)

    def WaitEvent(self, topic, event, referenceUUID, idIterator: str) -> EventMessage:
        return self.clientEvent.WaitEvent(topic, event, referenceUUID, idIterator)

    def WaitTopic(self, topic, referenceUUID, idIterator: str) -> EventMessage:
        return self.clientEvent.WaitTopic(topic, referenceUUID, idIterator)

    def CreateIteratorCommand(self) -> IteratorMessage:
        return self.clientCommand.CreateIteratorCommand()

    def CreateIteratorEvent(self) -> IteratorMessage:
        return self.clientEvent.CreateIteratorEvent()
