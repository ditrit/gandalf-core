#! /usr/bin/env python3
# coding: utf-8

from ..WorkerAws import WorkerAws

import json

class WorkerIAM(WorkerAws):
    def __init__(self, version, commandes):
        super().__init__(version, commandes)

    def Execute(self, clientGandalf, version):
        #fetch the config or something here
        pass

    def CreateUser(self):
        id = self.clientGandalf.CreateIteratorCommand()
        print(id)

        command = self.clientGandalf.WaitCommand("CREATE_USER", id, self.version)
        print(command)

        payload = json.loads(command.Payload)
        


    def Run(self):
        pass
