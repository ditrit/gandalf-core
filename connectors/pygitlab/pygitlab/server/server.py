

import logging
import sys 
import json
from flask import Flask, request, abort

from pyclient.ClientGandalf import ClientGandalf
from pyclient.models import Options
from pygitlab.server.serverUrl import Urls

config = json.loads(sys.stdin.read())
url = Urls.returnHashURLS()

application = Flask(__name__)

@application.route( url, methods=['POST'])
def index():
    
    if request.method != 'POST':
        abort(501)
        
    if request.method == 'POST':
        content_length = int(request.headers.get('Content-Length')) # <--- Gets the size of data
        post_data = request.rfile.read(content_length) # <--- Gets the data itself
       
        clientGandalf=ClientGandalf( identity=config.Identity, clientConnections=config.ClientEventConnection)
        clientGandalf.SendEvent("Gitlab", "HOOK", Options("", post_data.decode('utf-8')))
        
            
        


