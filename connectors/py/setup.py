from setuptools import setup, find_packages

setup(
    name='pyworker',
    version='0.1',
    description = 'Base worker for Python gandalf connectors',
    packages=find_packages(),
    install_requires = [
        'pyclient'
    ],
    dependency_links = [
        '../../libraries/pyclient'
    ]
)