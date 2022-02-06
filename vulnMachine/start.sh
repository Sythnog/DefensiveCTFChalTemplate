#!/bin/bash
$(/usr/sbin/sshd -D) &
$(python3 vuln_service.py) &
tail -f /dev/null