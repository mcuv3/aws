#!/usr/bin/bash
protoc \                                                                
  --plugin=protoc-gen-ts=./web/node_modules/.bin/protoc-gen-ts \
  -I ./proto \
  --js_out=import_style=commonjs,binary:./web/src/gen \
  --ts_out=service=grpc-web:./web/src/gen \
  ./proto/*.proto