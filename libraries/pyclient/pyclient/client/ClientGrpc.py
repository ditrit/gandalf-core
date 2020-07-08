#! /usr/bin/env python3
# coding: utf-8

from typing import List

from ..command.ClientCommand import ClientCommand
from ..event.ClientEvent import ClientEvent
from ..base.ClientBase import ClientBase


class ClientGrpc:

    @property
    def Identity(self):
        return self._Identity

    @Identity.setter
    def Identity(self, value):
        self._Identity = value

    @property
    def ClientConnection(self):
        return self._ClientConnection

    @ClientConnection.setter
    def ClientConnection(self, value):
        self._ClientConnection = value

    @property
    def ClientBase(self):
        return self._ClientBase

    @ClientBase.setter
    def ClientBase(self, value):
        self._ClientBase = value

    @property
    def ClientCommand(self):
        return self._ClientCommand

    @ClientCommand.setter
    def ClientCommand(self, value):
        self._ClientCommand = value

    @property
    def ClientEvent(self):
        return self._ClientEvent

    @ClientEvent.setter
    def ClientEvent(self, value):
        self._ClientEvent = value

    def __init__(self, identity: str, clientConnection: str):
        self.Identity = identity
        self.ClientConnection = clientConnection

        self.ClientBase = ClientBase(identity, clientConnection)
        self.ClientCommand = ClientCommand(identity, clientConnection)
        self.ClientEvent = ClientEvent(identity, clientConnection)

    def SendCommandList(self, version: int, commands: List[str]) -> Empty:
        return self.ClientBase.SendCommandList(version, commands)

    def SendCommand(self, connectorType, command, timeout, payload: str) -> CommandMessageUUID:
        return self.ClientCommand.SendCommand(connectorType, command, timeout, payload)

    def SendEvent(self, topic, event, referenceUUID, timeout, payload: str) -> Empty:
        return self.ClientEvent.SendEvent(topic, event, referenceUUID, timeout, payload)

    def WaitCommand(self, command, idIterator: str, version: int) -> CommandMessage:
        return self.ClientCommand.WaitCommand(command, idIterator, version)

    def WaitEvent(self, topic, event, referenceUUID, idIterator: str) -> EventMessage:
        return self.ClientEvent.WaitEvent(topic, event, referenceUUID, idIterator)

    def WaitTopic(self, topic, referenceUUID, idIterator: str) -> EventMessage:
        return self.ClientEvent.WaitTopic(topic, referenceUUID, idIterator)

    def CreateIteratorCommand(self) -> IteratorMessage:
        return self.ClientCommand.CreateIteratorCommand()

    def CreateIteratorEvent(self) -> IteratorMessage:
        return self.ClientEvent.CreateIteratorEvent()
