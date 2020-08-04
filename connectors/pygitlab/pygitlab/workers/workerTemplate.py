
# import de fonction lié au repo
import json
from threading import Thread
from pygitlab.hook import templatePayload
from pygitlab.hook import template
from pygitlab.client.clientGitlab import ClientGitlab
from pyclient import ClientGandalf
from pyclient.models import Options



class WorkerTemplate(Thread):

    def __init__(self, clientGitlab, clientGandalf, version):

        Thread.__init__(self)
        self.clientGandalf = clientGandalf
        self.clientGitlab = clientGitlab
        self.version = version
        
    

    def Run(self):
        TemplateThread = Thread(target=self.CreateTemplate, args=(self,))
        TemplateThread.start()


    def CreateTemplate(self):
        id = self.clientGandalf.CreateIteratorCommand()

        while True:
            command = self.clientGandalf.WaitCommand("CREATE_TEMPLATE", id, self.version)


            # TODO ERROR CHECKING, CHECK IF THE ISSUEPAYLOAD IS FULL
            

            result = template.CreateTemplate(self.clientGitlab)

            if result :
                self.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), Options("",""))
            else:
                self.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), Options("",""))
                

