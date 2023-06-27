#!/usr/bin/bash

DIR=$(dirname -- "$0";)
go build -ldflags="-w -s" $DIR/crayon.go
time $DIR/crayon $1