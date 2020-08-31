#! /usr/bin/env python3
# coding: utf-8

import json

from ..WorkerAws import WorkerAws

from ..plugins.PluginsLoader import PluginsLoader
from threading import Thread

class WorkerPlugins(WorkerAws):
    loader: PluginsLoader

    def __init__(self, version, commandes, config):
        super().__init__(version, commandes, config)

        self.loader = PluginsLoader()

    
    def Execute(self, clientGandalf, version):
        self.loader.detectPlugins()

    def runPlugin(self):
        id = self.clientGandalf.CreateIteratorCommand()
        print(id)

        command = self.clientGandalf.WaitCommand(
            "RUN_PLUGIN", id, self.version)
        print(command)

        payload = json.loads(command.Payload)

        plugin = payload['plugin'] if 'plugin' in payload else None

        if not plugin:
            self.reportError(uuid=command.UUID, command="RUN_PLUGIN", error="Plugin name must be provided")

        try:
            response = self.loader.runPlugin(plugin)    
        except ValueError as err:
            self.reportError(uuid=command.UUID, command="RUN_PLUGIN", error=err)
        else:
            self.clientGandalf.SendEvent(command.UUID, "SUCCES", {"10000", "Plugin executed with success !"})

    def Run(self):
        runPlugin = Thread(target=self.runPlugin())
    