
.PHONY: proto unittest e2etest

proto:
	protoc -I/usr/local/include  -I=./proto -I=${GOPATH}/src   --go_out=plugins=grpc:./proto  proto/*.proto
