from pyworker import Worker
from pyclient.ClientGandalf import ClientGandalf

from abc import ABCMeta, abstractmethod

class WorkerVisionning(Worker, metaclass=abc.ABCMeta) :
    
    def __init__(self, version: int, commands: list[str]):
        super().__init__(version, commands)
        
    @abc.abstractmethod   
    def CreateProject(self, clientGandalf: ClientGandalf, version: int):
        pass
    @abc.abstractmethod
    def CreateHook(self, clientGandalf: ClientGandalf, version: int):
        pass
    @abc.abstractmethod
    def CreateUser(self, clientGandalf: ClientGandalf, version: int):
        pass
        
        
    def Run(self):
        super().Run()
        self.CreateProject(self.clientGandalf, self.version)
        self.CreateHook(self.clientGandalf, self.version)
        self.CreateUser(self.clientGandalf, self.version)
