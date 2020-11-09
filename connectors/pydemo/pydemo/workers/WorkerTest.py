#! /usr/bin/env python3
# coding: utf-8

from ..WorkerDemo import WorkerDemo

import json

class WorkerTest(WorkerDemo):

    def __init__(self, version, commandes, config):
        super().__init__(version, commandes, config)

    def Execute(self, clientGandalf, version):
        return super().Execute(clientGandalf, version)

    def runTest(self):
        id = self.clientGandalf.CreateIteratorCommand()
        print(id)

        command = self.clientGandalf.WaitCommand("RUN_TEST_1", id, self.version)
        print(command)

        payload = json.loads(command.Payload)

        print(payload)

        self.clientGandalf.SendEvent(command.UUID, "SUCCES", {"10000", "RUN_TEST_1 works !"})