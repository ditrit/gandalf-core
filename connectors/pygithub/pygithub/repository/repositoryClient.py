class ClientRepository():
    def __init__(self,repositoryName, clientGithub ):
        self.repositoryName = repositoryName
        self.clientGithub = clientGithub
        self.repository = clientGithub.client.get_repo()

    def CreateIssue(self, title, body):
        self.repository.create_issue(title=title, body = body)
        #ERROR CHECKING TODO

        return True
