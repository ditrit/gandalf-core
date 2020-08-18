# import de fonction lié au repo
import json
from threading import Thread
from pygitlab.project import projectPayload
from pygitlab.project import project

from pyclient.models import Options



class WorkerAddMember(Thread):

    def __init__(self, clientGitlab, clientGandalf, version):

        Thread.__init__(self)
        self.clientGandalf = clientGandalf
        self.clientGitlab = clientGitlab
        self.version = version
        

    def Run(self):
        AddMemberThread = Thread(target=self.AddMember, args=(self,))
        AddMemberThread.start()


    def AddMember(self):
            id = self.clientGandalf.CreateIteratorCommand()
    
            while True:
                command = self.clientGandalf.WaitCommand("ADD_MEMBER", id, self.version)
    
                jsonProjectPayload = json.load(command.GetPayload())
                addMemberPayload = projectPayload.AddMemberProjectPayload(jsonProjectPayload)
    
                # TODO ERROR CHECKING, CHECK IF THE ISSUEPAYLOAD IS FULL
                if addMemberPayload != "":
                    
    
                    result = project.AddMember(clientGitlab = self.clientGitlab,  project_ID=addMemberPayload.projectID, user_email=addMemberPayload.userEmail)
    
                    if result :
                        self.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), Options("",""))
                    else:
                        self.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), Options("",""))
                        
                        
                        
class WorkerRemoveMember(Thread):

    def __init__(self, clientGitlab, clientGandalf, version):

        Thread.__init__(self)
        self.clientGandalf = clientGandalf
        self.clientGitlab = clientGitlab
        self.version = version
        

    def Run(self):
        RemoveMemberThread = Thread(target=self.RemoveMember, args=(self,))
        RemoveMemberThread.start()
    
    def RemoveMember(self):
                id = self.clientGandalf.CreateIteratorCommand()
        
                while True:
                    command = self.clientGandalf.WaitCommand("REMOVE_MEMBER", id, self.version)
        
                    jsonProjectPayload = json.load(command.GetPayload())
                    removeMemberPayload = projectPayload.RemoveMemberProjectPayload(jsonProjectPayload)
        
                    # TODO ERROR CHECKING, CHECK IF THE ISSUEPAYLOAD IS FULL
                    if removeMemberPayload != "":
                        
        
                        result = project.RemoveMember(clientGitlab=self.clientGitlab, project_id=removeMemberPayload.projectID, user_email=removeMemberPayload.userEmail, )
        
                        if result :
                            self.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), Options("",""))
                        else:
                            self.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), Options("",""))
        
