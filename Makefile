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
	@echo "building article-app..."
	@go build -o articleapp ./cmd/articleapp/
	@echo "build success!"

run:
	@echo "starting app..."
	@./articleapp & wait

all: install build run
