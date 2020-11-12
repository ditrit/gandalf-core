#! /usr/bin/env python3
# coding: utf-8

from sys import argv
from typing import List

from pyclient.ClientGandalf import ClientGandalf


def SendCommands(clientGandalf: ClientGandalf, major: int, minor: int, commands: List[str]) -> bool:
    print("SEND COMMAND LIST WORKER")
    validate = clientGandalf.SendCommandList(major, minor, commands)

    return validate.Valid
