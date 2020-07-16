# import de fonction lié au repo
import json
from threading import Thread

from pygithub.client.clientGithub import ClientGithub
from pyclient.ClientGandalf import ClientGandalf
from pyclient.models import Options


from pygithub.repository.issue.issuePayload import CreateIssuePayload
from pygithub.repository.issue.issue import CreateIssue

from pygithub.repository.repositoryPayload import CreateRepositoryPayload
from pygithub.repository.repository import CreateRepository




class WorkerRepository(Thread):

    def __init__(self, clientGithub, clientGandalf, version):

        Thread.__init__(self)
        self.clientGandalf = clientGandalf
        self.clientGithub = clientGithub
        self.version = version
        
    

    def Run(self):
        CreateIssueThread = Thread(target=self.CreateIssue, args=(self,))
        CreateRepositoryThread = Thread(target=self.CreateRepository, args=(self,))

        CreateIssueThread.start()
        CreateRepositoryThread.start()


    def CreateIssue(self):
        id = self.clientGandalf.CreateIteratorCommand()

        while True:
            command = self.clientGandalf.WaitCommand("CREATE_ISSUE", id, version)

            jsonIssuePayload = json.load(command.GetPayload())
            issuePayload = CreateIssuePayload(jsonIssuePayload)

            # TODO ERROR CHECKING, CHECK IF THE ISSUEPAYLOAD IS FULL
            if issuePayload != "":

                result = CreateIssue(self.clientGithub, issuePayload.repositoryName ,issuePayload.title,issuePayload.body)

                if result :
                    self.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), Options("",""))
                else:
                    self.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), Options("",""))

    def CreateRepository(self):
        id = self.clientGandalf.CreateIteratorCommand()

        while True:
            command = self.clientGandalf.WaitCommand("CREATE_REPOSITORY", id, version)

            jsonRepositoryPayload = json.load(command.GetPayload())
            repositoryPayload = CreateRepositoryPayload(jsonRepositoryPayload)

            # TODO ERROR CHECKING, CHECK IF THE ISSUEPAYLOAD IS FULL
            if repositoryPayload != "":

                result = CreateRepository(self.clientGithub, repositoryPayload.name ,repositoryPayload.private)

                if result :
                    self.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), Options("",""))
                else:
                    self.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), Options("",""))
