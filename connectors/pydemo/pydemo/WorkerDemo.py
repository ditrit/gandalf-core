#! /usr/bin/env python3
# coding: utf-8

from .WorkerInterface import WorkerInterface
from .workers.WorkerTest import WorkerTest

import json, sys
from typing import List, Dict
from threading import Thread

class WorkerDemo(WorkerInterface):
    config: Dict

    def __init__(self, version: int, commandes: List[str], config):
        super().__init__(version, commandes)

        self.config = config
    
    def Execute(self, clientGandalf, version):
        workerTest = WorkerTest(clientGandalf, version, self.config)
        threadWorkerTest = Thread(target=workerTest.Run())
        threadWorkerTest.start()
        threadWorkerTest.join()

if __name__ == "__main__":
    commands = list()
    version = int()

    config = json.loads(sys.stdin.read())

    workerDemo = WorkerDemo(version, commands, config)

    workerDemo.Run()