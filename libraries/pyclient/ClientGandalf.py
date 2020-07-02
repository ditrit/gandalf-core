#! /usr/bin/env python3
# coding: utf-8

from pyclient.command.ClientCommand import ClientCommand
from pyclient.event.ClientEvent import ClientEvent

class ClientGandalf(self):
    
    @property
    def identity(self):
        return self._identity
    @identity.setter
    def identity(self, value):
        self._identity = value

    @property
    def clientCommandConnection(self):
        return self._clientCommandConnection
    @clientCommandConnection.setter
    def clientCommandConnection(self, value):
        self._clientCommandConnection = value
    
    @property
    def clientEventConnection(self):
        return self._clientEventConnection
    @clientEventConnection.setter
    def clientEventConnection(self, value):
        self._clientEventConnection = value

    @property
    def clientCommand(self):
        return self._clientCommand
    @clientCommand.setter
    def clientCommand(self, value):
        self._clientCommand = value
    
    @property
    def clientEvent(self):
        return self._clientEvent
    @clientEvent.setter
    def clientEvent(self, value):
        self._clientEvent = value

    def __init__(self, identity: str, clientCommandConnection: str, clientEventConnection: str):
        super().__init__()

        self.identity = identity
        self.clientCommandConnection = clientCommandConnection
        self.clientEvenConenction = clientEventConnection

        self.clientCommand = ClientCommand(self.clientCommandConnection, self.identity)
        self.clientEvent = ClientEvent(self.clientEventConnection, self.identity)
    