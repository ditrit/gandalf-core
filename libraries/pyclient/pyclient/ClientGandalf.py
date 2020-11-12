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

    identity: str
    clientConnections: List[str]
    clients: List[ClientGrpc]
    timeout: str
    clientIndex: int

    def __init__(self, identity: str, timeout: str, clientConnections: List[str]):
        self.identity = identity
        self.clientConnections = clientConnections
        self.clients = []
        self.clientIndex = 0

        if timeout == "":
            self.timeout = DEFAULT_TIMEOUT
        else:
            self.timeout = timeout

        for connection in self.clientConnections:
            self.clients.append(ClientGrpc(self.identity, connection))

    # Need to be checked
    def SendCommand(self, request: str, options: Options) -> CommandMessageUUID:
        notSend: bool
        requestSplit = request.split(".")
        timeout = options.timeout
        if timeout == "":
            timeout = self.timeout

        stay = True
        timeoutLoop = Queue(maxsize=0)
        Thread(target=lambda: time.sleep(1) or timeoutLoop.put(0)).start()

        while stay:
            commandMessageUUID = self.clients[self.getClientIndex(self.clients, True)].SendCommand(
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
            timeout = self.timeout

        stay = True
        timeoutLoop = Queue(maxsize=0)
        Thread(target=lambda: time.sleep(1) or timeoutLoop.put(0)).start()

        while stay:
            empty = self.clients[self.getClientIndex(self.clients, True)].SendEvent(
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
            timeout = self.timeout

        stay = True
        timeoutLoop = Queue(maxsize=0)
        Thread(target=lambda: time.sleep(1) or timeoutLoop.put(0)).start()

        while stay:
            empty = self.clients[self.getClientIndex(self.clients, True)].SendEvent(
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
        return self.clients[self.getClientIndex(self.clients, True)].SendCommandList(major, minor, commands)

    def SendStop(self, major: int, minor: int) -> Validate:
        return self.clients[self.getClientIndex(self.clients, True)].SendStop(major, minor)

    def WaitCommand(self, command, idIterator: str, version: int) -> CommandMessage:
        return self.clients[self.getClientIndex(self.clients, False)].WaitCommand(command, idIterator, version)

    def WaitEvent(self, topic, event, idIterator: str) -> EventMessage:
        return self.clients[self.getClientIndex(self.clients, False)].WaitEvent(topic, event, "", idIterator)

    def WaitReplyByEvent(self, topic, event, referenceUUID, idIterator: str) -> EventMessage:
        return self.clients[self.getClientIndex(self.clients, False)].WaitEvent(topic, event, referenceUUID, idIterator)

    def WaitTopic(self, topic, idIterator: str) -> EventMessage:
        return self.clients[self.getClientIndex(self.clients, False)].WaitTopic(topic, "", idIterator)

    def WaitReplyByTopic(self, topic, referenceUUID, idIterator: str) -> EventMessage:
        return self.clients[self.getClientIndex(self.clients, False)].WaitTopic(topic, referenceUUID, idIterator)

    def WaitAllReplyByTopic(self, topic, referenceUUID, idIterator, version: str) -> List[EventMessage]:
        client = self.clients[self.getClientIndex(self.clients, False)]
        eventMessages = []
        loop = True

        while loop:
            message = client.WaitTopic(topic, referenceUUID, idIterator)
            eventMessages.append(message)

            if message.Event == "SUCCES":
                loop = False

        return eventMessages

    def CreateIteratorCommand(self) -> str:
        return self.clients[self.getClientIndex(self.clients, False)].CreateIteratorCommand().Id

    def CreateIteratorEvent(self) -> str:
        return self.clients[self.getClientIndex(self.clients, False)].CreateIteratorEvent().Id

    def getClientIndex(self, conns: List[ClientGrpc], updateIndex: bool) -> int:
        aux = self.clientIndex

        if updateIndex:
            self.clientIndex += 1
            if self.clientIndex >= len(conns):
                self.clientIndex = 0

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
