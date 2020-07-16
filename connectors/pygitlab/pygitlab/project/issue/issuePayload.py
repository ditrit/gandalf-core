class IssuePayload :
    def __init__(self, clientGitlab, title, body, project_id): # constructor
        self.clientGitlab= clientGitlab
        self.title= title
        self.body= body
        self.project_id=project_id
