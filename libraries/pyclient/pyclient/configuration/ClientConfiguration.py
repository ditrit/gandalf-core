#! /usr/bin/env python3
# coding: utf-8

import json


class ClientConfiguration:

    ClientCommandConnection: str
    ClientEventConnection: str
    Identity: str

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
