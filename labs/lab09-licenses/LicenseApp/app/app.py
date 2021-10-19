from flask import Flask, render_template, request
import requests
import json


app = Flask(__name__)


@app.route('/')
def license_check():
    message_to_display = ''
    response = requests.get(
    "http://kotsadm:3000/license/v1/license",
    headers={
        "content-type":"application/json"
    },
    )
    response_json = response.json()
    message_to_display += '<br> License Assigned To:' + response_json["assignee"]
    custom_fields = response_json['fields']
    for custom_field in custom_fields:
        print (message_to_display)
        if custom_field['field']=='subscription-tier' :
           message_to_display += '<br><br><br> The Current Subscription Tier is: ' + custom_field["value"]
        
    return message_to_display
if __name__ == '__main__':
    app.run(host='0.0.0.0')