from setuptools import setup, find_packages
from os.path import dirname, abspath

setup(
    name='pyworker',
    version='0.1.1',
    description = 'Base worker for Python gandalf connectors',
    packages=find_packages(),
    install_requires = [
        f'pyclient @ file://localhost{dirname(abspath(__file__))}/../../libraries/pyclient'
    ]
)