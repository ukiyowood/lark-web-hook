#!/bin/bash

# GOOS=linux GOARCH=amd64 go build -o webhook-linux webhook.go
# GOOS=linux GOARCH=amd64 go build -o mongoclient mongo.go
GOOS=linux GOARCH=amd64 go build  redis_cli.go