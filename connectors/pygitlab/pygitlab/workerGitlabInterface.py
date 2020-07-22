#imports
from pyworker import Worker



class WorkerGitlabInterface:
        
    def CreateProject(self, clientGandalf, version):
        pass

    def CreateHook(self, clientGandalf, version):
        pass
    
    def CreateUser(self, clientGandalf, version):
        pass
        
        
    def Run(self):
        self.Worker.Run()
        #done := make(chan bool)??
        self.CreateProject(self.clientGandalf, self.version)
        self.CreateHook(self.clientGandalf, self.version)
        self.CreateUser(self.clientGandalf, self.version)
        #<-done??

        
    