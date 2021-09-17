// Function used to send a POST request to start the judge
function StartJudge() {
    // Give visual feedback to the user that the button press worked ;)
    document.getElementById("feedbackText").innerText = "Running the check, plese wait :)"

    // Make async request to get the data
    let httpReq = new XMLHttpRequest();

    // When we get a response, we will do this:
    httpReq.onreadystatechange = function() {
        if (httpReq.readyState == XMLHttpRequest.DONE) {
            // Parse the respone, and check if the backend failed
            const respJson = JSON.parse(httpReq.response);
            if (respJson["failed"] == true) {
                // As we failed, then redirect user to the failed page
                window.location = `http://${window.location.hostname}/failed`;
            } else {
                // Send the user to the done page, where they will get the result from the validation/attack
                window.location = `http://${window.location.hostname}/done`;
            }
            console.log(options)
        }
    }

    // Set the URL to /start and send request
    httpReq.open("post", `http://${window.location.hostname}/start`, true);
    httpReq.send(); 
}

// Function used to go back to the frontpage
function Homepage() {
    // Change the browser location to the current domainname
    window.location = `http://${window.location.hostname}/`;
}