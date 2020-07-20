
def AddHook(clientGithub, repositoryName:str, hookName:str, events:list, host:str, endpoint:str, active:bool, contentType: str):

    config = {
        "url": "http://{host}/{endpoint}".format(host=host, endpoint = endpoint),
        "content_type:": contentType
        }

    repository = clientGithub.client.get_repo(repositoryName)
    repository.create_hook(name = hookName, config = config, events = events, active = active)

    #ERROR HANDLING TODO

    return True


def GetHook(clientGithub, repositoryName, hookID):


    repository = clientGithub.client.get_repo(repositoryName)
    hook = repository.get_hook(hookID)

    #TODO ERROR HANDLING
    
    return hook

def GetHooks(clientGithub, repositoryName):

    repository = clientGithub.client.get_repo(repositoryName)
    hooks = repository.get_hook()

    #TODO ERROR HANDLING

    return hooks

def DeleteHook(clientGithub, repositoryName, hookID):

    repository = clientGithub.client.get_repo(repositoryName)
    hook = repository.get_hook(hookID)
    hook.delete()

    #TODO ERROR HANDLING
    
    return True

'''

def EditHook(clientGithub, ):

    repository = clientGithub.client.get_repo(repositoryName)
    hook = repository.get_hook(hookID)
    hook.edit(name = hookName, conifg = config, events = events, add_events = add_events, active = active)

    #TODO ERROR HANDLING

    pass

'''
