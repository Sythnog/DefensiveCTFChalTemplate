// Send a request to the API to get the judge to start an audit
function StartJudge() {
    // Write feedback message on the current page
    document.getElementById("feedbackText").innerText = "Running the judge, please wait for the result :D"

    // Make async request to get the data
    let httpReq = new XMLHttpRequest();

    // When we get a response, we will do this:
    httpReq.onreadystatechange = function() {
        if (httpReq.readyState == XMLHttpRequest.DONE) {
            // As the request is done, then send the user to see their result
            window.location = `${window.location.origin}/result`;
        }
    }

    // Set the URL to /audit and send an async request
    httpReq.open("post", `${window.location.origin}/audit`, true);
    httpReq.send(); 
}