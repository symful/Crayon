#!/usr/bin/bash

DIR=$(dirname -- "$0";)
antlr4 -visitor -Dlanguage=Go -o $DIR/../parser $DIR/Crayon.g4