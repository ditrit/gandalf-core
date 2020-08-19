import unittest
import uuid
import os
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


def make_orderer():
    order = {}

    def ordered(f):
        order[f.__name__] = len(order)
        return f

    def compare(a, b):
        return [1, -1][order[a] < order[b]]

    return ordered, compare


ordered, compare = make_orderer()
unittest.defaultTestLoader.sortTestMethodsUsing = compare


class TestS3(unittest.TestCase):

    @property
    def s3(self) -> S3:
        return self._s3

    @s3.setter
    def s3(self, value: S3):
        self._s3: S3 = value

    @property
    def s3control(self) -> S3Control:
        return self._s3control

    @s3control.setter
    def s3control(self, value: S3Control):
        self._s3control: S3Control = value

    @property
    def testBucket(self) -> str:
        return self._testBucket

    @testBucket.setter
    def testBucket(self, value: str):
        self._testBucket: str = value

    @classmethod
    def setUpClass(cls):
        cls.s3: S3 = S3(regionName=aws_region_name, accessKeyId=aws_access_key_id,
                        secretAccessKey=aws_secret_access_key)
        cls.s3control: S3Control = S3Control(regionName=aws_region_name, accessKeyId=aws_access_key_id,
                        secretAccessKey=aws_secret_access_key)
        cls.testBucket: str = str(uuid.uuid4())

    @ordered
    def test_create_bucket(self):
        self.assertTrue(self.s3.createBucket(self.testBucket))

    @ordered
    def test_get_bucket_list(self):
        bucketList = self.s3.getBucketList()

        self.assertGreater(len(bucketList), 0)

    @ordered
    def test_put_bucket_website(self):
        bucket: Bucket = self.s3.buckets[self.testBucket]

        self.assertTrue(bucket.putWebsite(FIXED_TEST_WEBSITE_CONFIG))

    @ordered
    def test_get_bucket_website(self):
        bucket: Bucket = self.s3.buckets[self.testBucket]

        self.assertGreater(len(bucket.getWebsite()), 0)

    @ordered
    def test_upload_file_to_bucket(self):
        bucket: Bucket = self.s3.buckets[self.testBucket]

        with BytesIO(FIXED_TEST_INDEX_FILE) as f:
            bucket.uploadFile("index.html", "index.html", {'ACL': 'public-read'}, f)

    @ordered
    def test_make_public_read_bucket(self):
        bucket: Bucket = self.s3.buckets[self.testBucket]

        bucket.putBucketAcl('public-read')
    
    @ordered
    def test_aclpub(self):
        bucket: Bucket = self.s3.buckets[self.testBucket]
        self.s3control.createAccessPoint(bucket=bucket.name, name="TestAP")

    # @ordered
    # def test_delete_bucket_website(self):
    #     bucket: Bucket = self.s3.buckets[self.testBucket]

    #     self.assertTrue(bucket.deleteWebsite())

    # @ordered
    # def test_delete_bucket(self):
    #     self.assertTrue(self.s3.deleteBucket(self.testBucket))


if __name__ == '__main__':
    unittest.main()
