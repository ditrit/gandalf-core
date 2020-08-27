import unittest
import uuid
import os
from io import BytesIO

from pyaws.AWS.IAM.IAM import IAM

aws_region_name = os.environ["aws_region_name"]
aws_access_key_id = os.environ["aws_access_key_id"]
aws_secret_access_key = os.environ["aws_secret_access_key"]

FIXED_TEST_USER_NAME = "userTEST"
FIXED_TEST_USER_NEWNAME = "newUserTest"


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


class TestIAM(unittest.TestCase):

    iam: IAM
    userInfo = None

    @classmethod
    def setUpClass(cls):
        cls.iam = IAM(regionName=aws_region_name, accessKeyId=aws_access_key_id,
                      secretAccessKey=aws_secret_access_key)

    @ordered
    def test_create_user(self):
        self.assertNotEqual(self.iam.createUser(
            userName=FIXED_TEST_USER_NAME), None)

    @ordered
    def test_update_user(self):
        self.assertNotEqual(self.iam.updateUser(
            userName=FIXED_TEST_USER_NAME, newUserName=FIXED_TEST_USER_NEWNAME), None)

    @ordered
    def test_get_user(self):
        self.userInfo = self.iam.getUser(
            userName=FIXED_TEST_USER_NEWNAME)['User']

        self.assertEqual(self.userInfo['UserName'], FIXED_TEST_USER_NEWNAME)

    @ordered
    def test_delete_user(self):
        self.assertEqual(self.iam.deleteUser(
            userName=FIXED_TEST_USER_NEWNAME)['ResponseMetadata']['HTTPStatusCode'], 200)


if __name__ == '__main__':
    unittest.main()
