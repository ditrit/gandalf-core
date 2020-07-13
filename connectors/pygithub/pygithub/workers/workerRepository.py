# import de fonction lié au repo
import json
from threading import Thread

from pygithub.repository.issue.issuePayload import CreateIssuePayload
from pygithub.repository.repositoryClient import ClientRepository
from pygithub.client.clientGithub import ClientGithub
from pyclient.ClientGandalf import ClientGandalf
from pyclient.models import Options


class WorkerRepository(Thread):

    def __init__(self, clientGithub, clientGandalf, version):

        Thread.__init__(self)
        self.clientGandalf = clientGandalf
        self.clientGithub = clientGithub
        self.version = version
        
    

    def Run(self):
        CreateIssueThread = Thread(target=self.CreateIssue, args=(self,))
        CreateIssueThread.start()


    def CreateIssue(self):
        id = self.clientGandalf.CreateIteratorCommand()

        while True:
            command = self.clientGandalf.WaitCommand("CREATE_ISSUE", id, version)

            jsonIssuePayload = json.load(command.GetPayload())
            issuePayload = CreateIssuePayload(jsonIssuePayload)

            # TODO ERROR CHECKING, CHECK IF THE ISSUEPAYLOAD IS FULL
            if issuePayload != "":
                clientRepository = ClientRepository(issuePayload.RepositoryName, self.clientGithub)

                result = clientRepository.CreateIssue(issuePayload.title,issuePayload.body)

                if result :
                    self.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), Options("",""))
                else:
                    self.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), Options("",""))
