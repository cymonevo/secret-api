#!/usr/bin/env bash

export NOW = $(shell date --rfc-3339=ns)

build:
	@echo "building article-app..."
	@go build -o articleapp ./cmd/articleapp/
	@echo "building success!"

run:
	@echo "starting app..."
	@./articleapp & wait

all: build run
