

#from ...libraries.pyclient import ClientGandalf
from pyworker import Worker
from pygitlab import WorkerGitlabInterface
from pygitlab.workers import WorkerProject
from pygitlab.workers import WorkerHook
from pygitlab.workers import WorkerUser
from pygitlab.client.clientGitlab import ClientGitlab

import json


class WorkerGitlab(WorkerGitlabInterface):
    
    
    def CreateProject(self, version):
        '''
        configuration = { 
            "token" : None,
             }
        '''
        configFile = open('./config.json')
        configuration = json.load(configFile)
        
        clientGitlab= ClientGitlab(configuration.token, configuration.url)
        if not clientGitlab.isValidClient() :
            raise ValueError('Invalid client')
        
         # Thread creation
        workerProject = WorkerProject(self.clientGitlab, self.clientGandalf, self.version)
        workerProject.WorkerProject.Run()

        configFile.close()
        pass
    
    def CreateHook(self, version):
        '''
        configuration = { 
            "token" : None,
             }
        '''
        configFile = open('./config.json')
        configuration = json.load(configFile)
        
        clientGitlab= ClientGitlab(configuration.token, configuration.url)
        if not clientGitlab.isValidClient() :
            raise ValueError('Invalid client')
        
         # Thread creation
        workerHook = WorkerHook(self.clientGitlab, self.clientGandalf, self.version)
        workerHook.WorkerHook.Run()

        configFile.close()
        pass
    
    def CreateUser(self, version):
        '''
        configuration = { 
            "token" : None,
             }
        '''
        configFile = open('./config.json')
        configuration = json.load(configFile)
        
        clientGitlab= ClientGitlab(configuration.token, configuration.url)
        if not clientGitlab.isValidClient() :
            raise ValueError('Invalid client')
        
         # Thread creation
        workerUser = WorkerUser(self.clientGitlab, self.clientGandalf, self.version)
        workerUser.WorkerUser.Run()

        configFile.close()
        pass
        


#MAIN
commands = ["COMMAND_1","COMMAND_2"]
version = 2

#gérer les entrées standards ici

workerGitlab = Worker(version, commands) 
workerGitlab.WorkerGitlabInterface.Run()



