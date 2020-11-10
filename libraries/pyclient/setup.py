from setuptools import setup, find_packages

setup(
    name='pyclient',
    version='0.1.1',
    description='Python Grpc client library for Gandalf',
    packages=find_packages(),
    install_requires = [
        'grpcio',
        'grpcio-tools',
        'protobuf',
        'six'
    ]
)
