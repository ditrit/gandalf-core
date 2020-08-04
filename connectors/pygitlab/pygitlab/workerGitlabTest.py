

#from ...libraries.pyclient import ClientGandalf
from pyworker import Worker
from pyworker import WorkerVisionning
from pygitlab.workers import WorkerProject
from pygitlab.workers import WorkerAddHook
from pygitlab.workers import WorkerDeleteHook
from pygitlab.workers import WorkerAddMember
from pygitlab.workers import WorkerRemoveMember
from pygitlab.workers import WorkerServer
from pygitlab.workers import WorkerTemplate
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
    
    def AddMember(self, version):
        
        clientGitlab= ClientGitlab(config.token, config.url)
        if not clientGitlab.isValidClient() :
            raise ValueError('Invalid client')
        
         # Thread creation
        workerAddMember = WorkerAddMember(self.clientGitlab, self.clientGandalf, self.version)
        workerAddMember.WorkerAddMember.Run()

        pass
    
    def RemoveMember(self, version):
        
        clientGitlab= ClientGitlab(config.token, config.url)
        if not clientGitlab.isValidClient() :
            raise ValueError('Invalid client')
        
         # Thread creation
        workerRemoveMember = WorkerRemoveMember(self.clientGitlab, self.clientGandalf, self.version)
        workerRemoveMember.WorkerRemoveMember.Run()

        pass
    
    def CreateTemplate(self, version):
        
        clientGitlab= ClientGitlab(config.token, config.url)
        if not clientGitlab.isValidClient() :
            raise ValueError('Invalid client')
        
         # Thread creation
        workerTemplate = WorkerTemplate(self.clientGitlab, self.clientGandalf, self.version)
        workerTemplate.WorkerTemplate.Run()

        pass
    

    
    def AddHook(self, version):
        
        clientGitlab= ClientGitlab(config.token, config.url)
        if not clientGitlab.isValidClient() :
            raise ValueError('Invalid client')
        
         # Thread creation
        workerAddHook = WorkerAddHook(self.clientGitlab, self.clientGandalf, self.version)
        workerAddHook.WorkerAddHook.Run()

        pass
    
    def DeleteHook(self, version):
        
        clientGitlab= ClientGitlab(config.token, config.url)
        if not clientGitlab.isValidClient() :
            raise ValueError('Invalid client')
        
         # Thread creation
        workerDeleteHook = WorkerDeleteHook(self.clientGitlab, self.clientGandalf, self.version)
        workerDeleteHook.WorkerDeleteHook.Run()

        pass
    
    def NewTag(self, version):
        
        clientGitlab= ClientGitlab(config.token, config.url)
        if not clientGitlab.isValidClient() :
            raise ValueError('Invalid client')
        
         # Thread creation
        workerTag = WorkerServer(self.clientGitlab, self.clientGandalf, self.version)
        workerTag.WorkerServer.Run()

        pass
    
    def MergeRequest(self, version):
        
        clientGitlab= ClientGitlab(config.token, config.url)
        if not clientGitlab.isValidClient() :
            raise ValueError('Invalid client')
        
         # Thread creation
        workerMerge = WorkerServer(self.clientGitlab, self.clientGandalf, self.version)
        workerMerge.WorkerServer.Run()

        pass
        
    


#MAIN
commands = list()
version = int()

config = json.loads(sys.stdin.read())

workerGitlab = Worker(version, commands) 
workerGitlab.WorkerVisionning.Run()



