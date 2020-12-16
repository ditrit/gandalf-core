from setuptools import setup, find_packages
from os.path import dirname, abspath, join

setup(
    name="pydemo",
    version="0.0.1",
    description="Python demo connector for Gandalf",
    long_description=open(join(dirname(__file__), "README.md")).read(),
    long_description_content_type="text/markdown",
    packages=find_packages(),
    include_package_data=True,
    install_requires=[
        f"pyworker @ file://localhost{dirname(abspath(__file__))}/../py",
        f"pyclient @ file://localhost{dirname(abspath(__file__))}/../../libraries/pyclient",
    ],
    entry_points={
        "console_scripts": [
            "gandalfpydemo=pydemo.__main__:main",
        ]
    },
)