#! /usr/bin/env python3
# coding: utf-8

from typing import BinaryIO, Dict
from botocore.exceptions import ClientError
from mypy_boto3_s3 import S3Client

class Bucket:

    client: S3Client
    name: str
    region: str
    accountId: str
    bucket: Dict

    def __init__(self, client: S3Client, accountId: str, name: str, region: str, bucket: Dict = None):
        self.client: S3Client = client
        self.accountId = accountId
        self.name = name
        self.region = region

        self.bucket = bucket if bucket is not None else self.client.create_bucket(
            Bucket=name, CreateBucketConfiguration={'LocationConstraint': region})

    def uploadFile(self, fileName: str, objectName: str, extraArgs, fileObject: BinaryIO) -> bool:
        if objectName is None:
            objectName = fileName

        try:
            if fileObject is None:
                if extraArgs is None:
                    self.client.upload_file(fileName, self.name, objectName)
                else:
                    self.client.upload_file(
                        fileName, self.name, objectName, ExtraArgs=extraArgs)
            else:
                if extraArgs is None:
                    self.client.upload_fileobj(
                        fileObject, self.name, objectName)
                else:
                    self.client.upload_fileobj(
                        fileObject, self.name, objectName, ExtraArgs=extraArgs)
        except ClientError as err:
            print(err)
            return False

        return True

    def downloadFile(self, fileName: str, objectName: str, fileObject: BinaryIO) -> bool:
        if objectName is None:
            objectName = fileName

        try:
            if fileObject is None:
                self.client.download_file(self.name, objectName, fileName)
            else:
                self.client.download_fileobj(self.name, objectName, fileObject)
        except ClientError as err:
            print(err)
            return False

        return True

    def putWebsite(self, config: Dict) -> bool:
        try:
            self.client.put_bucket_website(
                Bucket=self.name, WebsiteConfiguration=config)
        except ClientError as err:
            print(err)
            return False

        return True

    def getWebsite(self) -> Dict:
        return self.client.get_bucket_website(Bucket=self.name)

    def deleteWebsite(self) -> bool:
        try:
            self.client.delete_bucket_website(Bucket=self.name)
        except ClientError as err:
            print(err)
            return False

        return True

    def getBucketAcl(self) -> Dict:
        try:
            return self.client.get_bucket_acl(Bucket=self.name)
        except ClientError as err:
            print(err)
        return None

    def putBucketAcl(self, acl: str):
        try:
            self.client.put_bucket_acl(Bucket=self.name, ACL=acl)
        except ClientError as err:
            print(err)
            return False

        return True

    def putPublicAccessBlock(self, publicAccessBlockConfiguration):
        self.client.put_public_access_block(
            Bucket=self.name, PublicAccessBlockConfiguration=publicAccessBlockConfiguration)
