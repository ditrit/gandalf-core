# from /workers import worker1,worker2 etc

from ...libraries.pyclient import ClientGandalf

# Import de la classe worker de base pour python
from ..python import Worker

import json


#MAIN
commands = ["COMMAND_1","COMMAND_2"]
version = 2

class WorkerGithub(Worker):
    def Execute(self, client, version):

        configuration = { 
            "token" : None,
             }
        
        configFile = open('./config.json')
        configuration = json.load(configFile)


        #Start Worker

        configFile.close()
        pass
