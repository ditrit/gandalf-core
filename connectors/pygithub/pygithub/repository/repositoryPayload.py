
class CreateRepositoryPayload:
    def __init__(self,  payload):
        self.name = payload['name']
        self.private = bool(payload['private'])
