from flask import Flask, request, jsonify
app = Flask(__name__)

import redis
redisInst = redis.StrictRedis(host='172.16.88.145', port=6379, db=0)


@app.route("/set", methods=['POST'])
def setredis():
    key = request.json['key']
    value = request.json['value']
    result = redisInst.set(key, value)
    return jsonify(
        key=key,
        value=value,
        result=result
    )


@app.route("/get/<key>", methods=['GET'])
def getredis(key):
    value = redisInst.get(key)
    return jsonify(
        key=key,
        value=value.decode()
    )