#! /usr/bin/env python3
# coding: utf-8

import time
from queue import Queue
from threading import Thread
from typing import List

from .client.ClientGrpc import ClientGrpc
from .models.Options import Options

from .grpc.connector_pb2 import *
from .grpc.connectorCommand_pb2 import *
from .grpc.connectorEvent_pb2 import *

DEFAULT_TIMEOUT = "10000"

class ClientGandalf:

    Identity: str
    ClientConnections: List[str]
    Clients: List[ClientGrpc]
    Timeout: str
    ClientIndex: int

    def __init__(self, identity: str, timeout: str, clientConnections: List[str]):
        self.Identity = identity
        self.ClientConnections = clientConnections
        self.Clients = []
        self.ClientIndex = 0

        if timeout == "":
            self.Timeout = DEFAULT_TIMEOUT
        else:
            self.Timeout = timeout

        for connection in self.ClientConnections:
            self.Clients.append(ClientGrpc(self.Identity, connection))

    # Need to be checked
    def SendCommand(self, request: str, options: Options) -> CommandMessageUUID:
        notSend: bool
        requestSplit = request.split(".")
        timeout = options.timeout
        if timeout == "":
            timeout = self.Timeout

        stay = True
        timeoutLoop = Queue(maxsize=0)
        Thread(target=lambda: time.sleep(1) or timeoutLoop.put(0)).start()

        while stay:
            commandMessageUUID = self.Clients[self.getClientIndex(self.Clients, True)].SendCommand(
                requestSplit[0], requestSplit[1], timeout, options.payload)
            if commandMessageUUID != None:
                notSend = False
                break

            for which, _ in select(timeoutLoop):
                if which is timeoutLoop:
                    stay = False
                    notSend = True
                    return

        if notSend:
            return None

        return commandMessageUUID

    def SendEvent(self, topic, event: str, options: Options) -> Empty:
        notSend: bool
        timeout = options.timeout
        if timeout == "":
            timeout = self.Timeout

        stay = True
        timeoutLoop = Queue(maxsize=0)
        Thread(target=lambda: time.sleep(1) or timeoutLoop.put(0)).start()

        while stay:
            empty = self.Clients[self.getClientIndex(self.Clients, True)].SendEvent(
                topic, event, "", timeout, options.payload)
            if empty != None:
                notSend = False
                break

            for which, _ in select(timeoutLoop):
                if which is timeoutLoop:
                    stay = False
                    notSend = True
                    return

        if notSend:
            return None

        return empty

    def SendReply(self, topic, event, referenceUUID: str, options: Options) -> Empty:
        notSend: bool
        timeout = options.timeout

        if timeout == "":
            timeout = self.Timeout

        stay = True
        timeoutLoop = Queue(maxsize=0)
        Thread(target=lambda: time.sleep(1) or timeoutLoop.put(0)).start()

        while stay:
            empty = self.Clients[self.getClientIndex(self.Clients, True)].SendEvent(
                topic, event, referenceUUID, timeout, options.payload)
            if empty != None:
                notSend = False
                break

            for which, _ in select(timeoutLoop):
                if which is timeoutLoop:
                    stay = False
                    notSend = True
                    return

        if notSend:
            return None

        return empty

    def SendCommandList(self, major: int, minor: int, commands: List[str]) -> Validate:
        return self.Clients[self.getClientIndex(self.Clients, True)].SendCommandList(major, minor, commands)

    def SendStop(self, major: int, minor: int) -> Validate:
        return self.Clients[self.getClientIndex(self.Clients, True)].SendStop(major, minor)

    def WaitCommand(self, command, idIterator: str, version: int) -> CommandMessage:
        return self.Clients[self.getClientIndex(self.Clients, False)].WaitCommand(command, idIterator, version)

    def WaitEvent(self, topic, event, idIterator: str) -> EventMessage:
        return self.Clients[self.getClientIndex(self.Clients, False)].WaitEvent(topic, event, "", idIterator)

    def WaitReplyByEvent(self, topic, event, referenceUUID, idIterator: str) -> EventMessage:
        return self.Clients[self.getClientIndex(self.Clients, False)].WaitEvent(topic, event, referenceUUID, idIterator)

    def WaitTopic(self, topic, idIterator: str) -> EventMessage:
        return self.Clients[self.getClientIndex(self.Clients, False)].WaitTopic(topic, "", idIterator)

    def WaitReplyByTopic(self, topic, referenceUUID, idIterator: str) -> EventMessage:
        return self.Clients[self.getClientIndex(self.Clients, False)].WaitTopic(topic, referenceUUID, idIterator)

    def WaitAllReplyByTopic(self, topic, referenceUUID, idIterator, version: str) -> List[EventMessage]:
        client = self.Clients[self.getClientIndex(self.Clients, False)]
        eventMessages = []
        loop = True

        while loop:
            message = client.WaitTopic(topic, referenceUUID, idIterator)
            eventMessages.append(message)

            if message.Event == "SUCCES":
                loop = False

        return eventMessages

    def CreateIteratorCommand(self) -> str:
        return self.Clients[self.getClientIndex(self.Clients, False)].CreateIteratorCommand().Id

    def CreateIteratorEvent(self) -> str:
        return self.Clients[self.getClientIndex(self.Clients, False)].CreateIteratorEvent().Id

    def getClientIndex(self, conns: List[ClientGrpc], updateIndex: bool) -> int:
        aux = self.ClientIndex

        if updateIndex:
            self.ClientIndex += 1
            if self.ClientIndex >= len(conns):
                self.ClientIndex = 0

        return aux


def select(*queues):
    combined = Queue(maxsize=0)

    def listen_and_forward(queue):
        while True:
            combined.put((queue, queue.get()))
    for queue in queues:
        t = Thread(target=listen_and_forward, args=(queue,))
        t.daemon = True
        t.start()
    while True:
        yield combined.get()
