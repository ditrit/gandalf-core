from setuptools import setup, find_packages
from os.path import dirname, join

setup(
    name='pydemo',
    version='0.0.1',
    description='Python demo connector for Gandalf',
    long_description=open(join(dirname(__file__), 'README.md')).read(),
    long_description_content_type="text/markdown",
    packages=find_packages(),
    install_requires = [
        'pyworker@git+https://github.com/ditrit/gandalf.git@pyaws-connector#egg=pyworker&subdirectory=connectors/py',
        'pyclient@git+https://github.com/ditrit/gandalf.git@pyaws-connector#egg=pyclient&subdirectory=libraries/pyclient'
    ]
)