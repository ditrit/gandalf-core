# import de fonction lié au repo
import json
from threading import Thread

from ..repository.issue.issuePayload import CreateIssuePayload
from ..repository.repositoryClient import ClientRepository
from ..client.client_github import ClientGithub
from ....libraries.pyclient import ClientGandalf
from ....libraries.pyclient.models import Options


class workerRepository(Thread):

    def __init__(self,clientGandalf, token, version):

        Thread.__init__(self)
        self.clientGandalf = clientGandalf
        self.clientGithub = ClientGithub(token)
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

                result = clientRepository.CreateIssue(issuePayload.title,issuePayload.body, issuePayload.number)

                if result :
                    self.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), Options("",""))
                else:
                    self.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), Options("",""))
