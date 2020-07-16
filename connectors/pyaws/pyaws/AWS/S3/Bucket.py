#! /usr/bin/env python3
# coding: utf-8

class Bucket:

    @property
    def client(self):
        return self._client

    @client.setter
    def client(self, value):
        self._client = value

    @property
    def name(self) -> str:
        return self._name

    @name.setter
    def name(self, value: str):
        self._name = value

    @property
    def region(self) -> str:
        return self._region

    @region.setter
    def region(self, value: str):
        self._region = value

    def __init__(self, client, region: str, name: str):
        self.client = client
        self.region = region
        self.name = name