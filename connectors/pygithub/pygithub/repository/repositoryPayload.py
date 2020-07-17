
class CreateRepositoryPayload:
    def __init__(self,  payload):
        self.name = payload['name']
        self.private = bool(payload['private'])

class AddUserToCollaboratorsPayload:
    def __init__(self, payload):

        #TODO 

        pass

class RemoveUserFromCollaboratorsPayload:
    def __init__(self, payload):

        #TODO

        pass
