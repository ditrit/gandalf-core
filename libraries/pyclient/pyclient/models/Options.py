#! /usr/bin/env python3
# coding: utf-8

class Options:

    timeout: str
    payload: str

    def __init__(self, timeout: str, payload: str):
        self.timeout = timeout
        self.payload = payload
