import unittest
import uuid
import os

from pyaws.AWS.S3.S3 import S3

aws_region_name = os.environ["aws_region_name"]
aws_access_key_id = os.environ["aws_access_key_id"]
aws_secret_access_key = os.environ["aws_secret_access_key"]

class TestS3(unittest.TestCase):

    @property
    def s3(self) -> S3:
        return self._s3

    @s3.setter
    def s3(self, value: S3):
        self._s3 = value

    @classmethod
    def setUpClass(cls):
        cls.s3 = S3(regionName=aws_region_name, accessKeyId=aws_access_key_id, secretAccessKey=aws_secret_access_key)
    
    def test_create_bucket(self):
        self.assertTrue(self.s3.createBucket(str(uuid.uuid4())))
    
    def test_get_bucket_list(self):
        self.assertGreater(len(self.s3.getBucketList()), 0)
    
if __name__ == '__main__':
    unittest.main()
    