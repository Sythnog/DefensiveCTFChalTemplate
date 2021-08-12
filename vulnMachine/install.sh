#!/bin/bash
# No need to spend build-time updating this to the absolute newest version?
apt update && apt install python3 python3-pip -y
pip3 install -r requirements.txt

# TODO: Setup SSH access