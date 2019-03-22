#!/usr/bin/env bash

export NOW = $(shell date --rfc-3339=ns)

test:

build:
    @echo "building article-app..."
    @go build -o articleapp ./cmd/articleapp/
    @echo "building success!"

run:
    @echo "starting app..."
    @./articleapp & wait

all: test build
