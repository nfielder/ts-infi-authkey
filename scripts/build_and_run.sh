#!/usr/bin/env bash

# Get the parent directory of the script
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE")/.." && pwd)"

# change into that directory
cd "$DIR"

make clean
make build

./bin/ts-infi-authkey "$@"
