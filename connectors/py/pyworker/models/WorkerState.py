
from threading import Lock


class State:
    ONGOING = 0
    STOPPING = 1


class WorkerState(Lock):
    state: State = State.ONGOING

    def getState(self) -> State:
        """
        Return the state
        """
        return self.state

    def setOngoingWorkerState(self):
        with self:
            self.state = State.ONGOING

    def setStoppingWorkerState(self):
        with self:
            self.state = State.STOPPING
