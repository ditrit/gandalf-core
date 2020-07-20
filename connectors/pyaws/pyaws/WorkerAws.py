#! /usr/bin/env python3
# coding: utf-8

from pyworker.WorkerWorkflow import WorkerWorkflow
from pyclient.ClientGandalf import ClientGandalf

import json
import sys

class WorkerAws(WorkerWorkflow):

    def __init__(self, version, commandes):
        super().__init__(version, commandes)

    def Upload(self, clientGandalf: ClientGandalf, version: int):
        # run stuff here
        print("WorkerAws running")

        # workerA = WorkerA(clientGandalf, version)
        # threadWorkerA = Thread(target=workerA.Run())
        # threadWorkerA.start()

        # workerB = WorkerB(clientGandalf, version)
        # threadWorkerB = Thread(target=workerB.Run())
        # threadWorkerB.start()

        # threadWorkerA.join()
        # threadWorkerB.join()


if __name__ == "__main__":
    commands = list()
    version = int()
    
    config = json.loads(sys.stdin.readlines())

    workerAws = WorkerAws(version, commands)

    workerAws.Run()
