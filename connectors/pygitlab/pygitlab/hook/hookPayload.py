
class AddHookPayload :
    def __init__(self, url, token, push_events, tag_push_events, merge_requests_events, repository_update_events, enable_ssl_verification): # constructor
        self.url= url
        self.token= token
        self.push_events=push_events
        self.tag_push_events=tag_push_events
        self.merge_requests_events= merge_requests_events
        self.repository_update_events= repository_update_events
        self.enable_ssl_verification=enable_ssl_verification

class DeleteHookPayload :
    def __init__(self, hook_id): # constructor
        self.hook_id= hook_id