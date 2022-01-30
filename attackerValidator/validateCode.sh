#!/bin/bash

# Test/validate each of the .go files
echo -e "{Various checks on the files}\n"
for goFile in $(find . | grep "\.go")
do
    echo "[$goFile]"
    go vet $goFile
    go fmt $goFile
    ~/.local/bin/semgrep --config "p/golang"
done

# Check the Dockerfile
echo "[Dockerfile]"
~/.local/bin/semgrep --config "p/dockerfile"

# Use gosec for another SAST security check
echo -e "\n{Gosec}\n"
~/go/bin/gosec .

# Finally run golangcli-lint on all of the code
echo -e "\n{Golangcli-lint}\n"
~/go/bin/golangci-lint run .