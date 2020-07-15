#! /usr/bin/env python3
# coding: utf-8

from botocore.exceptions import ClientError

from ..Client import Client
from .Bucket import Bucket

class S3(Client):
    
    @property
    def buckets(self):
        return self._buckets

    @buckets.setter
    def buckets(self, value):
        self._buckets = value

    def __init__(self, regionName: str, accessKeyId: str, secretAccessKey: str):
        super().__init__('s3', regionName, accessKeyId, secretAccessKey)

        self.buckets = {}

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