#! /usr/bin/env python3
# coding: utf-8

from ..WorkerAws import WorkerAws
from ..AWS.IAM.IAM import IAM
from botocore.exceptions import ClientError

from typing import List
from threading import Thread
import json

# TODO : Implement more detailed return types including the API errors returned, and stronger payload checks
#        for known mandatory parameters


class WorkerIAM(WorkerAws):
    iamClient: IAM

    def __init__(self, version: int, commandes: List[str], config):
        super().__init__(version, commandes, config)

        self.iamClient = IAM(
            config["aws_region_name"], config["aws_access_key_id"], config["aws_secret_access_key"])

    def Execute(self, clientGandalf, version):
        # fetch the config or something here
        pass

    def reportClientError(self, uuid: str, command: str, error: ClientError, timeout: str = "10000"):
        payload = (
            "{}: An AWS error occurenced while executing command #{}\n".format(command, uuid) +
            "Code: {}, message: \n{}\n".format(error['Error']['Code'], error['Error']['Message']) +
            "AWS Metadata :\nRequest ID:{}, Host ID:{}, HTTP status: {}".format(error['ResponseMetadata']['RequestId'], 
                                                                                error['ResponseMetadata']['HostId'],
                                                                                error['ResponseMetadata']['HTTPStatusCode'])
        )
        self.clientGandalf.SendEvent(uuid, "FAIL", {timeout, payload})

    # TODO : Create user access key and return it along with the created user
    def CreateUser(self):
        id = self.clientGandalf.CreateIteratorCommand()
        print(id)

        command = self.clientGandalf.WaitCommand(
            "CREATE_USER", id, self.version)
        print(command)

        payload = json.loads(command.Payload)

        userName = payload['name'] if 'name' in payload else None
        path = payload['path'] if 'path' in payload else None
        tags = payload['tags'] if 'tags' in payload else []
        permissions = payload['permissions'] if 'permissions' in payload else None
        policies = payload['policies'] if 'policies' in payload else []
        groups = payload['groups'] if 'groups' in payload else []

        if not userName:
            self.reportError(uuid=command.UUID, command="CREATE_USER", error="User name must be provided")

        try:
            response = self.iamClient.createUser(
                userName=userName, groups=groups, policies=policies, permissions=permissions, tags=tags, path=path)           
        except ClientError as err:
            self.reportClientError(uuid=command.UUID, command="CREATE_USER", error=err)
        else:
            self.clientGandalf.SendEvent(command.UUID, "SUCCES", {"10000", "User created !"})


    def UpdateUser(self):
        id = self.clientGandalf.CreateIteratorCommand()
        print(id)

        command = self.clientGandalf.WaitCommand(
            "UPDATE_USER", id, self.version)
        print(command)

        payload = json.loads(command.Payload)

        userName = payload['name'] if 'name' in payload else None
        newUserName = payload['newName'] if 'newName' in payload else None
        newPath = payload['newPath'] if 'newPath' in payload else None
        policies = payload['policies'] if 'policies' in payload else []
        groups = payload['groups'] if 'groups' in payload else []

        if not userName:
            self.reportError(uuid=command.UUID, command="UPDATE_USER", error="User name must be provided")
        
        try:
            response = self.iamClient.updateUser(
                userName=userName, newUserName=newUserName, newPath=newPath, groups=groups, policies=policies)
        except ClientError as err:
            self.reportClientError(uuid=command.UUID, command="UPDATE_USER", error=err)
        else:
            self.clientGandalf.SendEvent(command.UUID, "SUCCES", {"10000", "User updated !"})


    def DeleteUser(self):
        id = self.clientGandalf.CreateIteratorCommand()
        print(id)

        command = self.clientGandalf.WaitCommand(
            "DELETE_USER", id, self.version)
        print(command)

        payload = json.loads(command.Payload)

        userName = payload['name'] if 'name' in payload else None

        if not userName:
            self.reportError(uuid=command.UUID, command="DELETE_USER", error="User name must be provided")

        try:
            response = self.iamClient.deleteUser(userName=userName)
        except ClientError as err:
            self.reportClientError(uuid=command.UUID, command="DELETE_USER", error=err)
        else:
            self.clientGandalf.SendEvent(command.UUID, "SUCCES", {
                                         "10000", "User deleted !"})


    # TODO : Check exsiting policies and create if necesary
    def CreateGroup(self):
        id = self.clientGandalf.CreateIteratorCommand()
        print(id)

        command = self.clientGandalf.WaitCommand(
            "CREATE_GROUP", id, self.version)
        print(command)

        payload = json.loads(command.Payload)

        groupName = payload['name'] if 'name' in payload else None
        path = payload['path'] if 'path' in payload else None
        policies = payload['policies'] if 'policies' in payload else []

        if not groupName:
            self.reportError(uuid=command.UUID, command="CREATE_GROUP", error="Group name must be provided")

        try:
            response = self.iamClient.createGroup(groupName=groupName, policies=policies, path=path)
        except ClientError as err:
            self.reportClientError(uuid=command.UUID, command="CREATE_GROUP", error=err)
        else:
            self.clientGandalf.SendEvent(command.UUID, "SUCCES", {
                                         "10000", "Group created !"})


    def UpdateGroup(self):
        id = self.clientGandalf.CreateIteratorCommand()
        print(id)

        command = self.clientGandalf.WaitCommand(
            "UPDATE_GROUP", id, self.version)
        print(command)

        payload = json.loads(command.Payload)
        name = payload['name'] if 'name' in payload else None
        newName = payload['newName'] if 'newName' in payload else None
        newPath = payload['newPath'] if 'newPath' in payload else None

        if not name:
            self.reportError(uuid=command.UUID, command="UPDATE_GROUP", error="Group name must be provided")
            
        try:
            response = self.iamClient.updateGroup(groupName=name, newGroupName=newName, newPath=newPath)
        except ClientError as err:
            self.reportClientError(uuid=command.UUID, command="UPDATE_GROUP", error=err)
        else:
            self.clientGandalf.SendEvent(command.UUID, "SUCCES", {
                                         "10000", "Group updated !"})

    def DeleteGroup(self):
        id = self.clientGandalf.CreateIteratorCommand()
        print(id)

        command = self.clientGandalf.WaitCommand(
            "DELETE_GROUP", id, self.version)
        print(command)

        payload = json.loads(command.Payload)
        name = payload['name'] if 'name' in payload else None

        if not name:
            self.reportError(uuid=command.UUID, command="DELETE_GROUP", error="Group name must be provided")
            
        try:
            response = self.iamClient.deleteGroup(groupName=name)
        except ClientError as err:
            self.reportClientError(uuid=command.UUID, command="DELETE_GROUP", error=err)
        else:
            self.clientGandalf.SendEvent(command.UUID, "SUCCES", {
                                         "10000", "Group deleted !"})

    # NOTE : Access keys removed from here, not relevant to the larger-scale Gandalf thingy. They should be automatically
    #        managed by the worker.

    def Run(self):
        createUser = Thread(target=self.CreateUser())
        updateUser = Thread(target=self.UpdateUser())
        deleteUser = Thread(target=self.DeleteUser())

        createUser.start()
        updateUser.start()
        deleteUser.start()

        createUser.join()
        updateUser.join()
        deleteUser.join()

        pass
