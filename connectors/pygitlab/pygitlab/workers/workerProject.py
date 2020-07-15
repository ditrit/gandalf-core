
# import de fonction lié au repo
import json
from threading import Thread
from pygitlab.project.ClientProject import ClientProject
from pygitlab.project.issue import issuePayload
#from pygitlab.issuePayload import CreateIssuePayload
#from ..project.issue.issuePayload import CreateIssuePayload
#from pygitlab.ClientProject import ClientProject
#from ..client.clientGitlab import ClientGitlab
#from ....libraries.pyclient import ClientGandalf
#from ....libraries.pyclient.models import Options


class WorkerProject(Thread):

    def __init__(self, clientGitlab, clientGandalf, version):

        Thread.__init__(self)
        self.clientGandalf = clientGandalf
        self.clientGitlab = clientGitlab
        self.version = version
        
    

    def Run(self):
        CreateIssueThread = Thread(target=self.CreateIssue, args=(self,))
        CreateIssueThread.start()


    def CreateIssue(self):
        id = self.clientGandalf.CreateIteratorCommand()

        while True:
            command = self.clientGandalf.WaitCommand("CREATE_ISSUE", id, self.version)

            jsonIssuePayload = json.load(command.GetPayload())
            issuePayload = CreateIssuePayload(jsonIssuePayload)

            # TODO ERROR CHECKING, CHECK IF THE ISSUEPAYLOAD IS FULL
            if issuePayload != "":
                clientProject = ClientProject(issuePayload.ProjectName, self.clientGithub)

                result = clientProject.CreateIssue(issuePayload.title,issuePayload.body)

                if result :
                    self.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), Options("",""))
                else:
                    self.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), Options("",""))
