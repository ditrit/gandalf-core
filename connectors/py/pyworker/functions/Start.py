#! /usr/bin/env python3
# coding: utf-8

from sys import argv

from pyclient.ClientGandalf import ClientGandalf


def Start() -> ClientGandalf:
    return ClientGandalf(argv[0], argv[1], argv[2].split(","))
