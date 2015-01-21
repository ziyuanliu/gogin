from flask import Flask
from flask import request,redirect, make_response
from json import loads,dumps

app = Flask(__name__)

@app.route('/<stuff>',methods=['GET','POST'])
def hello_world(stuff):
    response = make_response("")

    if stuff=='subscribe':
        time = request.args.get('_')
        channels = request.args.get('channels')
        url = "/ws?_={0}&tag=&time=&eventid=&channels={1}".format(time,channels)
        response.headers["X-Accel-Redirect"] = url
        return response
    else:
        key = request.args.get('id')
        url = "/pub?id={0}".format(key)
        print "redirecting to",url
        return redirect(url,code=307)
        # response = make_response(request.data)
        # response.data = request.data
        # response.headers['Content-Type'] = ""
        # response.headers['X-Accel-Redirect']=url
        # response.headers['X-Accel-Buffering'] = "no"
        # response.headers['Content-Type'] = 'text/plain'
        # response.headers['Cache-Control'] = 'no-cache'
        # return response


if __name__ == '__main__':
    app.run(port=8000)
