
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
        
        print("Listening on localhost:", self.port)
        print(server.application.run(debug=True, host='0.0.0.0'))


                        
    
    
    def GetUrl(self):
        return "http://" + self.rooturl + self.url

