class CreateIssuePayload :

    def __init__(self, payload):
        self.title = payload['title']
        self.body = payload['body']
        self.repositoryName = payload['repository']
