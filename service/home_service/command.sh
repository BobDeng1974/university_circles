#!/bin/bash

make
gsed -i "s/,omitempty//g" $(grep ",omitempty" -rl ./pb/home/*)

GOOS=linux GOARCH=amd64 go build -o home_services main.go plugin.go