#! /usr/bin/env python3
# coding: utf-8

from botocore.exceptions import ClientError
from mypy_boto3_iam import IAMClient
from typing import Dict

from ..Client import Client

class IAM(Client):

    client: IAMClient

    def __init__(self, regionName: str, accessKeyId: str, secretAccessKey: str):
        Client.__init__(self, 'iam', regionName, accessKeyId, secretAccessKey)

    def createUser(self, name: str):
        try:
            response = self.client.create_user(UserName=name)
            print(response)
        except ClientError as err:
            raise err

        return True
