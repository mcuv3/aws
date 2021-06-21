#!/bin/bash
export BACKEND_CONFIG=$GOPATH/src/github.com/MauricioAntonioMartinez/aws/policy/configs/dev.yaml
go run $GOPATH/src/github.com/MauricioAntonioMartinez/aws/policy/cli/main.go
