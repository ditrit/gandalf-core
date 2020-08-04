import hashlib

class Urls:
    def __init__(self, static_path, gandalf_path, hook_path):
        self.static_path = static_path
        self.gandalf_path = gandalf_path
        self.hook_path = gandalf_path
    
    
    def returnHashURLS():
        sha_512 = hashlib.sha512(b"/gandalf/gitlab/").hexdigest()
        return "/" + sha_512
    
        
    