#!/bin/bash
echo -e "Removing containers first, please wait :)\n"
docker stop defensive_ctf_vuln_running && docker rm defensive_ctf_vuln_running
docker run -it --name defensive_ctf_vuln_running defensive_ctf_vuln_build