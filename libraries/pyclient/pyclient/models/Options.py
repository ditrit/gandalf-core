#! /usr/bin/env python3
# coding: utf-8

class Options:

    Timeout: str
    Payload: str

    def __init__(self, timeout: str, payload: str):
        self.Timeout = timeout
        self.Payload = payload
