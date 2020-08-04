
# import de fonction lié au repo
import json
from threading import Thread
from pygitlab.server import ServerPayload
from pygitlab.server import server
from pygitlab.server import ServerUrl
from pygitlab.client.clientGitlab import ClientGitlab
from pyclient import ClientGandalf
from pyclient.models import Options


import logging


#Déclaration de variables
server_address = "localhost"
server_port    = ":3003"

class WorkerServer(Thread):

    def __init__(self, clientGitlab, clientGandalf, version, identity, connections):

        Thread.__init__(self)
        #self.version = version
        #self.clientGitlab = clientGitlab
        self.clientGandalf = ClientGandalf(identity, connections)
        
        self.address = server_address
        self.port = server_port
        self.rooturl = server_address + server_port
        self.url = ServerUrl.returnHashURLS()
        print(self.url)
        
        self.router = server.Hook(self.url, self.clientGandalf)
        

    def Run(self):
        #start the workerServer
        serverThread = Thread(target=self.SendEvent, args=(self,))
        serverThread.start()
        print("Listening on localhost:", self.port)
        print(server.application.run(debug=True, host='0.0.0.0'))


    def SendEvent(self):

        id = self.clientGandalf.CreateIteratorEvent()

        while True:
            event = self.clientGandalf.WaitEvent("Gitlab", "HOOK", id)
            
            #TODO create a server payload?

            jsonEventPayload = json.load(event.GetPayload())
            newEventPayload = serverPayload.eventPayload(jsonEventPayload)

            # TODO ERROR CHECKING, CHECK IF THE ISSUEPAYLOAD IS FULL
            if newEventPayload != "":
                print(newEventPayload)
                
                #a faire differentes classes et indépendnat du hook au cas d'autres se rajoute??
                '''if newEventPayload.event == "new_tag":
    
                    result = event.NewTag(self.clientGitlab, "")
    
                    if result :
                        self.clientGandalf.SendReply(event.GetEvent(), "SUCCES", event.GetUUID(), Options("",""))
                    else:
                        self.clientGandalf.SendReply(event.GetEvent(), "FAIL", event.GetUUID(), Options("",""))
                        
                        
                if newEventPayload.topic == "merge_request":
        
                        result = event.MergeRequest(self.clientGitlab, "")
        
                        if result :
                            self.clientGandalf.SendReply(event.GetEvent(), "SUCCES", event.GetUUID(), Options("",""))
                        else:
                            self.clientGandalf.SendReply(event.GetEvent(), "FAIL", event.GetUUID(), Options("",""))'''
                        
                        
    
    
    def GetUrl(self):
        return "http://" + self.rooturl + self.url

