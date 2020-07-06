#! /usr/bin/env python3
# coding: utf-8

import uuid
import time

from ..grpc.connectorEvent_pb2 import *
from ..grpc.connectorEvent_pb2_grpc import *

class ClientEvent(self):
    @property
    def ClientEventConnection(self):
        return self._ClientEventConnection
    @ClientEventConnection.setter
    def ClientEventConnection(self, value):
        self._ClientEventConnection = value

    @property
    def Identity(self):
        return self._Identity
    @Identity.setter
    def Identity(self, value):
        self._Identity = value

    @property
    def client(self):
        return self._client
    @client.setter
    def client(self, value):
        self._client = value

    def __init__(self, identity: str, clientEventConnection: str):
        self.Identity = identity
        self.ClientEventConnection = clientEventConnection

        conn = grpc.insecure_channel(clientEventConnection)
        self.client = ConnectorEvent(conn)
    
    def SendEvent(self, topic, event, referenceUUID, timeout, payload: str) -> Empty:
        eventMessage = EventMessage()
        eventMessage.Topic = topic
        eventMessage.Timeout = timeout
        eventMessage.UUID = uuid.uuid4()
        eventMessage.Event = event
        eventMessage.Payload = payload
        eventMessage.ReferenceUUID = referenceUUID

        return self.client.SendEvent(eventMessage)

    def WaitEvent(self, topic, event, referenceUUID, idIterator: str) -> EventMessage:
        eventMessageWait = EventMessageWait()
        eventMessageWait.WorkerSource = self.Identity
        eventMessageWait.Topic = topic
        eventMessageWait.Event = event
        eventMessageWait.IteratorId = idIterator
        eventMessageWait.ReferenceUUID = referenceUUID
        eventMessage = self.client.WaitEventMessage(eventMessageWait)

        while eventMessage == None:
            time.sleep(1)

        return eventMessage

    def WaitTopic(self, topic, referenceUUID, idIterator: str) -> EventMessage:
        topicMessageWait = TopicMessageWait()
        topicMessageWait.WorkerSource = self.Identity
        topicMessageWait.Topic = topic
        topicMessageWait.IteratorId = idIterator
        topicMessageWait.ReferenceUUID = referenceUUID
        eventMessage = self.client.WaitTopiceMessage(topicMessageWait)

        while eventMessage == None:
            time.sleep(1)

        return eventMessage

    def CreateIteratorEvent(self) -> IteratorMessage:
        return self.client.CreateIteratorEvent()