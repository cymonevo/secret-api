#!/usr/bin/env bash
SHELL := /bin/bash

export NOW = $(shell date --rfc-3339=ns)

install:
	@echo "installation task started"
	@echo "- configuring dep..."
	@dep init
	@echo "-- success"
	@echo "- installing dependencies..."
	@dep ensure -v
	@echo "-- success"
	@echo "- configuring environment..."
	@echo "$(cat .ignore)" >> .git/info/exclude
	@echo "-- success"
	@echo "installation success!"

ignore:
	@echo "ingoring files..."
	@source command.sh; ignore

unignore:
	@echo "revert ingoring files..."
	@source command.sh; unignore

update:
	@echo "update task started"
	@echo "- updating dependencies..."
	@dep ensure -v
	@echo "-- success"
	@echo "update success!"

build:
	@echo "build task started"
	@echo "- building main-app..."
	@go build -o mainapp ./cmd/mainapp/
	@echo "-- success"
	@echo "build success!"

run:
	@echo "starting app..."
	@./mainapp

mq:
	@echo "starting mq server..."
	@nsqlookupd & nsqd --lookupd-tcp-address=127.0.0.1:4160 & nsqadmin --lookupd-http-address 127.0.0.1:4161

all: install build run
quick: build run