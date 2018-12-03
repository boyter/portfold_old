#!/usr/bin/env bash

pkill portfold

echo "Running go fmt..."
go fmt ./...

echo "Running unit tests..."
go test ./... || exit

echo "Building application and starting..."
go build && ./portfold > /dev/null 2> /dev/null &

echo "Waiting for server to startup... 10 seconds..."
sleep 7
echo "3..."
sleep 1
echo "2..."
sleep 1
echo "1..."
sleep 1

echo "Running integration tests..."

GREEN='\033[1;32m'
RED='\033[0;31m'
NC='\033[0m' # no colour

if go test -tags=integration ./... ; then
	echo -e "${GREEN}--------- ALL TESTS PASSED ---------"
else
	exit
fi

echo -e "${NC}"
pkill portfold