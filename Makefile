#!/usr/bin/env bash
SHELL := /bin/bash

export NOW=$(shell date --rfc-3339=ns)

install:
	@echo "installation task started"
	@echo -n "- configuring dep... "
	@dep init
	@echo "ok"
	@echo -n "- installing dependencies... "
	@dep ensure -v
	@echo "ok"
	@echo -n "- configuring environment... "
	@echo "$(cat .ignore)" >> .git/info/exclude
	@echo "ok"

ignore:
	@echo "ingoring files..."
	@source command.sh; ignore

unignore:
	@echo "revert ingoring files..."
	@source command.sh; unignore

update:
	@echo "update task started"
	@echo -n "- updating dependencies... "
	@dep ensure -v
	@echo "ok"

build:
	@echo "build task started"
	@echo -n "- building main-app... "
	@go build -o mainapp ./cmd/mainapp/
	@echo "ok"

run:
	@echo "starting app..."
	@./mainapp

mq:
	@echo "starting mq server..."
	@nsqlookupd & nsqd --lookupd-tcp-address=127.0.0.1:4160 & nsqadmin --lookupd-http-address 127.0.0.1:4161

all: install build run
quick: build run