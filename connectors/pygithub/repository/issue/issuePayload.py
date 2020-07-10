class CreateIssuesPayload :
    def __init__(self, title, body, number, repository): # constructor
        self.title= title
        self.body= body
        self.number= number
        self.repository=repository
