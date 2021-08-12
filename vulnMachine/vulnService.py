from flask import Flask, render_template_string

# Setup the flask app
app = Flask(__name__)

# -== Endpoint functionality ==-

# Index page
@app.route("/")
def Index():
    indexContent = "<h1>This is not the page you are looking for</h1><p>No pages will be found on this website...</p>"
    return render_template_string(indexContent)

# Will send a HTML page stating x was not found
@app.route("/<path>")
def NotFound(path):
    notFoundStr = "<h1>{} was not found :(</h1>".format(path)
    return render_template_string(notFoundStr)

# -== Start server ==-
app.run(threaded=True, debug=False, port=80, host="0.0.0.0")