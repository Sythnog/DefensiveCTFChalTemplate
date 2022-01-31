#!/bin/bash

# Test/validate each of the .go files
echo -e "{Various checks on the files}\n"
for goFile in $(find . | grep "\.go")
do
    echo "[$goFile]"
    go vet $goFile
    go fmt $goFile
done

# Check the Dockerfile
echo "[Dockerfile]"
~/.local/bin/semgrep --config "p/dockerfile"

# Perform SAST security checks
echo -e "\n{Gosec}\n"
~/go/bin/gosec .
echo -e "\n{Semgrep: Golang}\n"
~/.local/bin/semgrep --config "p/golang"
echo -e "\n{Semgrep: CI}\n"
~/.local/bin/semgrep --config "p/ci"
echo -e "\n{Semgrep: OWASP Top 10}\n"
~/.local/bin/semgrep --config "p/owasp-top-ten"
echo -e "\n{Semgrep: R2C Security Audit}\n"
~/.local/bin/semgrep --config "p/r2c-security-audit"

# Finally run golangcli-lint on all of the code
echo -e "\n{Golangcli-lint}\n"
~/go/bin/golangci-lint run .