import json
from flask import Flask
from flask import request

app = Flask(__name__)

@app.route("/get", methods=["GET"])
def get():
  if request.headers.get('Authorization') != "Bearer foobar":
    return "", 403
  
  return "", 200

@app.route("/post", methods=["POST"])
def post():
  if request.cookies.get("admin") != "1":
    return "", 403
  
  body = json.loads(request.data, strict=False)
  
  if body.get("username") != "guest" or body.get("password") != "guest":
    return "", 403
  
  return "", 200

if __name__ == '__main__':
    app.run()
