from gitlab import Gitlab
from gitlab import exceptions 
import logging


class ClientGitlabBase:
    def __init__(self, Token, url):

        git = Gitlab(url, Token)
        self.client = git
        self.token=Token

    # This method returns true if the client is valid 
    def isValidClient(self):

        user = self.client.get_user()
        try:
            login = user.login
            return True
        except:
            return False
        return True
    
    
class Singleton(type):
    _instances = {}
    def __call__(cls, *args, **kwargs):
        if cls not in cls._instances:
            cls._instances[cls] = super(Singleton, cls).__call__(*args, **kwargs)
        return cls._instances[cls]

class ClientGitlab(ClientGitlabBase, metaclass=Singleton):
    pass
