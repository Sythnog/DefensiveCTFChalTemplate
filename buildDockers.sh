#!/bin/bash
cd attackerValidator && docker build -t defensive_ctf_attval_build . && cd ..
cd vulnMachine && docker build -t defensive_ctf_vuln_build . && cd ..