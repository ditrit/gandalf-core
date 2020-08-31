import uuid
import os
import time
from io import BytesIO

from pyaws.AWS.S3.S3 import S3
from pyaws.AWS.S3.S3Control import S3Control
from pyaws.AWS.S3.Bucket import Bucket

aws_region_name = os.environ["aws_region_name"]
aws_access_key_id = os.environ["aws_access_key_id"]
aws_secret_access_key = os.environ["aws_secret_access_key"]

FIXED_TEST_WEBSITE_CONFIG = {
    'ErrorDocument': {'Key': 'error.html'},
    'IndexDocument': {'Suffix': 'index.html'}
}

FIXED_TEST_INDEX_FILE = b"<html><head><title>Gandalf AWS Connector S3 Static Website Test (Index)</title><meta charset=\"UTF-8\"></head><body><h1>It's working !</h1></body></html>"
FIXED_TEST_ERROR_FILE = b"<html><head><title>Gandalf AWS Connector S3 Static Website Test (Error)</title><meta charset=\"UTF-8\"></head><body><h1>Error !!!</h1></body></html>"


if __name__ == '__main__':
    print("Script has started !")
    s3: S3 = S3(regionName=aws_region_name, accessKeyId=aws_access_key_id,
                secretAccessKey=aws_secret_access_key)
    s3control: S3Control = S3Control(regionName=aws_region_name, accessKeyId=aws_access_key_id,
                                     secretAccessKey=aws_secret_access_key)
    testBucket: str = str(uuid.uuid4())

    # Create a bucket
    s3.createBucket(testBucket)

    # Get the bucket list
    bucketList = s3.getBucketList()
    print(len(bucketList))

    # Get the bucket
    bucket: Bucket = s3.buckets[testBucket]

    # Put the files config in the bucket
    bucket.putWebsite(FIXED_TEST_WEBSITE_CONFIG)

    # Verify if files config are added to the bucket
    print(len(bucket.getWebsite()))

    # Upload files to bucket
    with BytesIO(FIXED_TEST_INDEX_FILE) as f:
        bucket.uploadFile("index.html", "index.html",
                          {'ACL': 'public-read'}, f)

    # create an acl for the bucket
    bucket.putBucketAcl('public-read')

    # create an access point
    s3control.createAccessPoint(bucket=bucket.name, name="TestAP")

    # delete file config
    bucket.deleteWebsite()

    # delete bucket
    s3.deleteBucket(testBucket)

    print("end of the script ...")