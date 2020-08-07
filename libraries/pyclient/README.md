# Python gRPC Client
This is the python port of the gRPC Client for Gandalf connectors.

## BUILD
You need to have installed python3, pip and virtualenv.

### Install python dependencies
#### First step is to create the env folder
```python3 -m virtualenv env```
#### Then activate it
On Windows  : ```env\Scripts\activate.bat``` \
On Linux    : ```source env/bin/activate```
#### And finally install the dependencies
```pip install -r requirements.txt``` \
*You can do this step without setting up the virtualenv but the dependencies will be installed globally* \
*If you are using Spyder you can install using the console with : ```!pip install -r requirements.txt```*

## Testing
```python3 -m unittest discover```

## USAGE
### Quickstart
#### Add pyclient to your project
Setup you virtualenv install if needed. \
Then you can add pyclient directly from the github repository (it's only available through this link for now) :  \
```pip install "git+https://github.com/ditrit/gandalf.git@armorique#egg=pyclient&subdirectory=libraries/pyclient"``` \

Or you can load it from somewhere else, as an example for a connector (located at ```/connectors/myconnector```) in this project : \ 
```pip install -e "../../libraries/pyclient"``` \
#### Example file
```
from pyclient.ClientGandalf import ClientGandalf

if __name__ == "__main__":
  demoClient = ClientGandalf(identity="TestIdentity", timeout="10000", clientConnections=["IP:PORT"])  # Create the client, IP and PORT need to be changed
  
  demoClient.SendCommandList(version=42, commands=["CommandA", "CommandB", "CommandC"]) # Send the command list
```
  
  
