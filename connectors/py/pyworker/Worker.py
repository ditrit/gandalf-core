#! /usr/bin/env python3
# coding: utf-8

from typing import Dict, List, Callable
from threading import Thread
import time

from pyclient.ClientGandalf import ClientGandalf
from pyclient.models.Options import Options
from pyclient.grpc.connectorCommand_pb2 import CommandMessage
from pyclient.grpc.connectorEvent_pb2 import EventMessage

from .functions.Start import Start
from .functions.SendCommands import SendCommands
from .functions.Stop import Stop
from .models.workerState import WorkerState
from .models.ongoingTreatments import OngoingTreatments
from .models.topicEvent import TopicEvent


class Worker:

    major: int
    minor: int
    clientGandalf: ClientGandalf
    CommandsFuncs: Dict[Callable[[ClientGandalf, int, CommandMessage], int]]
    EventsFunc: Dict[Callable[[ClientGandalf, int, EventMessage], int]]
    WorkerState: List[WorkerState]
    OngoingTreatments: List[OngoingTreatments]

    def Start(self, clientGandalf: ClientGandalf):
        pass

    def Stop(self, clientGandalf: ClientGandalf, major: int, minor: int, workerState: List[WorkerState]):
        pass

    def SendCommands(self, clientGandalf: ClientGandalf, major: int, minor: int, commands: List[str]):
        pass

    def __init__(self, major: int, minor: int):
        self.major = major
        self.minor = minor
        self.CommandsFuncs = {}
        self.EventsFunc = {}
        self.OngoingTreatments = OngoingTreatments()
        self.WorkerState = WorkerState()

        self.Start = Start
        self.Stop = Stop
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
                    Thread(target=self.waitCommands(id, key, function)))
                joinList[len(joinList)-1].start()

            for key, function in self.EventsFunc:
                id = self.clientGandalf.CreateIteratorEvent()

                joinList.append(
                    Thread(target=self.waitEvents(id, key, function)))
                joinList[len(joinList)-1].start()

            for wstate in self.WorkerState:
                if wstate.GetState() == 0:
                    pass

            for ontreatment in self.OngoingTreatments:
                if ontreatment.GetIndex() > 0:
                    time.Sleep(2)
        else:
            # SEND EVENT INVALID CONFIGURATION
            # self.clientGandalf.SendReply(
            #     event.Event, "FAIL", event.UUID, Options("", ""))
            pass

        for threadWait in joinList:
            threadWait.join()

    def waitCommands(self, id, commandName: str, function: Callable[[ClientGandalf, int, CommandMessage], int]):
        print("[{}](waitCommands) Start wait".format(id))

        joinList: List[Thread] = []

        for wstate in self.WorkerState:
            if wstate.GetState() == 0:
                print("[{}](waitCommands) Wait for {}".format(id, commandName))
                command = self.clientGandalf.WaitCommand(
                    commandName, id, self.major)
                print("[{}](waitCommands) command :\n{}".format(id, command))

                joinList.append(
                    Thread(target=self.executeCommands(command, function)))
                joinList[len(joinList)-1].start()

        for ontreatment in self.OngoingTreatments:
            if ontreatment.GetIndex() > 0:
                time.Sleep(2)

        print("[{}](waitCommands) Wait for tasks to finish".format(id))
        for threadWait in joinList:
            threadWait.join()
        print("[{}](waitCommands) End Wait".format(id))

    def executeCommands(self, command: CommandMessage, function: Callable[[ClientGandalf, int, CommandMessage], int]):
        print("[{}](executeCommands) Execute command : {}".format(
            id, command.Command))
        self.OngoingTreatments.IncrementOngoingTreatments()
        result = function(self.clientGandalf, self.major, command)

        if result == 0:
            self.clientGandalf.SendReply(
                command.Command, "SUCCES", command.UUID, Options("", ""))
        else:
            self.clientGandalf.SendReply(
                command.Command, "FAIL", command.UUID, Options("", ""))

        self.OngoingTreatments.DecrementOngoingTreatments()

    def waitEvents(self, id: str, topicEvent: TopicEvent, function: Callable[[ClientGandalf, int, EventMessage], int]):
        print("[{}](waitEvents) Start wait".format(id))

        joinList: List[Thread] = []

        for wstate in self.WorkerState:
            if wstate.GetState() == 0:
                print("[{}](waitEvents) Wait for {}".format(
                    id, topicEvent.Event))
                event = self.clientGandalf.WaitEvent(
                    topicEvent.Topic, topicEvent.Event, id)
                print("[{}](waitEvents) event :\n{}".format(id, event))

                joinList.append(
                    Thread(target=self.executeEvents(event, function)))
                joinList[len(joinList)-1].start()

        for ontreatment in self.OngoingTreatments:
            if ontreatment.GetIndex() > 0:
                time.Sleep(2)

        print("[{}](waitEvents) Wait for tasks to finish".format(id))
        for threadWait in joinList:
            threadWait.join()
        print("[{}](waitEvents) End Wait".format(id))

    def executeEvents(self, event: EventMessage, function: Callable[[ClientGandalf, int, EventMessage], int]):
        print("[{}](executeEvents) Execute event : {}".format(id, event.Event))
        self.OngoingTreatments.IncrementOngoingTreatments()
        result = function(self.clientGandalf, self.major, event)

        if result == 0:
            self.clientGandalf.SendReply(
                event.Event, "SUCCES", event.UUID, Options("", ""))
        else:
            self.clientGandalf.SendReply(
                event.Event, "FAIL", event.UUID, Options("", ""))

        self.OngoingTreatments.DecrementOngoingTreatments()
