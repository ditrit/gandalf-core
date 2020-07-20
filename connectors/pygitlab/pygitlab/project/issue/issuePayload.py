class CreateIssuePayload :
    def __init__(self, project_id, title, body): # constructor
        self.title= title
        self.body= body
        self.project_id=project_id
