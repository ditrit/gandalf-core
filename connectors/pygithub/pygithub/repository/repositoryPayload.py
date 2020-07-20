
class CreateRepositoryPayload:
    def __init__(self,  payload):
        self.name = payload['name']
        self.private = bool(payload['private'])

class AddUserToCollaboratorsPayload:
    def __init__(self, payload):
        self.name  = payload['name']
        self.username = payload['username']
        self.permission = payload['permission']



class RemoveUserFromCollaboratorsPayload:
    def __init__(self, payload):
        self.name  = payload['name']
        self.username = payload['username']

        pass
