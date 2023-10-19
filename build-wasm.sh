#!/bin/sh

cd cmd/markdown-to-html
GOOS=js GOARCH=wasm go build -o main.wasm
