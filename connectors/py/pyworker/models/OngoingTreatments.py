
from threading import Lock

class OngoingTreatments(Lock):

    index: int = 0

    def getIndex(self) -> int:
        with self:
            index: int = self.index

        return index
    
    def Increment(self):
        with self:
            self.index += 1

    def Decrement(self):
        with self:
            self.index -= 1