import json
import sys

from .WorkerDemo import WorkerDemo


def main():
    major, minor = 1, 0

    print("Pydemo connector - VERSION {}.{}".format(major, minor))

    print("START 0 : Load config from standard input")
    config = json.loads(sys.stdin.read())
    print("Loaded config :\n{}".format(config))

    print("START 1 : Instanciate worker")
    workerDemo = WorkerDemo(major, minor, config)

    commands = {"RUN_TEST_1": workerDemo.runTest1, "RUN_TEST_2": workerDemo.runTest2}

    for k, v in commands:
        print("START REGISTER : {}".format(k))
        workerDemo.RegisterCommandsFuncs(k, v)

    workerDemo.Run()


if __name__ == "__main__":
    main()
