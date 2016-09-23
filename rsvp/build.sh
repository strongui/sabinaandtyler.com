#!/bin/bash

set -e
set -x

GOOS=linux GOARCH=amd64 go build -o main
zip -r lambda.zip main index.js
rm main
