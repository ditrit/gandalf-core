import json, sys
from .WorkerDemo import WorkerDemo

if __name__ == "__main__":
    commands = list()
    version = int()

    config = json.loads(sys.stdin.read())

    workerDemo = WorkerDemo(version, commands, config)

    workerDemo.Run()