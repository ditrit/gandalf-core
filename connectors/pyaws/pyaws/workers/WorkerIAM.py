#! /usr/bin/env python3
# coding: utf-8

from ..WorkerAws import WorkerAws
from ..AWS.IAM.IAM import IAM

from typing import List
import json

class WorkerIAM(WorkerAws):
    iamClient: IAM

    def __init__(self, version: int, commandes: List[str], config):
        super().__init__(version, commandes, config)

        self.iamClient = IAM(config["aws_region_name"], config["aws_access_key_id"], config["aws_secret_access_key"])

    def Execute(self, clientGandalf, version):
        #fetch the config or something here
        pass

    def CreateUser(self):
        id = self.clientGandalf.CreateIteratorCommand()
        print(id)

        command = self.clientGandalf.WaitCommand("CREATE_USER", id, self.version)
        print(command)

        payload = json.loads(command.Payload)


        response = self.iamClient.createUser(payload["name"])
        if response == None:
            self.clientGandalf.SendEvent(command.UUID, "FAIL",{"10000", "User not created !"})
        else:
            self.clientGandalf.SendEvent(command.UUID, "SUCCES", {"10000", "User created !"})
    
    def UpdateUser(self):
        id = self.clientGandalf.CreateIteratorCommand()
        print(id)

        command = self.clientGandalf.WaitCommand("UPDATE_USER", id, self.version)
        print(command)

        payload = json.loads(command.Payload)

        response = self.iamClient.updateUser(userName=payload["name"], newUserName=payload["newName"])
        if response == None:
            self.clientGandalf.SendEvent(command.UUID, "FAIL",{"10000", "User not updated !"})
        else:
            self.clientGandalf.SendEvent(command.UUID, "SUCCES", {"10000", "User updated !"})
    
    def DeleteUser(self):
        id = self.clientGandalf.CreateIteratorCommand()
        print(id)

        command = self.clientGandalf.WaitCommand("DELETE_USER", id, self.version)
        print(command)

        payload = json.loads(command.Payload)

        response = self.iamClient.deleteUser(userName=payload["name"])
        if response == None:
            self.clientGandalf.SendEvent(command.UUID, "FAIL",{"10000", "User not deleted !"})
        else:
            self.clientGandalf.SendEvent(command.UUID, "SUCCES", {"10000", "User deleted !"})

    def CreateGroup(self):
        id = self.clientGandalf.CreateIteratorCommand()
        print(id)

        command = self.clientGandalf.WaitCommand("CREATE_GROUP", id, self.version)
        print(command)

        payload = json.loads(command.Payload)


        # response = self.iamClient.createUser(payload["name"])
        # if response == None:
        #     self.clientGandalf.SendEvent(command.UUID, "FAIL",{"10000", "User not created !"})
        # else:
        #     self.clientGandalf.SendEvent(command.UUID, "SUCCES", {"10000", "User created !"})

    def UpdateGroup(self):
        id = self.clientGandalf.CreateIteratorCommand()
        print(id)

        command = self.clientGandalf.WaitCommand("UPDATE_GROUP", id, self.version)
        print(command)

        payload = json.loads(command.Payload)


        # response = self.iamClient.createUser(payload["name"])
        # if response == None:
        #     self.clientGandalf.SendEvent(command.UUID, "FAIL",{"10000", "User not created !"})
        # else:
        #     self.clientGandalf.SendEvent(command.UUID, "SUCCES", {"10000", "User created !"})

    def DeleteGroup(self):
        id = self.clientGandalf.CreateIteratorCommand()
        print(id)

        command = self.clientGandalf.WaitCommand("DELETE_GROUP", id, self.version)
        print(command)

        payload = json.loads(command.Payload)


        # response = self.iamClient.createUser(payload["name"])
        # if response == None:
        #     self.clientGandalf.SendEvent(command.UUID, "FAIL",{"10000", "User not created !"})
        # else:
        #     self.clientGandalf.SendEvent(command.UUID, "SUCCES", {"10000", "User created !"})


    def Run(self):
        pass
