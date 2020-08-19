#! /usr/bin/env python3
# coding: utf-8

from botocore.exceptions import ClientError
from mypy_boto3_s3 import S3Client
from typing import Dict

from ..Client import Client
from .Bucket import Bucket


class S3(Client):

    client: S3Client
    buckets: Dict[str, Bucket]

    def __init__(self, regionName: str, accessKeyId: str, secretAccessKey: str):
        Client.__init__(self, 's3', regionName, accessKeyId, secretAccessKey)

        self.buckets = {}
        self.loadBuckets()

    def loadBuckets(self):
        bucketList = self.client.list_buckets()['Buckets']

        for bucket in bucketList:
            self.buckets[bucket['Name']] = Bucket(self.client, self.awsAccessKeyId, bucket['Name'], self.regionName, bucket)

    def createBucket(self, name: str, region=None):
        if region is None:
            region = self.regionName

        try:
            self.buckets[name] = Bucket(self.client, self.awsAccessKeyId, name, region)
        except ClientError as err:
            raise err

        return True

    def deleteBucket(self, name: str):
        try:
            self.client.delete_bucket(Bucket=name)
        except ClientError as err:
            print(err)
            return False

        return True

    def getBucketList(self):
        return self.client.list_buckets()['Buckets']
