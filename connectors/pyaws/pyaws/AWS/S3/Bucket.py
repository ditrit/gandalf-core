#! /usr/bin/env python3
# coding: utf-8

from typing import BinaryIO, Dict
from botocore.exceptions import ClientError


class Bucket:

    client: str
    name: str
    region: str

    def __init__(self, client, region: str, name: str):
        self.client = client
        self.region = region
        self.name = name

    def uploadFile(self, fileName: str, objectName: str, fileObject: BinaryIO) -> bool:
        if objectName is None:
            objectName = fileName

        try:
            if fileObject is None:
                self.client.upload_file(fileName, self.name, objectName)
            else:
                self.client.upload_fileobj(fileObject, self.name, objectName)
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

    def putBucketAcl(self, acl: str, acp: Dict):
        try:
            self.client.put_bucket_acl(
                Bucket=self.name, ACL=acl, AccessControlPolicy=acp)
        except ClientError as err:
            print(err)
            return False

        return True
