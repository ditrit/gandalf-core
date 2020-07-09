#! /usr/bin/env python3
# coding: utf-8

from .Worker import Worker
from pyclient.ClientGandalf import ClientGandalf

class WorkerAws(Worker):

    def __init__(self, version, commandes):
        super().__init__(version, commandes)

    def Execute(self, clientGandalf: ClientGandalf, version: int):
        # run stuff here
        print("WorkerAws running")

        # workerA = WorkerA(clientGandalf, version)
        # threadWorkerA = Thread(target=workerA.Run())
        # threadWorkerA.start()

        # workerB = WorkerB(clientGandalf, version)
        # threadWorkerB = Thread(target=workerB.Run())
        # threadWorkerB.start()

        # threadWorkerB.join()


if __name__ == "__main__":
    commands = ["CommandA", "CommandB"]
    version = 2

    workerAws = WorkerAws(version, commands)

    workerAws.Run()
