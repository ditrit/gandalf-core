# Functions related to Repository

def CreateRepository(clientGithub, name, private):

    user = clientGithub.client.get_user()
    user.create_repo(name = name, private = private)
    
    # ERROR HANDLING TODO

    return True


def AddUserToCollaborators(clientGithub, repositoryName, username, permission):
    #permission = 'pull', 'push' or 'admin'
    repository = clientGithub.client.get_repo(repositoryName)
    repository.add_to_collaborator(collaborator = username, permission = permission)

    # ERROR HANDLING TODO

    return True

def RemoveUserFromCollaborators( clientGithub, repositoryName, username):

    repository = repository = clientGithub.client.get_repo(repositoryName)
    repository.remove_from_collaborators( collaborator = username )

    # ERROR HANDLING TODO

    return True
