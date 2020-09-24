PACKAGES=$(shell go list ./... )
all: install

install: go-get go-install

go-get:
	go get ./...

go-install:
	@echo 'install all packages'
	go install ./...
test: 
	@echo  'installing all packages'
	go test ./...

help: Makefile 
	  @echo 'nothing'

.PHONY: help test install
