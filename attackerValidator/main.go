package main

// -=== Setup and init ===-

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os/exec"
	"strings"
)

// Allow for us to easily group challenge data
type Challenge struct {
	name                    string
	solved                  bool
	statusMsg               string
	webLocation             string
	webLocationFull         string
	flag                    string
	SSHPass                 string
	fileLoc                 string
	restartServiceCMD       string
	auditNormalFunc         string
	auditNormalFuncNumTests int
	auditVuln               string
}

// -=== Assessment functionality ===-

// Perform a security check using Nuclei
// Returns a bool (true if successful audit, false if not satisfiable audit)
func PerformCheck(AuditTemplatePath string, VulnMachineDomain string, CheckVuln bool, AuditNum int) bool {
	// Run the nuclei command (input validation on VulnMachineDomain not done, as it is hard-coded... can be done for added security)
	out, err := exec.Command("/home/attValUser/go/bin/nuclei", "-nts", "-silent", "-t", AuditTemplatePath, "-u", VulnMachineDomain).Output()

	// Upon a cmd-failure, return false and print the error for debugging purposes
	if err != nil {
		fmt.Printf("%s", err)
		return false
	}

	// Get the amount of lines in the output, for further analysis
	linesInOut := len(strings.Split(string(out), "\n"))

	// Check if nothing was found (except for the default newline)
	if linesInOut == 1 {
		// If we are checking for vulns, then this is the success criteria
		if CheckVuln {
			return true
		}
		// If we have no normal-func audits, this is success...
		if AuditNum == 0 {
			return true
		}
		// ...otherwise, this indicates that a normal-func test did not succeeed
		return false
	}

	// If we are checking normal-func, then success criteria is AuditNum + the default newline
	if !CheckVuln {
		expectedLines := AuditNum + 1
		if linesInOut == expectedLines {
			return true
		}
	}

	// As no other check succeeded, then default to returning false for all other cases
	return false
}

// -=== Main functionality ===-

// The main function, which outlines and constructs the service
func main() {
	// Initialize the Gin framework, incl. setting upd static files and HTML templates
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob("webTemplates/*")
	router.Static("/static/", "./webStatic")

	// Store the challenge data in a struct
	// NOTE: Since they are hard-coded, no input validation will be written
	challInfo := Challenge {
		name:                    "SSTI Example Challenge",
		solved:                  false,
		statusMsg:               "No audit has run, please start one on the frontpage",
		webLocation:             "172.17.0.3",
		webLocationFull:         "http://172.17.0.3",
		flag:                    "flag{ex@mp1e_FlaG_123}",
		SSHPass:                 "notS0S3cr3tSSHPa55",
		fileLoc:                 "/opt/notFoundService",
		restartServiceCMD:       "killall python3; python3 /opt/notFoundService/vuln_service.py",
		auditNormalFunc:         "/opt/attVal/auditTemplates/normalFuncCheck.yml",
		auditNormalFuncNumTests: 3,
		auditVuln:               "/opt/attVal/auditTemplates/vulnCheck.yml",
	}

	// -== Endpoints =-

	// The main page of the judge
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H {
			"name": challInfo.name,
		})
	})

	// The FAQ page about the concept and challegne
	router.GET("/faq", func(c *gin.Context) {
		c.HTML(200, "faq.html", gin.H {
			"name":              challInfo.name,
			"webLocation":       challInfo.webLocation,
			"SSHPass":           challInfo.SSHPass,
			"fileLoc":           challInfo.fileLoc,
			"restartServiceCMD": challInfo.restartServiceCMD,
		})
	})

	// The page which will diplay the result, once an assessment has been made
	router.GET("/result", func(c *gin.Context) {
		// If the challenge has been solved, then return the flag
		if challInfo.solved {
			c.HTML(200, "resultSolved.html", gin.H {
				"name":           challInfo.name,
				"challStatusMsg": challInfo.statusMsg,
				"flag":           challInfo.flag,
			})
			return
		}

		// Upon an invalid result, default to tell this to the user
		c.HTML(200, "resultBad.html", gin.H {
			"name":           challInfo.name,
			"challStatusMsg": challInfo.statusMsg,
		})
	})

	// Run an audit against the vulnerable machine
	router.POST("/audit", func(c *gin.Context) {
		// Run both tests
		vulnCheck := PerformCheck(challInfo.auditVuln, challInfo.webLocationFull, true, 0)
		normalFuncCheck := PerformCheck(challInfo.auditNormalFunc, challInfo.webLocationFull, false, challInfo.auditNormalFuncNumTests)

		// If both passed, then set chall to have been solved
		if vulnCheck && normalFuncCheck {
			challInfo.solved = true
			challInfo.statusMsg = "You solved the challenge!"
		}

		// If only vuln was solved, tell this
		if vulnCheck && !normalFuncCheck {
			challInfo.statusMsg = "You have solved the vulnerability, BUT also removed too much of the normal/expected functionality"
			challInfo.solved = false
		}

		// If vuln not solved, but normal is working
		if !vulnCheck && normalFuncCheck {
			challInfo.statusMsg = "The vulnerbility has not been solved, but the normal/expected functionality works"
			challInfo.solved = false
		}

		// If both are false
		if !vulnCheck && !normalFuncCheck {
			challInfo.statusMsg = "The vulnerability still exist, and the normal/expected functionality no longer works"
			challInfo.solved = false
		}

		// Return the state to the user
		c.JSON(200, gin.H {
			"isVulnerable":         vulnCheck,
			"normalFuncStillWorks": normalFuncCheck,
		})
	})

	// -== Run the service ==-

	// Run the service, and set the port
	if err := router.Run(":80"); err != nil {
		fmt.Print("Unable to start the service :(\n")
	}
}
