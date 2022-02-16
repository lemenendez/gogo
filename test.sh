#!/bin/bash
rm $PWD/cover.out
go test ./pkg/ -coverprofile=cover.out -coverpkg=./...
# Command cover
# https://golang.org/cmd/cover/
rm $PWD/cover.html
go tool cover -html=$PWD/cover.out -o $PWD/cover.html