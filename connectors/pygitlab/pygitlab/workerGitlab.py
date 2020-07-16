

#from ...libraries.pyclient import ClientGandalf
#from .workers import WorkerProject
from pygitlab.workers import WorkerProject
# Import de la classe worker de base pour python
from pyworker import Worker
from pygitlab.client.clientGitlab import ClientGitlab

import json


class WorkerGitlab(Worker):
    def Execute(self, version):
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
        workerProject.start()

        configFile.close()
        pass


#MAIN
commands = ["COMMAND_1","COMMAND_2"]
version = 2

workerGitlab = WorkerGitlab(version, commands)
workerGitlab.Run()



