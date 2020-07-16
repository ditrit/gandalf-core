
# import de fonction lié au repo
import json
from threading import Thread
from pygitlab.project import ProjectPayload
from pygitlab.project import project
from pygitlab.project.issue import IssuePayload
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
        CreateIssueThread = Thread(target=self.CreateIssue, args=(self,))
        CreateProjectThread = Thread(target=self.CreateProject, args=(self,))
        CreateIssueThread.start()
        CreateProjectThread.start()


    def CreateProject(self):
        id = self.clientGandalf.CreateIteratorCommand()

        while True:
            command = self.clientGandalf.WaitCommand("CREATE_PROJECT", id, self.version)

            jsonProjectPayload = json.load(command.GetPayload())
            projectPayload = ProjectPayload(jsonProjectPayload)

            # TODO ERROR CHECKING, CHECK IF THE ISSUEPAYLOAD IS FULL
            if projectPayload != "":
                

                result = project.CreateProject(projectPayload.clientGitlab, projectPayload.name)

                if result :
                    self.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), Options("",""))
                else:
                    self.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), Options("",""))
                    
                    
    def CreateIssue(self):
        id = self.clientGandalf.CreateIteratorCommand()

        while True:
            command = self.clientGandalf.WaitCommand("CREATE_ISSUE", id, self.version)

            jsonIssuePayload = json.load(command.GetPayload())
            issuePayload = IssuePayload(jsonIssuePayload)

            # TODO ERROR CHECKING, CHECK IF THE ISSUEPAYLOAD IS FULL
            if issuePayload != "":
                

                result = issue.CreateIssue(issuePayload.clientGitlab, issuePayload.project_id, issuePayload.title, issuePayload.body)

                if result :
                    self.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), Options("",""))
                else:
                    self.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), Options("",""))

