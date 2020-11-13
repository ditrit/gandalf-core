#! /usr/bin/env python3
# coding: utf-8

from typing import Dict

from pyclient.ClientGandalf import ClientGandalf
from pyclient.grpc.connectorCommand_pb2 import CommandMessage

from .WorkerInterface import WorkerInterface


class WorkerDemo(WorkerInterface):
    config: Dict

    def __init__(self, major: int, minor: int, config):
        super().__init__(major, minor)

        self.config = config

    def runTest1(self, clientGandalf: ClientGandalf, major: int, command: CommandMessage) -> int:
        print("EXECUTE RUN_TEST_1")
        print("COMMAND\n{}".format(command))
        print("PAYLOAD\n{}".format(command.Payload))

        return 0

    def runTest2(self, clientGandalf: ClientGandalf, major: int, command: CommandMessage) -> int:
        print("EXECUTE RUN_TEST_2")
        print("COMMAND\n{}".format(command))
        print("PAYLOAD\n{}".format(command.Payload))

        return 0
