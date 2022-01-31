package main

// -=== Setup and init ===-

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Allow for us to easily group challenge data
type Challenge struct {
	challName              string
	challSolved            bool
	challWebLocation       string
	challFlag              string
	challSSHPass           string
	challFileLoc           string
	challRestartServiceCMD string
}

// -=== Main functionality ===-

// The main function, which outlines and constructs the service
func main() {
	// Initialize the Gin framework, and HTML templates
	r := gin.Default()
	r.LoadHTMLGlob("webTemplates/*")

	// Store the challenge data in a struct
	// NOTE: Since they are hard-coded, no input validation will be written
	challInfo := Challenge {
		challName:              "SSTI Example Challenge",
		challSolved:            false,
		challWebLocation:       "172.17.0.3",
		challFlag:              "flag{ex@mp1e_FlaG_123}",
		challSSHPass:           "notS0S3cr3tSSHPa55",
		challFileLoc:           "/opt/notFoundService",
		challRestartServiceCMD: "killall python3; python3 /opt/notFoundService/vuln_service.py",
	}

	// -== Endpoints =-

	// The main page of the judge
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"challName": challInfo.challName,
		})
	})

	// The FAQ page about the concept and challegne
	r.GET("/faq", func(c *gin.Context) {
		c.HTML(200, "faq.html", gin.H{
			"challName":              challInfo.challName,
			"challWebLocation":       challInfo.challWebLocation,
			"challSSHPass":           challInfo.challSSHPass,
			"challFileLoc":           challInfo.challFileLoc,
			"challRestartServiceCMD": challInfo.challRestartServiceCMD,
		})
	})

	// The page which will diplay the result, once an assessment has been made
	r.GET("/result", func(c *gin.Context) {
		// If the challenge has been solved, then return the flag
		if challInfo.challSolved {
			c.HTML(200, "resultSolved.html", gin.H{
				"challName":      challInfo.challName,
				"challStatusMsg": "You solved the challenge!",
				"challFlag":      challInfo.challFlag,
			})
			return
		}

		// Upon an invalid result, default to tell this to the user
		c.HTML(200, "resultBad.html", gin.H{
			"challName":      challInfo.challName,
			"challStatusMsg": "The challenge has not been solved :(",
		})
	})

	// -== Run the service ==-

	// Run the service, and set the port
	if err := r.Run(":80"); err != nil {
		fmt.Print("Unable to start the service :(")
	}
}
