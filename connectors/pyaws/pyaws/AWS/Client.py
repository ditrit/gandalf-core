#! /usr/bin/env python3
# coding: utf-8

import boto3


class Client:

    @property
    def service(self):
        return self._service

    @service.setter
    def service(self, value):
        self._service = value

    @property
    def awsAccessKeyId(self):
        return self._awsAccessKeyId

    @awsAccessKeyId.setter
    def awsAccessKeyId(self, value):
        self._awsAccessKeyId = value

    @property
    def awsSecretAccessKey(self):
        return self._awsSecretAccessKey

    @awsSecretAccessKey.setter
    def awsSecretAccessKey(self, value):
        self._awsSecretAccessKey = value

    @property
    def regionName(self):
        return self._regionName

    @regionName.setter
    def regionName(self, value):
        self._regionName = value

    @property
    def client(self):
        return self._client

    @client.setter
    def client(self, value):
        self._client = value

    def __init__(self, service: str, regionName: str, accessKeyId: str, secretAccessKey: str):
        self.service = service
        self.regionName = regionName
        self.accessKeyId = accessKeyId
        self.awsSecretAccessKey = secretAccessKey

        self.client = boto3.client(service, region_name=regionName,
                                   aws_access_key_id=accessKeyId, aws_secret_access_key=secretAccessKey)
