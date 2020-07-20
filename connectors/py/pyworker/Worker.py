#! /usr/bin/env python3
# coding: utf-8

from typing import List, Callable

from pyclient.ClientGandalf import ClientGandalf

from .functions.Start import Start
from .functions.SendCommands import SendCommands


class Worker:

    version: str
    commandes: List[str]
    clientGandalf: ClientGandalf

    def Start(self, clientGandalf: ClientGandalf):
        pass

    def SendCommands(self, clientGandalf: ClientGandalf, version: int, commands: List[str]):
        pass

    def __init__(self, version: int, commandes: List[str]):
        self.version = version
        self.commandes = commandes

        self.Start = Start
        self.SendCommands = SendCommands

    def Run(self):
        self.clientGandalf = self.Start()

        self.SendCommands(self.clientGandalf, self.version, self.commandes)
