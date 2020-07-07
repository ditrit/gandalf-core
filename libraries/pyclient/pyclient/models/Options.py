#! /usr/bin/env python3
# coding: utf-8

class Options:
    @property
    def timeout(self):
        return self._timeout
    @timeout.setter
    def timeout(self, value):
        self._timeout = value

    @property
    def payload(self):
        return self._payload
    @payload.setter
    def payload(self, value):
        self._payload = value

    def __init__(self, timeout, payload: str):
        self.timeout = timeout
        self.payload = payload