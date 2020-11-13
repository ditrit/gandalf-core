#! /usr/bin/env python3
# coding: utf-8

from typing import Dict, List, Callable, Set
from threading import Thread
import time

from pyclient.ClientGandalf import ClientGandalf
from pyclient.models.Options import Options
from pyclient.grpc.connectorCommand_pb2 import CommandMessage
from pyclient.grpc.connectorEvent_pb2 import EventMessage

from .functions.Start import Start
from .functions.SendCommands import SendCommands
from .functions.Stop import Stop
from .functions.waitStop import waitStop

from .models.WorkerState import WorkerState
from .models.OngoingTreatments import OngoingTreatments
from .models.TopicEvent import TopicEvent


class Worker:

    major: int
    minor: int
    clientGandalf: ClientGandalf
    CommandsFuncs: Dict[Callable[[ClientGandalf, int, CommandMessage], int]]
    EventsFuncs: Dict[Callable[[ClientGandalf, int, EventMessage], int]]
    WorkerState: WorkerState
    OngoingTreatments: OngoingTreatments

    def Start(self, clientGandalf: ClientGandalf):
        pass

    def Stop(self, clientGandalf: ClientGandalf, major: int, minor: int, workerState: WorkerState):
        pass

    def waitStop(self, clientGandalf: ClientGandalf, major: int, minor: int, workerState: WorkerState, ongoingTreatment: OngoingTreatments):
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
        self.waitStop = waitStop

    def Run(self):
        joinList: List[Thread] = []

        # [RUN] Step 1 : Exec Start function
        self.clientGandalf = self.Start()

        # [RUN] Step 2 : Send keys of CommandsFuncs to router
        keys = self.CommandsFuncs.keys()
        valid = self.SendCommands(
            self.clientGandalf,
            self.major,
            self.minor,
            keys
        )
        if valid:
            # [RUN] Step 3 : Set state as "ongoing"
            self.WorkerState.setOngoingWorkerState()
            for key, function in self.CommandsFuncs:
                # [RUN] Step 4 : Run waitCommand
                id = self.clientGandalf.CreateIteratorCommand()

                joinList.append(
                    Thread(target=self.waitCommands(id, key, function))
                )
                joinList[len(joinList)-1].start()
                joinList.append(
                    Thread(target=self.waitStop(
                        self.clientGandalf,
                        self.major,
                        self.minor,
                        self.WorkerState,
                        self.OngoingTreatments
                    ))
                )
                joinList[len(joinList)-1].start()

            for key, function in self.EventsFunc:
                # [RUN] Step 4bis : Run waitEvent
                id = self.clientGandalf.CreateIteratorEvent()

                joinList.append(
                    Thread(target=self.waitEvents(id, key, function))
                )
                joinList[len(joinList)-1].start()
                joinList.append(
                    Thread(target=self.waitStop(
                        self.clientGandalf,
                        self.major,
                        self.minor,
                        self.WorkerState,
                        self.OngoingTreatments
                    ))
                )
                joinList[len(joinList)-1].start()
        else:
            # SEND EVENT INVALID CONFIGURATION
            # self.clientGandalf.SendReply(
            #     event.Event, "FAIL", event.UUID, Options("", ""))
            pass

        for threadWait in joinList:
            threadWait.join()

    def RegisterCommandsFuncs(self, command: str, function: Callable[[ClientGandalf, int, CommandMessage], int]):
        print("[{}](RegisterCommandsFuncs) REGISTER".format(command))
        self.CommandsFuncs[command] = function

    def RegisterEventsFuncs(self, topicevent: TopicEvent, function: Callable[[ClientGandalf, int, EventMessage], int]):
        print("[{}](RegisterEventsFuncs) REGISTER".format(topicevent.topic))
        self.EventsFuncs[topicevent] = function

    def waitCommands(self, id, commandName: str, function: Callable[[ClientGandalf, int, CommandMessage], int]):
        print("[{}](waitCommands) Start wait".format(id))

        joinList: List[Thread] = []

        while self.WorkerState.getState() == 0:
            print("[{}](waitCommands) Wait for {}".format(id, commandName))
            command = self.clientGandalf.WaitCommand(
                commandName,
                id,
                self.major
            )
            print("[{}](waitCommands) command :\n{}".format(id, command))

            joinList.append(
                Thread(target=self.executeCommands(command, function))
            )
            joinList[len(joinList)-1].start()

        while self.OngoingTreatments.getIndex() > 0:
            time.sleep(2)

        print("[{}](waitCommands) Wait for tasks to finish".format(id))
        for threadWait in joinList:
            threadWait.join()
        print("[{}](waitCommands) End Wait".format(id))

    def executeCommands(self, command: CommandMessage, function: Callable[[ClientGandalf, int, CommandMessage], int]):
        print(
            "[{}](executeCommands) Execute command : {}".format(
                id,
                command.Command
            )
        )
        self.OngoingTreatments.Increment()
        result = function(self.clientGandalf, self.major, command)

        if result == 0:
            self.clientGandalf.SendReply(
                command.Command,
                "SUCCES",
                command.UUID,
                Options("", "")
            )
        else:
            self.clientGandalf.SendReply(
                command.Command,
                "FAIL",
                command.UUID,
                Options("", "")
            )

        self.OngoingTreatments.Decrement()

    def waitEvents(self, id: str, topicEvent: TopicEvent, function: Callable[[ClientGandalf, int, EventMessage], int]):
        print("[{}](waitEvents) Start wait".format(id))

        joinList: List[Thread] = []

        while self.WorkerState.getState() == 0:
            print(
                "[{}](waitEvents) Wait for {}".format(
                    id,
                    topicEvent.event
                )
            )
            event = self.clientGandalf.WaitEvent(
                topicEvent.topic,
                topicEvent.event,
                id
            )
            print(
                "[{}](waitEvents) event :\n{}".format(
                    id,
                    event
                )
            )

            joinList.append(
                Thread(target=self.executeEvents(event, function))
            )
            joinList[len(joinList)-1].start()

        while self.OngoingTreatments.getIndex() > 0:
            time.sleep(2)

        print("[{}](waitEvents) Wait for tasks to finish".format(id))
        for threadWait in joinList:
            threadWait.join()
        print("[{}](waitEvents) End Wait".format(id))

    def executeEvents(self, event: EventMessage, function: Callable[[ClientGandalf, int, EventMessage], int]):
        print("[{}](executeEvents) Execute event : {}".format(id, event.Event))
        self.OngoingTreatments.Increment()
        result = function(self.clientGandalf, self.major, event)

        if result == 0:
            self.clientGandalf.SendReply(
                event.Event,
                "SUCCES",
                event.UUID,
                Options("", "")
            )
        else:
            self.clientGandalf.SendReply(
                event.Event,
                "FAIL",
                event.UUID,
                Options("", "")
            )

        self.OngoingTreatments.Decrement()
