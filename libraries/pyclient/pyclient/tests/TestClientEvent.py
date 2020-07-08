#! /usr/bin/env python3
# coding: utf-8

import unittest
import uuid
import grpc
from concurrent import futures
from threading import Thread
from time import sleep

from pyclient.event.ClientEvent import ClientEvent
from pyclient.grpc.connectorEvent_pb2_grpc import *
from pyclient.grpc.connectorEvent_pb2 import *
from pyclient.grpc.connector_pb2 import *

TEST_IDENTITY = "TestClientEvent"
DEFAULT_PORT = "50150"
SERVER_STOP_TIME = 1.0

FIXED_TEST_UUID = "3dc28f74-9ad3-4fc7-8b2b-03b7aab660c9"
FIXED_TEST_EVENT = "DummyEvent"
FIXED_TEST_TOPIC = "DummyTopic"
FIXED_TEST_ITERATOR_ID = "DummyIteratorId"
FIXED_TEST_PAYLOAD = "DummyPayload"
FIXED_TEST_TIMEOUT = "10000"

server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))


def serve():
    server.start()
    server.wait_for_termination()


class TestConnectorEventServicer(ConnectorEventServicer):
    """gRPC Server for unit tests"""

    def __init__(self):
        super().__init__()

    def SendEventMessage(self, request, context) -> Empty:
        print('Handle Send Event Message')

        return Empty()

    def WaitEventMessage(self, request, context) -> EventMessage:
        print('Handle Wait Event Message')

        return EventMessage(Topic=request.Topic, ReferenceUUID=request.ReferenceUUID)

    def WaitTopicMessage(self, request, context) -> EventMessage:
        print('Handle Wait Topic Message')

        return EventMessage(Topic=request.Topic, ReferenceUUID=request.ReferenceUUID)

    def CreateIteratorEvent(self, request, context) -> IteratorMessage:
        print('Handle Create Iterator Event')

        return IteratorMessage(Id=FIXED_TEST_ITERATOR_ID)


class TestClientEvent(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        add_ConnectorEventServicer_to_server(
            TestConnectorEventServicer(), server)
        server.add_insecure_port('[::]:'+DEFAULT_PORT)
        thread = Thread(target=serve)
        thread.start()
        cls.client = ClientEvent(TEST_IDENTITY, 'localhost:'+DEFAULT_PORT)

    @classmethod
    def tearDownClass(cls):
        server.stop(SERVER_STOP_TIME)

    def test_send_event(self):
        """
        Test Send Event
        """
        print('TEST ClientEvent.SendEvent(topic, event, referenceUUID, timeout, payload)')

        result = self.client.SendEvent(FIXED_TEST_TOPIC, FIXED_TEST_EVENT,
                                       FIXED_TEST_UUID, FIXED_TEST_TIMEOUT, FIXED_TEST_PAYLOAD)

    def test_wait_event(self):
        """
        Test Wait Event
        """
        print('TEST : ClientEvent.WaitEvent(topic, event, referenceUUID, iteratorId)')

        result = self.client.WaitEvent(
            FIXED_TEST_TOPIC, FIXED_TEST_EVENT, FIXED_TEST_UUID, FIXED_TEST_ITERATOR_ID)

        self.assertEqual(result.Topic, FIXED_TEST_TOPIC)
        self.assertEqual(result.ReferenceUUID, FIXED_TEST_UUID)

    def test_wait_topic(self):
        """
        Test Wait Topic
        """
        print('TEST : ClientEvent.WaitTopic(topic, referenceUUID, idIterator)')

        result = self.client.WaitTopic(
            FIXED_TEST_TOPIC, FIXED_TEST_UUID, FIXED_TEST_ITERATOR_ID)

        self.assertEqual(result.Topic, FIXED_TEST_TOPIC)
        self.assertEqual(result.ReferenceUUID, FIXED_TEST_UUID)

    def test_create_iterator_event(self):
        """
        Test Create Iterator Event
        """
        print('TEST : ClientEvent.CreateIteratorEvent()')

        result = self.client.CreateIteratorEvent()

        self.assertEqual(result.Id, FIXED_TEST_ITERATOR_ID)


if __name__ == '__main__':
    unittest.main()
