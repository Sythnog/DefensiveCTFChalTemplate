#!/bin/bash
docker stop defensive_ctf_attval_running && docker rm defensive_ctf_attval_running
docker stop defensive_ctf_vuln_running && docker rm defensive_ctf_vuln_running
docker system prune -a