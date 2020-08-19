#! /usr/bin/env python3
# coding: utf-8

from .WorkerInterface import WorkerInterface
from .workers.WorkerIAM import WorkerIAM

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
        threadWorkerIAM = Thread(target=workerIAM.Run())
        threadWorkerIAM.start()

        threadWorkerIAM.join()


if __name__ == "__main__":
    commands = list()
    version = int()
    
    config = json.loads(sys.stdin.read())

    workerAws = WorkerAws(version, commands, config)

    workerAws.Run()
