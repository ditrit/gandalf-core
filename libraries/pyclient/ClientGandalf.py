#! /usr/bin/env python3
# coding: utf-8

class ClientGandalf(self):
    
    @property
    def identity(self):
        return self._identity
    @identity.setter
    def identity(self, value):
        self._identity = value
    @identity.deleter
    def identity(self):
        del self._identity
    
    def __init__(self, identity: str, clientCommandConnection: str, clientEventConnection: str):
        super().__init__()

        self.identity = identity
        self.clientCommandConnection = clientCommandConnection
        self.clientEvenConenction = clientEventConnection

        # self.clientCommand = new ClientCommand(self.getClientCommandConnection(), this.getIdentity())
        # self.clientEvent = new ClientEvent(self.getClientEventConnection(), self.getIdentity())
    