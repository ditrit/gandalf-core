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

        if self.iamClient.createUser(payload["name"]):
            self.clientGandalf.SendEvent(command.UUID, "SUCCES", {"10000", "User created !"})
        else:
            self.clientGandalf.SendEvent(command.UUID, "FAIL",{"10000", "User not created !"})


    def Run(self):
        pass
