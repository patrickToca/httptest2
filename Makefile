#!/usr/bin/make -f

# This Makefile is an example of what you could feed to scantest's -command flag.

default: test

test: build
	go generate ./...
	go test -v ./...

build:
	go build ./...

cover:
	go build
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

docs:
	go get -u github.com/robertkrimen/godocdown/godocdown
	go install github.com/robertkrimen/godocdown/godocdown
	godocdown > README.md