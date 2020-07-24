

#from ...libraries.pyclient import ClientGandalf
from pyworker import Worker
from pyworker import WorkerVisionning
from pygitlab.workers import WorkerProject
from pygitlab.workers import WorkerHook
from pygitlab.workers import WorkerUser
from pygitlab.client.clientGitlab import ClientGitlab

import sys 
import json


class WorkerGitlab(WorkerVisionning):
    
    
    def CreateProject(self, version):
        
        clientGitlab= ClientGitlab(config.token, config.url)
        if not clientGitlab.isValidClient() :
            raise ValueError('Invalid client')
        
         # Thread creation
        workerProject = WorkerProject(self.clientGitlab, self.clientGandalf, self.version)
        workerProject.WorkerProject.Run()

        pass
    
    def CreateHook(self, version):
        
        clientGitlab= ClientGitlab(config.token, config.url)
        if not clientGitlab.isValidClient() :
            raise ValueError('Invalid client')
        
         # Thread creation
        workerHook = WorkerHook(self.clientGitlab, self.clientGandalf, self.version)
        workerHook.WorkerHook.Run()

        pass
    
    def CreateUser(self, version):
        
        clientGitlab= ClientGitlab(config.token, config.url)
        if not clientGitlab.isValidClient() :
            raise ValueError('Invalid client')
        
         # Thread creation
        workerUser = WorkerUser(self.clientGitlab, self.clientGandalf, self.version)
        workerUser.WorkerUser.Run()

        pass
        
    


#MAIN
commands = list()
version = int()

config = json.loads(sys.stdin.read())

workerGitlab = Worker(version, commands) 
workerGitlab.WorkerGitlabInterface.Run()



