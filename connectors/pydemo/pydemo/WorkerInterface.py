#! /usr/bin/env python3
# coding: utf-8

from abc import ABCMeta, abstractmethod

from pyclient.ClientGandalf import ClientGandalf
from pyclient.grpc.connectorCommand_pb2 import CommandMessage
from pyworker.Worker import Worker


class WorkerInterface(Worker, metaclass=ABCMeta):

    @abstractmethod
    def runTest1(self, clientGandalf: ClientGandalf, major: int, command: CommandMessage) -> int:
        raise NotImplementedError

    def __init__(self, major: int, minor: int):
        super().__init__(major, minor)

    def Run(self):
        super().Run()
