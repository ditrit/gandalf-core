

#from ...libraries.pyclient import ClientGandalf
from pyworker import Worker
from pyworker import WorkerVisionning
from pygitlab.workers import WorkerProject
from pygitlab.workers import WorkerAddHook
from pygitlab.workers import WorkerDeleteHook
from pygitlab.workers import WorkerAddMember
from pygitlab.workers import WorkerRemoveMember

from pygitlab.client.clientGitlab import ClientGitlab

from interface import implements
import sys 
import json


class WorkerGitlab(Worker, implements(WorkerVisionning)):
    
    def __init__(self, version: int, commands: list[str]):
        super().__init__(version, commands)
    
    
    def CreateProject(self, version):
        
        clientGitlab= ClientGitlab(url = config.url, token = config.token)
        if not clientGitlab.isValidClient() :
            raise ValueError('Invalid client')
        
         # Thread creation
        workerProject = WorkerProject(self.clientGitlab, self.clientGandalf, self.version)
        workerProject.WorkerProject.Run()

        pass
    
    def AddMember(self, version):
        
        clientGitlab= ClientGitlab(url = config.url, token = config.token)
        if not clientGitlab.isValidClient() :
            raise ValueError('Invalid client')
        
         # Thread creation
        workerAddMember = WorkerAddMember(self.clientGitlab, self.clientGandalf, self.version)
        workerAddMember.WorkerAddMember.Run()

        pass
    
    def RemoveMember(self, version):
        
        clientGitlab= ClientGitlab(url = config.url, token = config.token)
        if not clientGitlab.isValidClient() :
            raise ValueError('Invalid client')
        
         # Thread creation
        workerRemoveMember = WorkerRemoveMember(self.clientGitlab, self.clientGandalf, self.version)
        workerRemoveMember.WorkerRemoveMember.Run()

        pass
    

    
    def AddHook(self, version):
        
        clientGitlab= ClientGitlab(url = config.url, token = config.token)
        if not clientGitlab.isValidClient() :
            raise ValueError('Invalid client')
        
         # Thread creation
        workerAddHook = WorkerAddHook(self.clientGitlab, self.clientGandalf, self.version)
        workerAddHook.WorkerAddHook.Run()

        pass
    
    def DeleteHook(self, version):
        
        clientGitlab= ClientGitlab(url = config.url, token = config.token)
        if not clientGitlab.isValidClient() :
            raise ValueError('Invalid client')
        
         # Thread creation
        workerDeleteHook = WorkerDeleteHook(self.clientGitlab, self.clientGandalf, self.version)
        workerDeleteHook.WorkerDeleteHook.Run()

        pass
    


#MAIN
commands = list()
version = int()

config = json.loads(sys.stdin.read())

workerGitlab = Worker(version, commands) 
workerGitlab.WorkerVisionning.Run()



