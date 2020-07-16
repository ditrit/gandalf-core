
def CreateProject(clientGitlab, name):
    project = clientGitlab.client.projects.create({'name': name})
    #ERROR CHECKING TODO
    return project, True