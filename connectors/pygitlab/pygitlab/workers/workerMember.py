# import de fonction lié au repo
import json
from threading import Thread
from pygitlab.project import projectPayload
from pygitlab.project import project
from pygitlab.project.issue import issuePayload
from pygitlab.project.issue import issue
from pygitlab.client.clientGitlab import ClientGitlab
from pyclient import ClientGandalf
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
                    
    
                    result = project.AddMember(self.clientGitlab, addMemberPayload.new_member, addMemberPayload.project_id)
    
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
        CreateProjectThread = Thread(target=self.CreateProject, args=(self,))
        AddMemberThread = Thread(target=self.AddMember, args=(self,))
        RemoveMemberThread = Thread(target=self.RemoveMember, args=(self,))
        CreateProjectThread.start()
        AddMemberThread.start()
        RemoveMemberThread.start()
    
    def RemoveMember(self):
                id = self.clientGandalf.CreateIteratorCommand()
        
                while True:
                    command = self.clientGandalf.WaitCommand("REMOVE_MEMBER", id, self.version)
        
                    jsonProjectPayload = json.load(command.GetPayload())
                    removeMemberPayload = projectPayload.removeMemberProjectPayload(jsonProjectPayload)
        
                    # TODO ERROR CHECKING, CHECK IF THE ISSUEPAYLOAD IS FULL
                    if removeMemberPayload != "":
                        
        
                        result = project.RemoveMember(self.clientGitlab, removeMemberPayload.member, removeMemberPayload.project_id)
        
                        if result :
                            self.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), Options("",""))
                        else:
                            self.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), Options("",""))
        
