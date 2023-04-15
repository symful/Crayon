#!/usr/bin/bash

go build -ldflags="-w -s" crayon.go
time ./crayon $1