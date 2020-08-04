
# import de fonction lié au repo
import json
from threading import Thread
from pygitlab.project import projectPayload
from pygitlab.project import project
from pygitlab.project.issue import issuePayload
from pygitlab.project.issue import issue
from pygitlab.client.clientGitlab import ClientGitlab
from pyclient import ClientGandalf
from pyclient.models import Options



class WorkerProject(Thread):

    def __init__(self, clientGitlab, clientGandalf, version):

        Thread.__init__(self)
        self.clientGandalf = clientGandalf
        self.clientGitlab = clientGitlab
        self.version = version
        

    def Run(self):
        CreateProjectThread = Thread(target=self.CreateProject, args=(self,))
        CreateProjectThread.start()
    

    def CreateProject(self):
        id = self.clientGandalf.CreateIteratorCommand()

        while True:
            command = self.clientGandalf.WaitCommand("CREATE_PROJECT", id, self.version)

            jsonProjectPayload = json.load(command.GetPayload())
            createProjectPayload = projectPayload.CreateProjectPayload(jsonProjectPayload)

            # TODO ERROR CHECKING, CHECK IF THE ISSUEPAYLOAD IS FULL
            if createProjectPayload != "":
                

                result = project.CreateProject(self.clientGitlab, createProjectPayload.name)

                if result :
                    self.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), Options("",""))
                else:
                    self.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), Options("",""))
                    
                    

