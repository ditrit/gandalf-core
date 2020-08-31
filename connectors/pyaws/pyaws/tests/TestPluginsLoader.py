import unittest
import uuid
import os
from io import BytesIO

from pyaws.plugins.PluginsLoader import PluginsLoader


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


class TestPluginsLoader(unittest.TestCase):

    loader: PluginsLoader

    @classmethod
    def setUpClass(cls):
        cls.loader = PluginsLoader()

    @ordered
    def test_load_plugins(self):
        self.loader.detectPlugins()
    
    @ordered
    def test_start_plugins(self):
        self.loader.runPlugins()

    @ordered
    def test_read_output_plugins(self):
        for plugin in self.loader.loadedPlugins.values():
            print(plugin["output"])


if __name__ == '__main__':
    unittest.main()
