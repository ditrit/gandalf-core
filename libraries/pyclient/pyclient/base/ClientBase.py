#! /usr/bin/env python3
# coding: utf-8

from typing import List

from grpc import Channel

from ..grpc.connector_pb2 import *
from ..grpc.connector_pb2_grpc import *
from ..ClientWarper import ClientWarper


class ClientBase(ClientWarper):

    @property
    def clientBaseConnection(self) -> str:
        return self.clientConnection

    @clientBaseConnection.setter
    def clientBaseConnection(self, value: str):
        self.clientConnection = value

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
