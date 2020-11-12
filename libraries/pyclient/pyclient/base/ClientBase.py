#! /usr/bin/env python3
# coding: utf-8

from typing import List

from ..grpc.connector_pb2 import *
from ..grpc.connector_pb2_grpc import *
from ..ClientWarper import ClientWarper


class ClientBase(ClientWarper):

    @property
    def ClientBaseConnection(self) -> str:
        return self.ClientConnection

    @ClientBaseConnection.setter
    def ClientBaseConnection(self, value: str):
        self.ClientConnection = value

    def SendCommandList(self, major: int, minor: int, commands: List[str]) -> Empty:
        commandList = CommandList()
        commandList.Major = major
        commandList.Minor = minor
        commandList.Commands.extend(commands)

        print("SEND COMMAND LIST LIB")
        return self.Client.SendCommandList(commandList)

    def SendStop(self, major: int, minor: int):
        stop = Stop(major, minor)

        return self.Client.SendStop(stop)
