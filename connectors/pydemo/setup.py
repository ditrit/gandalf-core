from setuptools import setup, find_packages
from os.path import dirname, join

setup(
    name='pydemo',
    version='0.0.1',
    description='Python demo connector for Gandalf',
    long_description=open(join(dirname(__file__), 'README.md')).read(),
    long_description_content_type="text/markdown",
    packages=find_packages(),
    entry_points={
        'console_scripts': [
            'gandalf_pydemo=pydemo.WorkerDemo:main'
        ]
    }
)