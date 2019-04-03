#!/usr/bin/env bash

export NOW = $(shell date --rfc-3339=ns)

install:
	@echo "configuring dep"
	@dep init
	@echo "installing dependencies..."
	@dep ensure -v
	@echo "install success!"

update:
	@echo "updating dependencies..."
	@dep ensure -v
	@echo "update success!"

build:
	@echo "building main-app..."
	@go build -o mainapp ./cmd/mainapp/
	@echo "build success!"

run:
	@echo "starting app..."
	@./mainapp

all: install build run
quick: build run