
.PHONY: proto
proto:
	@protoc --proto_path=. -I${GOPATH}/src --go_out=. --go-grpc_out=. --goclient_out=svc=myservice-grpc,grpc-option=test/grpc:. test.proto
