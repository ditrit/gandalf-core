#! /usr/bin/env python3
# coding: utf-8

from botocore.exceptions import ClientError
from mypy_boto3_iam import IAMClient
from typing import Dict, List

from ..Client import Client


class IAM(Client):

    client: IAMClient

    def __init__(self, regionName: str, accessKeyId: str, secretAccessKey: str):
        Client.__init__(self, 'iam', regionName, accessKeyId, secretAccessKey)

    def createUser(self, name: str, permissions: str = "", tags: List = [], path: str = None):
        try:

            response = self.client.create_user(
                UserName=name, PermissionsBoundary=permissions, Tags=tags)

            return response
        except ClientError as err:
            raise err

        return None

    def updateUser(self, name: str, newName: str, newPath: str = None):
        try:

            response = self.client.update_user(
                UserName=name, NewUserName=newName)

            return response
        except ClientError as err:
            raise err

        return None
