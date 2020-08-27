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

    def createUser(self, userName: str, permissions: str = None, tags: List = [], path: str = None):
        try:
            if permissions == None:
                response = self.client.create_user(
                    UserName=userName, Tags=tags)
            else:
                response = self.client.create_user(
                    UserName=userName, PermissionsBoundary=permissions, Tags=tags)

            return response
        except ClientError as err:
            raise err

        return None

    def getUser(self, userName: str):
        try:
            response = self.client.get_user(UserName=userName)

            return response
        except ClientError as err:
            raise err

        return None

    def updateUser(self, userName: str, newUserName: str, newPath: str = None):
        try:
            response = self.client.update_user(
                UserName=userName, NewUserName=newUserName)

            return response
        except ClientError as err:
            raise err

        return None

    def deleteUser(self, userName: str):
        try:
            response = self.client.delete_user(UserName=userName)

            return response
        except ClientError as err:
            raise err

        return None

    def listUsers(self):
        return self.getPaginator('list_users')

    def getPaginator(self, interface: str):
        try:
            response = self.client.get_paginator(interface)

            return response
        except ClientError as err:
            raise err

        return None

    def createPolicy(self, policyName: str, policyDocument: str):
        try:
            response = self.client.create_policy(
                PolicyName=policyName, PolicyDocument=policyDocument)

            return response
        except ClientError as err:
            raise err

        return None

    def getPolicy(self, policyArn: str):
        try:
            response = self.client.get_policy(PolicyArn=policyArn)

            return response
        except ClientError as err:
            raise err

        return None

    def deletePolicy(self, policyArn: str):
        try:
            response = self.client.delete_policy(PolicyArn=policyArn)

            return response
        except ClientError as err:
            raise err

        return None

    def listPolicies(self, scope: str = 'All', onlyAttached: bool = False, pathPrefix: str = "", policyUsageFilter: str = "", marker: str = "", maxItems: int = 100):
        try:
            response = self.client.list_policies(Scope=scope, OnlyAttached=onlyAttached, PathPrefix=pathPrefix,
                                                 PolicyUsageFilter=policyUsageFilter, Marker=marker, MaxItems=maxItems)

            return response
        except ClientError as err:
            raise err

        return None

    def attachRolePolicy(self, policyArn: str, roleName: str):
        try:
            response = self.client.attach_role_policy(
                PolicyArn=policyArn, RoleName=roleName)

            return response
        except ClientError as err:
            raise err

        return None

    def detachRolePolicy(self, policyArn: str, roleName: str):
        try:
            response = self.client.detach_role_policy(
                PolicyArn=policyArn, RoleName=roleName)

            return response
        except ClientError as err:
            raise err

        return None

    def attachUserPolicy(self, policyArn: str, userName: str):
        try:
            response = self.client.attach_user_policy(
                PolicyArn=policyArn, UserName=userName)

            return response
        except ClientError as err:
            raise err

        return None

    def detachUserPolicy(self, policyArn: str, userName: str):
        try:
            response = self.client.detach_user_policy(
                PolicyArn=policyArn, UserName=userName)

            return response
        except ClientError as err:
            raise err

        return None

    def attachGroupPolicy(self, policyArn: str, groupName: str):
        try:
            response = self.client.attach_group_policy(
                PolicyArn=policyArn, GroupName=groupName)

            return response
        except ClientError as err:
            raise err

        return None

    def detachGroupPolicy(self, policyArn: str, groupName: str):
        try:
            response = self.client.detach_group_policy(
                PolicyArn=policyArn, GroupName=groupName)

            return response
        except ClientError as err:
            raise err

        return None

    def createGroup(self, groupName: str, path: str = None):
        try:
            if path == None:
                response = self.client.create_group(GroupName=groupName)
            else:
                response = self.client.create_group(
                    GroupName=groupName, Path=path)

            return response
        except ClientError as err:
            raise err

        return None

    def getGroup(self, groupName: str, marker: str = "", maxItems: int = 100):
        try:
            response = self.client.get_group(
                GroupName=groupName, MaxItems=maxItems, Marker=marker)

            return response
        except ClientError as err:
            raise err

        return None

    def updateGroup(self, groupName: str, newGroupName: str, newPath: str = None):
        try:
            if newPath == None:
                response = self.client.update_group(
                    GroupName=groupName, NewGroupName=newGroupName)
            else:
                response = self.client.update_group(
                    GroupName=groupName, NewGroupName=newGroupName, NewPath=newPath)

            return response
        except ClientError as err:
            raise err

        return None

    def deleteGroup(self, groupName: str):
        try:
            response = self.client.delete_group(GroupName=groupName)

            return response
        except ClientError as err:
            raise err

        return None


    def createUserAccessKey(self, userName: str):
        try:
            response = self.client.create_access_key(UserName=userName)
            return response['AccessKey']

        except ClientError as err:
            raise err

        return None


    def revokeUserAccessKey(self, keyId: str, userName: str = None):
        try:
            if userName == None:
                response = self.client.delete_access_key(AccessKeyId=keyId)
            else:
                response = self.client.delete_access_key(UserName=userName, AccessKeyId=keyId)

        except ClientError as err:
            raise err

        return None


    def updateUserAccessKey(self, keyId: str, status: bool, userName: str = None):

        status = "Active" if status == True else "Inactive"

        try:
            if userName == None:
                response = self.client.update_access_key(AccessKeyId=keyId, Status=status)
            else:
                response = self.client.update_access_key(AccessKeyId=keyId, Status=status, UserName=userName)
                
        except ClientError as err:
            raise err

    def listAccessKeys(self, userName: str = None, maxItems: int = 100):
        try:
            response = self.client.list_access_keys(UserName=userName, MaxItems=maxItems)
            keys = response['AccessKeyMetadata']

            while response['IsTruncated'] == True:
                response = self.client.list_access_keys(UserName=userName, Marker=response['Marker'], MaxItems=maxItems)
                keys += response['AccessKeyMetadata']
            
            return keys
            
        except ClientError as err:
            raise err

        return None