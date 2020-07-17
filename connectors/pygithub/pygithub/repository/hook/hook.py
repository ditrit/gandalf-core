
def AddHook(clientGithub, repositoryName:str, hookName:str, events:list, host:str, endpoint:str, active:bool, contentType: str):

    config = {
        "url": "http://{host}/{endpoint}".format(host=host, endpoint = endpoint),
        "content_type:": contentType
        }

    repository = clientGithub.client.get_repo(repositoryName)
    repository.create_hook(name = hookName, config = config, events = events, active = active)

    #ERROR HANDLING TODO

    return True


def GetHook():

    #TODO
    
    pass

def GetHooks():

    #TODO

    pass

def EditHook():

    #TODO

    pass
