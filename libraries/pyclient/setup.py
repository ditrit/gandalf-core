from setuptools import setup, find_packages

setup(
    name='pyclient',
    version='0.1',
    description='Python Grpc client librairy for Gandalf',
    packages=find_packages(),
    install_requires = [
        'grpcio',
        'grpcio-tools',
        'protobuf',
        'six'
    ]
)
