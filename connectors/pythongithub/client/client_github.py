from github import Github
import logging

class ClientGithub:
    def __init__(self, Token, clientGithub):

        git = Github(Token)

        if not self.isValidClient(git):
            logging.ERROR("Failed to create client")

        self.client = git


    def __init__(self, user, password, clientGithub):

        git = Github(user, password)

        if not self.isValidClient(git):
            logging.ERROR("Failed to create client")

        self.client = git

    # This method returns true if the client is valid 
    def isValidClient(self, gitclient):
        #TODO
        return True
