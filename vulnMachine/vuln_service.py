"""A simple API which will tell you that the page is not what you are looking for"""

# -== Imports and setup ==-

# Import Flask, and setup the flask app
from flask import Flask, render_template_string
app = Flask(__name__)

# -== Endpoint functionality ==-

@app.route("/")
def index():
    """The page a user will be shown when they go to the root-page of the website"""

    index_content = "<h1>This is not the page you are looking for</h1>" + \
                    "<p>No pages will be found on this website...</p>"
    return render_template_string(index_content)

@app.route("/<path>")
def not_found(path):
    """On any path: Send a HTML page stating x was not found"""

    not_found_str = f"<h1>{path} was not found :(</h1>"
    return render_template_string(not_found_str)

# -== Start server ==-
if __name__ == "__main__":
    app.run(threaded=True, debug=False, port=80, host="0.0.0.0")
