


import logging
from sys import stderr, hexversion
logging.basicConfig(stream=stderr)

import hmac
from hashlib import sha1
from json import loads, dumps
from subprocess import Popen, PIPE
from tempfile import mkstemp
from os import access, X_OK, remove, fdopen
from os.path import isfile, abspath, normpath, dirname, join, basename

import requests
from ipaddress import ip_address, ip_network
from flask import Flask, request, abort

from pyclient.models import Options
from pygitlab.server import Urls


url = Urls.returnHashURLS()

application = Flask(__name__)

@application.route( url, methods=['POST'])
def index():
    
    if request.method != 'POST':
        abort(501)
        
    if request.method == 'POST':
        content_length = int(request.headers.get('Content-Length')) # <--- Gets the size of data
        post_data = request.rfile.read(content_length) # <--- Gets the data itself
        #demander à Romain commment ca s'utilise
        clientGandalf=ClientGitlab(config.identity, config.connections)
        clientGandalf.SendEvent("Gitlab", "HOOK", Options("", post_data.decode('utf-8')))
        
            
        


