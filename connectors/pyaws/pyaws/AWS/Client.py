#! /usr/bin/env python3
# coding: utf-8

import boto3
import typing

from abc import abstractmethod

T = typing.TypeVar('T', bound='Client')


class Client:

    service: str

    awsAccessKeyId: str
    awsSecretAccessKey: str

    regionName: str

    @property
    @abstractmethod
    def client(self) -> T:
        return self._client

    @client.setter
    @abstractmethod
    def client(self, value: T):
        self._client: T = value

    def __init__(self, service: str, regionName: str, accessKeyId: str, secretAccessKey: str):
        self.service = service
        self.regionName = regionName
        self.awsAccessKeyId = accessKeyId
        self.awsSecretAccessKey = secretAccessKey

        self.client = boto3.client(service, region_name=regionName,
                                   aws_access_key_id=accessKeyId, aws_secret_access_key=secretAccessKey)
