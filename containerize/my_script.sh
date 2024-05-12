#!/bin/bash

echo "This is a bash script running inside a Docker container!"
echo "Here's some system information:"
echo "Operating System: $(uname -s)"
echo "Kernel Version: $(uname -r)"
echo "Architecture: $(uname -m)"

# https://docs.docker.com/engine/security/