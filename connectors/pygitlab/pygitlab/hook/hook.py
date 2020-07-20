
def ListHooks(clientGitlab) :
    hooks = clientGitlab.client.hooks.list()
    #TODO ERROR CHECKING 
    return hooks, True

def AddHook(clientGitlab, url, token, push_events, tag_push_events, merge_requests_events, repository_update_events, enable_ssl_verification):
    hook = clientGitlab.client.hooks.create({'url': url, 'token': token, 'push_events': push_events, 'tag_push_events': tag_push_events,
                                            'merge_requests_events': merge_requests_events, 'repository_update_events': repository_update_events,
                                            'enable_ssl_verification': enable_ssl_verification})
    #TODO ERROR CHECKING 
    return hook, True

def DeleteHook(clientGitlab, hook_id) :
    clientGitlab.client.hooks.delete(hook_id)
    #TODO ERROR CHECKING 
    return True
