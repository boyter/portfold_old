#!/usr/bin/env bash

echo "Running go fmt..."
go fmt ./...

echo "Running unit tests..."
go test ./... || exit

echo "Running integration tests..."
go test -tags=integration ./... || exit
