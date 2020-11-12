#! /usr/bin/env python3
# coding: utf-8

from typing import Dict, List
from threading import Thread
import time

from pyclient.ClientGandalf import ClientGandalf
from pyclient.models.Options import Options
from pyclient.grpc.connectorCommand_pb2 import CommandMessage
from pyclient.grpc.connectorEvent_pb2 import EventMessage

from .functions.Start import Start
from .functions.SendCommands import SendCommands


class Worker:

    major: int
    minor: int
    commandes: List[str]
    clientGandalf: ClientGandalf
    CommandsFuncs: Dict
    EventsFunc: Dict
    WorkerState: List[WorkerState]
    OngoingTreatments: List[OngoingTreatments]

    def Start(self, clientGandalf: ClientGandalf):
        pass

    def Stop(self, clientGandalf: ClientGandalf, major: int, minor: int, workerState: List[WorkerState]):
        pass

    def SendCommands(self, clientGandalf: ClientGandalf, major: int, minor: int, commands: List[str]):
        pass

    def __init__(self, major: int, minor: int, commandes: List[str]):
        self.major = major
        self.minor = minor
        self.commandes = commandes

        self.Start = Start
        self.SendCommands = SendCommands

    def Run(self):
        self.clientGandalf = self.Start()

        keys = self.CommandsFuncs.values()

        valid = self.SendCommands(
            self.clientGandalf, self.major, self.minor, keys)

        joinList: List[Thread] = []

        if valid:
            joinList.append(Thread(target=self.Stop(
                self.clientGandalf, self.major, self.minor, self.WorkerState)))
            joinList[len(joinList)-1].start()

            for key, function in self.CommandsFuncs:
                id = self.clientGandalf.CreateIteratorCommand()

                joinList.append(
                    Thread(target=self.WaitCommands(id, key, function)))
                joinList[len(joinList)-1].start()

            for key, function in self.EventsFunc:
                id = self.clientGandalf.CreateIteratorEvent()

                joinList.append(
                    Thread(target=self.WaitEvents(id, key, function)))
                joinList[len(joinList)-1].start()

            for wstate in self.WorkerState:
                if wstate.GetState() == 0:
                    pass

            for ontreatment in self.OngoingTreatments:
                if ontreatment.GetIndex() > 0:
                    time.Sleep(2)
        else:
            # SEND EVENT INVALID CONFIGURATION
            pass

        for threadWait in joinList:
            threadWait.join()

    def WaitCommands(self, id, commandName: str, function: function):
        joinList: List[Thread] = []

        for wstate in self.WorkerState:
            if wstate.GetState() == 0:
                print("wait {}".format(commandName))
                command = self.clientGandalf.WaitCommand(
                    commandName, id, self.major)
                print("command")
                print(command)

                joinList.append(
                    Thread(target=self.executeCommands(command, function)))
                joinList[len(joinList)-1].start()

        for ontreatment in self.OngoingTreatments:
            if ontreatment.GetIndex() > 0:
                time.Sleep(2)

        print("END WAIT")
        for threadWait in joinList:
            threadWait.join()

    def executeCommands(self, command: CommandMessage, function: function):
        print("execute")
        self.OngoingTreatments.IncrementOngoingTreatments()
        result = function(self.clientGandalf, self.major, command)

        if result == 0:
            self.clientGandalf.SendReply(command.Command, "SUCCES", command.UUID, Options("", ""))
        else:
            self.clientGandalf.SendReply(command.Command, "FAIL", command.UUID, Options("", ""))

        self.OngoingTreatments.DecrementOngoingTreatments()
    
    def waitEvents(self, id: str, topicEvent: TopicEvent, function: function):
        joinList: List[Thread] = []

        for wstate in self.WorkerState:
            if wstate.GetState() == 0:
                event = self.clientGandalf.WaitEvent(topicEvent.Topic, topicEvent.Event, id)

                joinList.append(Thread(target=self.executeEvents(event, function)))
                joinList[len(joinList)-1].start()

        for ontreatment in self.OngoingTreatments:
            if ontreatment.GetIndex() > 0:
                time.Sleep(2)

        print("END WAIT")
        for threadWait in joinList:
            threadWait.join()
    
    def executeEvents(self, event: EventMessage, function: function):
        self.OngoingTreatments.IncrementOngoingTreatments()
        result = function(self.clientGandalf, self.major, event)

        if result == 0:
            self.clientGandalf.SendReply(event.Event, "SUCCES", event.UUID, Options("", ""))
        else:
            self.clientGandalf.SendReply(event.Event, "FAIL", event.UUID, Options("", ""))

        self.OngoingTreatments.DecrementOngoingTreatments()