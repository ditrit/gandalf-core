import gitlab

def CreateProject(clientGitlab, name):
    project = clientGitlab.client.projects.create({'name': name})
    #TODO ERROR CHECKING 
    return project, True

def AddMember(clientGitlab, user_id, project_id):
    project = clientGitlab.client.projects.get(project_id)
    project.members.create({'user_id': user_id, 'access_level':
                                 gitlab.DEVELOPER_ACCESS})
    #TODO ERROR CHECKING 
    return True

def DeleteMember(clientGitlab, user_id, project_id):
    #TODO
    project = clientGitlab.client.projects.get(project_id)
    project.members.delete(user_id)
    #TODO ERROR CHECKING 
    return True