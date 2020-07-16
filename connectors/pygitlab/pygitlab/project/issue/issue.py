

def CreateIssue(clientGitlab, project_id, title, body):
    project = clientGitlab.client.projects.get(project_id)
    project.issues.create({'title' : title, 'description' : body})
    #ERROR CHECKING TODO
    return True