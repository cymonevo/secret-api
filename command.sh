#!/bin/bash

function ignore() {
  git update-index --assume-unchanged $(grep -v '^#' .ignore | tr '\n' ' ')
}

function unignore() {
  git update-index --no-assume-unchanged $(grep -v '^#' .ignore | tr '\n' ' ')
}
