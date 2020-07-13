
from pyclient.ClientGandalf import ClientGandalf
# Import de la classe worker de base pour python
from pyworker.Worker import Worker
from pygithub.workers.workerRepository import WorkerRepository
from pygithub.client.clientGithub import ClientGithub



import json

class WorkerGithub(Worker):
    def Execute(self, version):
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
        workerRepository = WorkerRepository(clientGithub, self.clientGandalf, self.version)
        workerRepository.start()

        configFile.close()

if __name__ == "__main__":

    commands = ["COMMAND_1","COMMAND_2"]
    version = 2

    workerGithub = WorkerGithub(version, commands)
    workerGithub.Run()
