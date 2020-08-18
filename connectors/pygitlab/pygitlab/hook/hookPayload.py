
class AddHookPayload :
    def __init__(self, url, token, pushEvents, tagPushEvents, mergeRequestsEvents, repositoryUpdateEvents, enableSslVerification): # constructor
        self.url= url
        self.token= token
        self.pushEvents=pushEvents
        self.tagPushEvents=tagPushEvents
        self.mergeRequestsEvents= mergeRequestsEvents
        self.repositoryUpdateEvents= repositoryUpdateEvents
        self.enableSslVerification=enableSslVerification

class DeleteHookPayload :
    def __init__(self, hookID): # constructor
        self.hookID= hookID
        