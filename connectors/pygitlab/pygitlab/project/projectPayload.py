class CreateProjectPayload :
    def __init__(self, project_id, name): # constructor
        self.project_id= project_id
        self.name= name
       
        
class AddMemberProjectPayload :
    def __init__(self, project_id, user_id): # constructor
        self.project_id= project_id
        self.user_id= user_id
       