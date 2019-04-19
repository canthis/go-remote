#!/bin/bash

set GOOS=linux
set GOARCH=386
rm "GoRemote-linux-i386"
go build -ldflags="-s -w" -o "GoRemote-linux-i386"
rice append --exec "GoRemote-linux-i386"

set GOOS=linux
set GOARCH=amd64
rm "GoRemote-linux-amd64"
go build -ldflags="-s -w" -o "GoRemote-linux-amd64"
rice append --exec "GoRemote-linux-amd64"

