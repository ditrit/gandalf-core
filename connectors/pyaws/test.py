#! /usr/bin/env python3
# coding: utf-8

from pyaws.AWS.S3.S3 import S3

if __name__ == "__main__":
    aws_access_key_id =  "AKIA3TDVY223GQROR77W"
    aws_secret_access_key = "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"
    region_name = "eu-west-3"

    s3 = S3(region_name, aws_access_key_id, aws_secret_access_key)
    s3.loadBuckets()
    # s3.createBucket("chentest42")

    buckets = s3.getBucketList()

    for bucketName in s3.buckets:
        print(s3.buckets[bucketName].name)