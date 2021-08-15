# AWS - mcuve

This is a replication of the most common aws services, build with grpc,docker,psql and kubernetes.

## Requirements

- Docker engine
- Golang v16 >
- Air cli
- Evans cli

## Steps

- `docker-compose up`
- `evans -r -p 6002` for IAM
- `evans -r -p 6003` for SQS
- `evans -r -p 6004` for Lambda
