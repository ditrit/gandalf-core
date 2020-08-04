
# import de fonction lié au repo
import json
from threading import Thread
from pygitlab.hook import hookPayload
from pygitlab.hook import hook
from pygitlab.client.clientGitlab import ClientGitlab
from pyclient import ClientGandalf
from pyclient.models import Options



class WorkerAddHook(Thread):

    def __init__(self, clientGitlab, clientGandalf, version):

        Thread.__init__(self)
        self.clientGandalf = clientGandalf
        self.clientGitlab = clientGitlab
        self.version = version
        
    

    def Run(self):
        AddHookThread = Thread(target=self.AddHook, args=(self,))
        AddHookThread.start()


    def AddHook(self):
        id = self.clientGandalf.CreateIteratorCommand()

        while True:
            command = self.clientGandalf.WaitCommand("ADD_HOOK", id, self.version)

            jsonHookPayload = json.load(command.GetPayload())
            addHookPayload = hookPayload.AddHookPayload(jsonHookPayload)

            # TODO ERROR CHECKING, CHECK IF THE ISSUEPAYLOAD IS FULL
            if addHookPayload != "":
                

                result = hook.AddHook(self.clientGitlab, addHookPayload.url, addHookPayload.token, addHookPayload.push_events, addHookPayload.tag_push_events, addHookPayload.merge_requests_events, addHookPayload.repository_update_events, addHookPayload.enable_ssl_verification)

                if result :
                    self.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), Options("",""))
                else:
                    self.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), Options("",""))


class WorkerDeleteHook(Thread):

    def __init__(self, clientGitlab, clientGandalf, version):

        Thread.__init__(self)
        self.clientGandalf = clientGandalf
        self.clientGitlab = clientGitlab
        self.version = version
        
    

    def Run(self):
        DeleteHookThread = Thread(target=self.DeleteHook, args=(self,))
        DeleteHookThread.start()

    def DeleteHook(self):
            id = self.clientGandalf.CreateIteratorCommand()
    
            while True:
                command = self.clientGandalf.WaitCommand("DELETE_HOOK", id, self.version)
    
                jsonHookPayload = json.load(command.GetPayload())
                deleteHookPayload = hookPayload.DeleteHookPayload(jsonHookPayload)
    
                # TODO ERROR CHECKING, CHECK IF THE ISSUEPAYLOAD IS FULL
                if deleteHookPayload != "":
                    
    
                    result = hook.DeleteHook(self.clientGitlab, deleteHookPayload.hook_id)
    
                    if result :
                        self.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), Options("",""))
                    else:
                        self.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), Options("",""))
    
