#!/bin/bash
$(/usr/sbin/sshd -D) &
$(python3 vulnService.py)