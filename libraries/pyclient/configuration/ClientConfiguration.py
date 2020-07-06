#! /usr/bin/env python3
# coding: utf-8

import json

class ClientConfiguration(self):

    @property
    def ClientCommandConnection(self):
        return self._ClientCommandConnection
    @ClientCommandConnection.setter
    def ClientCommandConnection(self, value):
        self._ClientCommandConnection = value

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

    def __init__(self, clientCommandConnection: str, clientEventConnection: str, identity: str):
        self.ClientCommandConnection = clientCommandConnection
        self.ClientEventConnection = clientEventConnection
        self.Identity = identity

    def __init__(self, path: str):
        self.LoadConfiguration(path)

    def LoadConfiguration(self, path: str):
        with open(path, 'r') as jsonFile:
            data = json.load(jsonFile)

            self.ClientCommandConnection = data['ClientCommandConnection']
            self.ClientEventConnection = data['ClientEventConnection']
            self.Identity = data['Identity']
