#! /usr/bin/env python3
# coding: utf-8

import unittest
import json
import os

from pyclient.configuration.ClientConfiguration import ClientConfiguration

FIXED_TEST_CONFIG_PATH = "./ClientTestConfig.json"
FIXED_TEST_CONFIG_CONTENTS = {
    "ClientCommandConnection": "DummyCommand",
    "ClientEventConnection": "DummyEvent",
    "Identity": "DummyClientIdentity"
}


class TestClientConfiguration(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        with open(FIXED_TEST_CONFIG_PATH, 'w', encoding='utf-8') as jsonFile:
            json.dump(FIXED_TEST_CONFIG_CONTENTS, jsonFile,
                      ensure_ascii=False, indent=4)

    @classmethod
    def tearDownClass(cls):
        if os.path.isfile(FIXED_TEST_CONFIG_PATH):
            os.remove(FIXED_TEST_CONFIG_PATH)

    def test_client_configuration(self):
        config = ClientConfiguration(FIXED_TEST_CONFIG_PATH)

        self.assertEqual(config.ClientCommandConnection,
                         FIXED_TEST_CONFIG_CONTENTS["ClientCommandConnection"])
        self.assertEqual(config.ClientEventConnection,
                         FIXED_TEST_CONFIG_CONTENTS["ClientEventConnection"])
        self.assertEqual(
            config.Identity, FIXED_TEST_CONFIG_CONTENTS["Identity"])
