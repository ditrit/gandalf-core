from setuptools import setup, find_packages

setup(
    name='pyworker',
    version='0.1.1',
    description = 'Base worker for Python gandalf connectors',
    packages=find_packages(),
    install_requires = [
        'pyclient@git+https://github.com/ditrit/gandalf.git@pyaws-connector#egg=pyclient&subdirectory=libraries/pyclient'
    ]
)