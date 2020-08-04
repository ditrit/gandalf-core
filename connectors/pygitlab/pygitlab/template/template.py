import gitlab

def CreateTemplate(clientGitlab, repository, project, name):
    project = clientGitlab.client.projects.create({'name': name})
    #TODO ERROR CHECKING 
    return True