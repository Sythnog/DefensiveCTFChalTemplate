import sys
from attacks.mainAttack import AttackWrapper
from flask_restful import Resource, Api, reqparse
from validators.mainValidator import ValidatorWrapper
from flask import Flask, make_response, render_template

# Setup the flask service
app = Flask(__name__)
api = Api(app)

# Global parameter to decide if the task was done
fulfilledTask = False
vulnMachineLocation = "http://172.17.0.3/"

# -== Helper functions ==-

# Function to return file and mime-type, as a Flask response
# Note: File has to be in ./templates folder
def MakeResponse(fileLocation, mimeType):
    # Make a response for Flask using its render_template functionality
    response = make_response(render_template(fileLocation))

    # Set the mime-type of the response, so that the browser knows how to handle the response
    response.headers["Content-Type"] = mimeType
    return response

# -== Endpoint functionality ==-

# The / endpoint, which displays the frontpage
class Root(Resource):
    def get(self):
        return MakeResponse("index.html", "text/html")

# The /start endpoint, which is responsible for triggering the validation
class Start(Resource):
    def post(self): # TODO: Multithreading-lock, in order to limit ability to DoS the hosting infrastructure
        # Get the location of the vuln machine
        global vulnMachineLocation
        
        # Run attacks the and validations
        # TODO: multithreading, so that the attacks and validations will be mixed as much as the python GIL allow
        attackResult = AttackWrapper(vulnMachineLocation)
        validationResult = ValidatorWrapper(vulnMachineLocation)

        # As we use bool to determine a non-failed check, then ensure it is bool
        if type(attackResult) == bool and type(validationResult) == bool:
            # Check if attack and validation both were good (returned True)
            if attackResult and validationResult:
                # Since the task has been judged sucessfull, then set the global variable to true
                # This way /done will always show the flag from this point, so that the users do not have...
                # ...to re-run the attack/validation if they accidentially close the window before submitting the flag.
                global fulfilledTask
                fulfilledTask = True

            # Return to the user that the validation was successfull -> /done
            return { "failed": False }
        else:    
            # Tell the user that an error occured
            # NOTE: For debug purposes, then this might be a good spot to print attackResult and/or validationResult
            return { "failed": True }, 500

# The /done endpoint, which is responsible for giving the flag if the objective was fulfilled
class Done(Resource):
    def get(self):
        # Obtain the global param to see if task was fullfilled, and return appropriate HTML
        global fulfilledTask
        if fulfilledTask:
            # Since we need to add content to the HTML, let's make the response here
            # TODO: Input validation on sys.argv[1]?
            response = make_response(render_template("doneGood.html", flag=sys.argv[1]))
            response.headers["Content-Type"] = "text/html"
            return response
        else:    
            return MakeResponse("doneBad.html", "text/html")

# The / endpoint, which displays the page letting users know an error occured
class Failed(Resource):
    def get(self):
        return MakeResponse("failed.html", "text/html")

# -== Endpoints ==-
api.add_resource(Root, "/")
api.add_resource(Start, "/start")
api.add_resource(Done, "/done")
api.add_resource(Failed, "/failed")

# -== Start server ==-
if not len(sys.argv) == 2:
    print("[!] ERROR! Please (only) enter a flag as the first parameter")
else:
    # Not threaded Flask, in order to have full control over concurrency ourselves
    app.run(threaded=False, debug=False, port=80, host="0.0.0.0")