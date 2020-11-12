#! /usr/bin/env python3
# coding: utf-8

from typing import List, Set
from pyclient.ClientGandalf import ClientGandalf


def Stop(clientGandalf: ClientGandalf, major: int, minor: int, workerState: WorkerState):
    validate = clientGandalf.SendStop(major, minor)

    if validate.Valid:
        workerState.SetStoppingWorkerState()
