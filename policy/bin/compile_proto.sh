#!/bin/bash
protoc \
    --go_out=$GOPATH/src/github.com/MauricioAntonioMartinez/aws/policy/backend \
    --go-grpc_out=$GOPATH/src/github.com/MauricioAntonioMartinez/aws/policy/backend \
    -I $GOPATH/src \
    -I $GOPATH/src/github.com/MauricioAntonioMartinez/aws/policy/proto \
    $GOPATH/src/github.com/MauricioAntonioMartinez/aws/policy/proto/backend.proto
