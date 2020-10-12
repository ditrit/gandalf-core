#! /usr/bin/env python3
# coding: utf-8


try:
    import importlib.resources as pkg_resources
except ImportError:
    # Try backported to PY<37 `importlib_resources`.
    import importlib_resources as pkg_resources

import pyaws.plugins as plugins
import os
import subprocess
from typing import List, Dict
from threading import Thread


class PluginsLoader:

    scripts: List
    loadedPlugins: Dict

    def __init__(self):
        self.scripts = []
        self.loadedPlugins = dict()

    def detectPlugins(self):
        # Get .py files and Filter PluginsLoader.py and __init__.py
        self.scripts = [plugin for plugin in pkg_resources.contents(
            plugins) if plugin[-3:] == '.py' and pkg_resources.is_resource(plugins, plugin) and plugin != "PluginsLoader.py" and plugin != "__init__.py"]
            
        if len(self.scripts) == 0:
            raise Exception("No plugins found !")

    def runPlugins(self):
        for plugin in self.scripts:
            self.loadedPlugins[plugin] = {
                "thread": None,
                "output": "",
                "state": "STOPPED"
            }

            self.loadedPlugins[plugin]["thread"] = Thread(
                target=self.runPlugin(plugin))

        for plugin in self.loadedPlugins.values():
            plugin["thread"].start()

        for plugin in self.loadedPlugins.values():
            plugin["thread"].join()

    def run(self, name):
        plugin = None
        for script in self.scripts:
            if script == name:
                plugin = script
                break

        if plugin is None:
            raise ValueError("This plugin wasn't found !")

        self.loadedPlugins[plugin] = {
            "thread": None,
            "output": "",
            "state": "STOPPED"
        }

        self.loadedPlugins[plugin]["thread"] = Thread(
            target=self.runPlugin(plugin))

        self.loadedPlugins[plugin]["thread"].start()
        self.loadedPlugins[plugin]["thread"].join()

    def runPlugin(self, name):
        self.loadedPlugins[name]["state"] = "RUNNING"

        try:
            with pkg_resources.path(plugins, name) as path:
                process = subprocess.Popen(
                    ["python", path], encoding="utf-8", stdout=subprocess.PIPE, shell=True)
        except FileNotFoundError:
            self.loadedPlugins[name]["output"] += "ERROR: " + \
                name+" cannot be executed !"
            self.loadedPlugins[name]["state"] = "ERROR"
        except subprocess.CalledProcessError as e:
            self.loadedPlugins[name]["output"] += "ERROR: " + \
                name+" has been executed with some errors !" + "\n"
            self.loadedPlugins[name]["output"] += e + "\n"
            self.loadedPlugins[name]["state"] = "ERROR"
        finally:
            while True:
                output = process.stdout.readline()
                if output == '' and process.poll() is not None:
                    break
                if output:
                    self.loadedPlugins[name]["output"] += output.strip() + "\n"

            rc = process.poll()
            process.stdout.close()
            self.loadedPlugins[name]["state"] = "STOPPED"

            return rc
