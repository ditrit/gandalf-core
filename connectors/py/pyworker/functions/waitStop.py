#! /usr/bin/env python3
# coding: utf-8

from typing import List, Set
import time
from pyclient.ClientGandalf import ClientGandalf
from ..models.workerState import WorkerState
from ..models.ongoingTreatments import OngoingTreatments

def waitStop(self, clientGandalf: ClientGandalf, major: int, minor: int, workerState: WorkerState, ongoingTreatment: OngoingTreatments):        
    # [WAIT_STOP] Step 1 & 2
    self.Stop(clientGandalf, major, minor, workerState)

    # [WAIT_STOP] Step 3
    while workerState.GetState() == 0 or ongoingTreatment.GetIndex() != 0 :
        time.Sleep(2)