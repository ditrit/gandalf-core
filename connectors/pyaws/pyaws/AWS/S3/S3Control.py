#! /usr/bin/env python3
# coding: utf-8

from botocore.exceptions import ClientError
from mypy_boto3_s3control import client as S3ControlClient
from typing import Dict

from ..Client import Client

class S3Control(Client):

    client: S3ControlClient

    def __init__(self, regionName: str, accessKeyId: str, secretAccessKey: str):
        Client.__init__(self, 's3control', regionName, accessKeyId, secretAccessKey)

    def createAccessPoint(self, bucket:str, name: str):
        try:
            self.client.create_access_point(AccountId=self.awsAccessKeyId, Name=name, Bucket=bucket, PublicAccessBlockConfiguration={'BlockPublicAcls': False, 'IgnorePublicAcls': False, 'BlockPublicPolicy': False, 'RestrictPublicBuckets': True})
        except ClientError as err:
            print(err)
            return False

        return True
    
