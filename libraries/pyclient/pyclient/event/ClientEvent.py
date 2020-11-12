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

    client: ConnectorEventStub

    @property
    def clientEventConnection(self) -> str:
        return self.clientConnection

    @clientEventConnection.setter
    def clientEventConnection(self, value: str):
        self.clientConnection = value

    def __init__(self, identity: str, clientEventConnection: str):
        super().__init__(identity, clientEventConnection)

        try:
            conn = grpc.insecure_channel(clientEventConnection)
            
            self.client = ConnectorEventStub(conn)
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

        return self.client.SendEventMessage(eventMessage)

    def WaitEvent(self, topic, event, referenceUUID, idIterator: str) -> EventMessage:
        eventMessageWait = EventMessageWait()
        eventMessageWait.WorkerSource = self.identity
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
        topicMessageWait.WorkerSource = self.identity
        topicMessageWait.Topic = topic
        topicMessageWait.IteratorId = idIterator
        topicMessageWait.ReferenceUUID = referenceUUID
        eventMessage = self.client.WaitTopicMessage(topicMessageWait)

        while eventMessage == None:
            time.sleep(1)

        return eventMessage

    def CreateIteratorEvent(self) -> IteratorMessage:
        return self.client.CreateIteratorEvent(Empty())
