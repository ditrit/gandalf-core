# from /workers import worker1,worker2 etc

from ...libraries.pyclient import ClientGandalf
from .workers import WorkerProject

# Import de la classe worker de base pour python
from ..python import Worker
from .client.clientGitlab import ClientGitlab

import json


#MAIN
commands = ["COMMAND_1","COMMAND_2"]
version = 2

workerGitlab = WorkerGitlab(version, commands)
workerGitlab.Run()


class WorkerGitlab(Worker):
    def Execute(self, client, version):
        '''
        configuration = { 
            "token" : None,
             }
        '''
        configFile = open('./config.json')
        configuration = json.load(configFile)
        
        clientGitlab= ClientGitlab(configuration.token)
        if not clientGitlab.isValidClient() :
            raise ValueError('Invalid client')
        
        
        

        # Thread creation
        workerProject = WorkerProject(clientGitlab, self.clientGandalf, self.version)
        workerProject.start()

        configFile.close()
        pass
