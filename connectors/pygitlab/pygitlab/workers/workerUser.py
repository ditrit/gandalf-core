
# import de fonction lié au repo
import json
from threading import Thread
from pygitlab.user import UserPayload
from pygitlab.user import user
from pygitlab.client.clientGitlab import ClientGitlab
from pyclient import ClientGandalf
from pyclient.models import Options



class WorkerUser(Thread):
    
    def __init__(self, clientGitlab, clientGandalf, version):

        Thread.__init__(self)
        self.clientGandalf = clientGandalf
        self.clientGitlab = clientGitlab
        self.version = version
        
    

    def Run(self):
        CreateUserThread = Thread(target=self.CreateIssue, args=(self,))
        CreateUserThread.start()


    def CreateUser(self):
        id = self.clientGandalf.CreateIteratorCommand()

        while True:
            command = self.clientGandalf.WaitCommand("CREATE_USER", id, self.version)

            jsonUserPayload = json.load(command.GetPayload())
            userPayload = UserPayload(jsonUserPayload)

            # TODO ERROR CHECKING, CHECK IF THE ISSUEPAYLOAD IS FULL
            if userPayload != "":
                

                result = user.CreateUser(self.clientGitlab, userPayload.name)

                if result :
                    self.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), Options("",""))
                else:
                    self.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), Options("",""))
                    