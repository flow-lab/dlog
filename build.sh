#!/usr/bin/env bash
set -e

GOOS=linux go build -o main
zip deployment.zip main