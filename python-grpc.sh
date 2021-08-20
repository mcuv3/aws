#!/usr/bin/bash

python3 -m grpc_tools.protoc  -I ./proto --python_out=./lambda/prebuild/python --grpc_python_out=./lambda/prebuild/python ./proto/lambda.proto