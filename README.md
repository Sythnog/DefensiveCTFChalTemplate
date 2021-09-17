# Defensive CTF - Challenge template
This repo is intended to be used as a template, or starting point, for the development of challenges for defensive CTFs.

## The idea behind this repo
We have two machines the "judge" which will run the attacks and validations against the other "vulnerable machine". If the judge deems the vulnerability solved, then the user will be given the flag.

## How to use
There are two docker scripts to build and run the containers. Be aware that currently the attackingIP is hardcoded to be http://172.17.0.3/, which is the second IP assigned by docker on most Linux devices. Aka. an assumtion is made that these are the only two running containers on the machine. Feel free to change the variable for IP in attackerValidator/webService.py to fit your needs.

SSH password for `root` on the vulnerable machine is by default `notS0S3cr3tSSHPa55`, which can be changed in the Dockerfile if desired.

Once you have used SSH to access the machine, and have made your changes, then the service can be restarted with `killall python3; python3 /opt/notFoundService/vulnService.py`

## The current example
The service is vulnerable to server-side template injection SSTI.