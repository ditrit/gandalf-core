#! /usr/bin/env python3
# coding: utf-8

import uuid
import time

from grpc import RpcError

from ..grpc.connectorEvent_pb2 import *
from ..grpc.connectorEvent_pb2_grpc import *

from ..grpc.connector_pb2 import *
from ..ClientWarper import ClientWarper

class ClientEvent(ClientWarper):

    Client: ConnectorEventStub

    @property
    def ClientEventConnection(self) -> str:
        return self.ClientConnection

    @ClientEventConnection.setter
    def ClientEventConnection(self, value: str):
        self.ClientConnection = value

    def __init__(self, identity: str, clientEventConnection: str):
        super().__init__(identity, clientEventConnection)

        try:
            conn = grpc.insecure_channel(clientEventConnection)
            
            self.Client = ConnectorEventStub(conn)
            print("clientEvent connect : {}".format(clientEventConnection))
        except RpcError as err:
            print("clientEvent failed dial : {}".format(err))

    def SendEvent(self, topic, event, referenceUUID, timeout, payload: str) -> Empty:
        eventMessage = EventMessage()
        eventMessage.Topic = topic
        eventMessage.Timeout = timeout
        eventMessage.UUID = str(uuid.uuid4())
        eventMessage.Event = event
        eventMessage.Payload = payload
        eventMessage.ReferenceUUID = referenceUUID

        return self.Client.SendEventMessage(eventMessage)

    def WaitEvent(self, topic, event, referenceUUID, idIterator: str) -> EventMessage:
        eventMessageWait = EventMessageWait()
        eventMessageWait.WorkerSource = self.Identity
        eventMessageWait.Topic = topic
        eventMessageWait.Event = event
        eventMessageWait.IteratorId = idIterator
        eventMessageWait.ReferenceUUID = referenceUUID
        eventMessage = self.Client.WaitEventMessage(eventMessageWait)

        while eventMessage == None:
            time.sleep(1)

        return eventMessage

    def WaitTopic(self, topic, referenceUUID, idIterator: str) -> EventMessage:
        topicMessageWait = TopicMessageWait()
        topicMessageWait.WorkerSource = self.Identity
        topicMessageWait.Topic = topic
        topicMessageWait.IteratorId = idIterator
        topicMessageWait.ReferenceUUID = referenceUUID
        eventMessage = self.Client.WaitTopicMessage(topicMessageWait)

        while eventMessage == None:
            time.sleep(1)

        return eventMessage

    def CreateIteratorEvent(self) -> IteratorMessage:
        return self.Client.CreateIteratorEvent(Empty())
