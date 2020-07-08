#! /usr/bin/env python3
# coding: utf-8

from sys import argv
from typing import List

from pyclient.ClientGandalf import ClientGandalf

def SendCommands(clientGandalf: ClientGandalf, version: int, commands: List[str]):
    clientGandalf.SendCommandList(version, commands)