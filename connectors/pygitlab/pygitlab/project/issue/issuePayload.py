class IssuePayload :
    def __init__(self, clientGitlab, project_id, title, body): # constructor
        self.clientGitlab= clientGitlab
        self.title= title
        self.body= body
        self.project_id=project_id
