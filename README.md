# Defensive CTF - Challenge template
This repo is intended to be used as a template, or starting point, for the development of challenges for defensive CTFs.

## The idea behind this repo
We have two machines the "judge" which will run the attacks and validations against the other "vulnerable machine". If the judge deems the vulnerability solved, then the user will be given the flag.

## How to use for testing
There are two docker scripts to build and run the containers. Be aware that currently the attackingIP is hardcoded to be http://172.17.0.3/, which is the second IP assigned by docker on most Linux devices. Aka. an assumption is made that these are the only two running containers on the machine. Feel free to change the variable for IP in attackerValidator/webService.py to fit your needs.

SSH password for `root` on the vulnerable machine is by default `notS0S3cr3tSSHPa55`, which can be changed in the Dockerfile if desired.

Once you have used SSH to access the machine, and have made your changes, then the service can be restarted with `killall python3; python3 /opt/notFoundService/vuln_service.py`

## How to develop your own challenge, based upon this repo
- Decide which vulnerability you wish for the participants to solve.
- Create the vulnerable service in the `vulnMachine` folder *(easiest to just edit the dockerfile in that folder)*
- Change the data for `challInfo` in `attackerValidator/main.go` to fit your challenge.
- Edit the following [Nuclei](https://nuclei.projectdiscovery.io/) scripts: `attackerValidator/auditTemplates/vulnCheck.yml` and `attackerValidator/auditTemplates/normalFuncCheck.yml`
- Make sure everything *(still)* works :D
- ???
- Profit!?

## Core Technologies used
- Golang for the validation service, using the [Gin](https://gin-gonic.com/) framework.
- [Nuclei](https://nuclei.projectdiscovery.io/) is used for checking the vulnerable service.
- Docker is used for encapsulation.

## How do I run the CTF?
We recommend using a platform where you are able to give the participants their own instances of both attacker and vulnerable machine.
Specifically, we recommend using the [HAAUKINS](https://docs.haaukins.com/) CTF platform.

## The current example
The service is vulnerable to server-side template injection SSTI.
