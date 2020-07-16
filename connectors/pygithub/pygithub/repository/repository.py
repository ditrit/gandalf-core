# Functions related to Repository

def CreateRepository(clientGithub, name, private):

    user = clientGithub.client.get_user()
    user.create_repo(name = name, private = private)
    
    # ERROR HANDLING TODO

    return True
