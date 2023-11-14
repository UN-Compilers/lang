#!/bin/sh

alias antlr4='java -jar $PWD/antlr-4.13.1-complete.jar'
antlr4 -Dlanguage=Go -o ../parsing -listener -visitor -package parsing *.g4