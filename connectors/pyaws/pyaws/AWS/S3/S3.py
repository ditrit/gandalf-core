#! /usr/bin/env python3
# coding: utf-8

from botocore.exceptions import ClientError
import mypy_boto3_s3 as botoS3
from typing import Dict

from ..Client import Client
from .Bucket import Bucket

class S3(Client):
    
    @property
    def client(self) -> botoS3.S3Client:
        return self._client

    @client.setter
    def client(self, value: botoS3.S3Client):
        self._client = value

    @property
    def buckets(self) -> Dict[str, Bucket]:
        return self._buckets

    @buckets.setter
    def buckets(self, value: Dict[str, Bucket]):
        self._buckets = value

    def __init__(self, regionName: str, accessKeyId: str, secretAccessKey: str):
        super().__init__('s3', regionName, accessKeyId, secretAccessKey)

        self.buckets = {}
        self.loadBuckets()

    def loadBuckets(self):
        bucketList = self.client.list_buckets()['Buckets']

        for bucket in bucketList:
            self.buckets[bucket['Name']] = Bucket(self.client, self.regionName, bucket['Name'])
    
    def createBucket(self, name: str):
        try:
            self.client.create_bucket(Bucket=name, CreateBucketConfiguration={'LocationConstraint': self.regionName})
            self.buckets[name] = Bucket(self.client, self.regionName, name)
        except ClientError as err:
            print(err)
            return False
        
        return True
    
    def getBucketList(self):
        return self.client.list_buckets()['Buckets']