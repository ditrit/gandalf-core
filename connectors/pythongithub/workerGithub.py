# from /workers import worker1,worker2 etc

from ...libraries.pyclient import ClientGandalf
from .workers import workerRepository

# Import de la classe worker de base pour python
from ..python import Worker
from .client.client_github import ClientGithub

import json


#MAIN
commands = ["COMMAND_1","COMMAND_2"]
version = 2

workerGithub = WorkerGithub(version, commands)
workerGithub.Run()


class WorkerGithub(Worker):
    def Execute(self, client, version):
        '''
        configuration = { 
            "token" : None,
             }
        '''
        configFile = open('./config.json')
        configuration = json.load(configFile)
        
        clientGithub= ClientGithub(configuration.token)
        if not clientGithub.isValidClient() :
            raise ValueError('Invalid client')
        
        
        

        # Thread creation
        workerRepository = workerRepository(clientGithub, self.clientGandalf, self.version)
        workerRepository.start()

        configFile.close()
        pass
