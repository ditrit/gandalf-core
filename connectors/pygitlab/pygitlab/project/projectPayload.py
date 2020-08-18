class CreateProjectPayload :
    def __init__(self, name, team, templateName = None): # constructor
        self.name= name
        self.team= team
        self.templateName= templateName
       
        
class AddMemberProjectPayload :
    def __init__(self, projectID, userEmail): # constructor
        self.projectID= projectID
        self.userEmail= userEmail
       
        
class RemoveMemberProjectPayload :
    def __init__(self, projectID, userEmail): # constructor
        self.projectID= projectID
        self.userEmail= userEmail
       
        
        