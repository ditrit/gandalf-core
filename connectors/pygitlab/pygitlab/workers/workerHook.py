
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
                

                result = hook.AddHook(clientGitlab=self.clientGitlab, url=addHookPayload.url, token=addHookPayload.token, push_events=addHookPayload.pushEvents, tag_push_events=addHookPayload.tagPushEvents, merge_requests_events=addHookPayload.mergeRequestsEvents, repository_update_events=addHookPayload.repositoryUpdateEvents, enable_ssl_verification=addHookPayload.enableSslVerification)

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
                    
                    #TODO
                    if isAuthorized(self.clientGitlab, "DELETE_HOOK") :  #vérification de l'abilitation de celui qui passe la commande?
    
                        result = hook.DeleteHook(clientGitlab=self.clientGitlab, hook_id=deleteHookPayload.hookID)
        
                        if result :
                            self.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), Options("",""))
                        else:
                            self.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), Options("",""))
        
