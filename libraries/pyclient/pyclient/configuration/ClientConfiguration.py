#! /usr/bin/env python3
# coding: utf-8

import json


class ClientConfiguration:

    clientCommandConnection: str
    clientEventConnection: str
    identity: str

    def __init__(self, clientCommandConnection: str, clientEventConnection: str, identity: str):
        self.clientCommandConnection = clientCommandConnection
        self.clientEventConnection = clientEventConnection
        self.identity = identity

    def __init__(self, path: str):
        self.LoadConfiguration(path)

    def LoadConfiguration(self, path: str):
        with open(path, 'r') as jsonFile:
            data = json.load(jsonFile)

            self.clientCommandConnection = data['ClientCommandConnection']
            self.clientEventConnection = data['ClientEventConnection']
            self.identity = data['Identity']
