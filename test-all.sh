#!/usr/bin/env bash

# For printing out coloured output
GREEN='\033[1;32m'
RED='\033[0;31m'
NC='\033[0m' # no colour

# Its possible that an existing portfold process is running so clean it up just in case
pkill portfold

echo "Running go fmt..."
go fmt ./...

echo "Running unit tests..."
# If this process fails we want this to exit which is what || achieves
go test ./... || exit

echo "Building application and starting..."
# Build the application and start it with all output redirected to /dev/null so we don't see it and
# then use & to fork the process off into the background.
go build && ./portfold > /dev/null 2> /dev/null &

echo "Waiting for server to startup..."
# Loop waiting for the application to run or until a timeout is reached
cnt=1;
while :
do
    echo "Attempt $cnt..."
    # -s is silent mode and -m sets the connection timeout to 1 second
    curl -s -m 1 "http://localhost:8080/health-check/" > /dev/null 2>&1

    # $? is the exit value of the last command and if it passes break out of the loop
    # NB if you put any command above this one this will then always pass
    if [[ $? -eq 0 ]]; then
    echo -e "Server has finished startup process commencing integration tests"
        break
    fi

    # After 9 attempts exit the loop
    if [[ cnt -eq 10 ]]; then
    echo
        echo -e "${RED}Server has not started expect integration tests to fail"
        echo -e "${NC}"
        break
    fi

    # Increment the count so we know how many attempts were made
    ((cnt++))
done

echo "Running integration tests..."

# Important to have -count=1 so the tests always run at least once
# We also set -tags to run integration tests
if go test -count=1 -tags=integration ./... ; then
	echo -e "${GREEN}--------- ALL TESTS PASSED ---------"
else
    echo -e "${RED}--------- TESTS FAILED ---------"
fi

echo -e "${NC}"
pkill portfold