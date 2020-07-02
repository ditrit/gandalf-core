#! /usr/bin/env python3
# coding: utf-8

import grpc
from pyclient.grpc import *

class ClientCommand(self):
    @property
    def clientCommandConnection(self):
        return self._clientCommandConnection
    @clientCommandConnection.setter
    def clientCommandConnection(self, value):
        self._clientCommandConnection = value

    @property
    def identity(self):
        return self._identity
    @identity.setter
    def identity(self, value):
        self._identity = value

    @property
    def channel(self):
        return self._channel
    @channel.setter
    def channel(self, value):
        self._channel = value

    def __init__(self, clientCommandConnection: str, identity: str):
        super().__init__()
        self.clientCommandConnection = clientCommandConnection
        self.identity = identity
        connections = self.clientCommandConnection.split(":")
        self.channel = grpc.insecure_channel(connections[0]+":"+connections[1]) # grpc.insecure_channel(connections) ??? need to be checked

    def SendCommand(self, context: str, timeout: str, uuid: str, connectorType: str, commandType: str, command: str, payload: str) -> str:
        stub = connectorCommand_pb2_grpc.ConnectorCommand(self.channel)
        
        commandMessage = CommandMessage()
        commandMessage.Context = context
        commandMessage.Timeout = timeout
        commandMessage.UUID = uuid
        commandMessage.ConnectorType = connectorType
        commandMessage.CommandType = commandType
        commandMessage.Command = command
        commandMessage.Payload = payload

        commandMessageUUID, _ = stub.sendCommandMessage(commandMessage)

        return commandMessageUUID


