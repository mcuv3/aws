
.PHONY: proto unittest e2etest

proto_sqs:
	protoc -I/usr/local/include  --proto_path=proto --go_out=plugins=grpc:proto \
	--grpc-gateway_out=:proto proto/sqs/sqs.proto --openapiv2_out=:docs

proto_iam:
	protoc -I/usr/local/include  --proto_path=proto --go_out=plugins=grpc:proto \
	--grpc-gateway_out=:proto proto/iam/iam.proto --openapiv2_out=:docs

proto_res:
	protoc -I/usr/local/include  --proto_path=proto --go_out=plugins=grpc:proto \
	--grpc-gateway_out=:proto proto/responses/responses.proto --openapiv2_out=:docs
gen:
	protoc --proto_path=proto proto/**/*.proto  --go_out=:pb --go-grpc_out=:pb --grpc-gateway_out=:pb --openapiv2_out=:swagger

python:
	python3 -m grpc_tools.protoc  -I ./proto --python_out=./proto --grpc_python_out=./proto ./proto/lambda.proto