#!/bin/bash
echo -e "Removing containters first, please wait :)\n"
docker stop defensive_ctf_attval_running && docker rm defensive_ctf_attval_running
docker run -it --name defensive_ctf_attval_running defensive_ctf_attval_build