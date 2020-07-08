# from /workers import worker1,worker2 etc

from ...libraries.pyclient import ClientGandalf
from .workers import workerRepository

# Import de la classe worker de base pour python
from ..python import Worker

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


        # Thread creation
        workerRepository = workerRepository(configuration.token, self.clientGandalf, self.version)
        workerRepository.start()

        configFile.close()
        pass
