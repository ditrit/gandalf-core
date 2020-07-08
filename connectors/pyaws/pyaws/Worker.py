#! /usr/bin/env python3
# coding: utf-8

from typing import List, Callable

from pyclient.ClientGandalf import ClientGandalf

from .functions.Start import Start
from .functions.SendCommands import SendCommands

class Worker:

    @property
    def version(self):
        return self._version

    @version.setter
    def version(self, value):
        self._version = value

    @property
    def commandes(self):
        return self._commandes

    @commandes.setter
    def commandes(self, value):
        self._commandes = value

    @property
    def clientGandalf(self):
        return self._clientGandalf

    @clientGandalf.setter
    def clientGandalf(self, value):
        self._clientGandalf = value

    def Start(self, clientGandalf: ClientGandalf):
        pass

    def SendCommands(self, clientGandalf: ClientGandalf, version: int, commands: List[str]):
        pass

    def Execute(self, clientGandalf: ClientGandalf, version: int):
        pass

    def __init__(self, version: int, commandes: List[str]):
        self.version = version
        self.commandes = commandes    

        self.Start = Start
        self.SendCommands = SendCommands

    def Run(self):
        self.clientGandalf = self.Start()

        self.SendCommands(self.clientGandalf, self.version, self.commandes)

        self.Execute(self.clientGandalf, self.version)

        # <-done need to be translated

        

        