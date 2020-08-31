#! /usr/bin/env python3
# coding: utf-8

from .WorkerInterface import WorkerInterface
from .workers.WorkerIAM import WorkerIAM
from .workers.WorkerPlugins import WorkerPlugins

from pyclient.ClientGandalf import ClientGandalf
from typing import List
from threading import Thread

import json
import sys

class WorkerAws(WorkerInterface):
    config

    def __init__(self, version: int, commandes: List[str], config):
        super().__init__(version, commandes)

        self.config = config

    def Execute(self, clientGandalf: ClientGandalf, version: int):
        print("WorkerAws running")

        workerIAM = WorkerIAM(clientGandalf, version, config)
        workerPlugins = WorkerPlugins(clientGandalf, version, config)

        threadWorkerIAM = Thread(target=workerIAM.Run())
        threadWorkerPlugins = Thread(target=workerPlugins.Run())

        threadWorkerIAM.start()
        threadWorkerPlugins.start()

        threadWorkerIAM.join()
        threadWorkerPlugins.join()
    
    def reportError(self, uuid: str, command: str, error: str, timeout: str = "10000"):
        payload = (
            "{}: Error in command #{}\n{}".format(command, uuid, error)
        )
        self.clientGandalf.SendEvent(uuid, "FAIL", {timeout, payload})


if __name__ == "__main__":
    commands = list()
    version = int()
    
    config = json.loads(sys.stdin.read())

    workerAws = WorkerAws(version, commands, config)

    workerAws.Run()
