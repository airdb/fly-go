#!/usr/bin/env bash

LDFLAGS="-s -w"

function until::build() {
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$LDFLAGS" -o ./output/main main.go
}

$1