#!/bin/bash
echo -e "-=== Removing old containers ===-\n"
docker stop defensive_ctf_attval_running && docker rm defensive_ctf_attval_running
docker stop defensive_ctf_vuln_running && docker rm defensive_ctf_vuln_running

echo -e "\n-=== Running new containers ===-\n"

# Attacker/Validator machine
attvalID=$(docker run -d --name defensive_ctf_attval_running defensive_ctf_attval_build)
attvalIP=$(docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' defensive_ctf_attval_running)
echo -e "Vuln machine can be found on http://$attvalIP\nDocker ID is $attvalID\n"

# Vuln machine
vulnID=$(docker run -d --name defensive_ctf_vuln_running defensive_ctf_vuln_build)
vulnIP=$(docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' defensive_ctf_vuln_running)
echo -e "Vuln machine can be found on http://$vulnIP\nDocker ID is $vulnID\n"
