#! /usr/bin/env python3
# coding: utf-8

from pyworker.Worker import Worker

from typing import List
from abc import ABCMeta, abstractmethod
from pyclient.ClientGandalf import ClientGandalf

class WorkerInterface(Worker, metaclass=ABCMeta):
    
    @abstractmethod
    def Execute(self, clientGandalf: ClientGandalf, version: int):
        raise NotImplementedError

    def __init__(self, version: int, commandes: List[str]):
        super().__init__(version, commandes)

    def Run(self):
        super().Run()

        self.Execute(self.clientGandalf, self.version)