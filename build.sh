#!/bin/bash

GOOS=linux;GOARCH=amd64;go build -o webhook-linux webhook.go