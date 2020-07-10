class ClientProject():
    def __init__(self,project_id, clientGitlab ):
        self.project_id = project_id
        self.clientGitlab = clientGitlab
        self.project = clientGitlab.client.projects.get(project_id)

    def CreateIssue(self, title, body):
        self.project.issues.create({'title' : title, 'description' : body})
        #ERROR CHECKING TODO

        return True