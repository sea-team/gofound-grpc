API_PROTO_PATH=./api
GO_OUT_PATH=$(API_PROTO_PATH)/gen/v1
PROTO_FILE=$(API_PROTO_PATH)/gofound.proto
GATEWAY_PROTO_YAML=$(API_PROTO_PATH)/gofound.yaml

.PHONY: init
init:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go mod tidy

.PHONY: api
api:
	mkdir -p $(GO_OUT_PATH)
	protoc -I=$(API_PROTO_PATH)  \
	--go_out=$(GO_OUT_PATH) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(GO_OUT_PATH) \
	--go-grpc_opt=paths=source_relative $(PROTO_FILE)


.PHONY: gateway
gateway:	
	 protoc -I=$(API_PROTO_PATH) \
	 --grpc-gateway_out=paths=source_relative,grpc_api_configuration=$(GATEWAY_PROTO_YAML):$(GO_OUT_PATH) $(PROTO_FILE)


.PHONY: build
build:	
	docker stop gofound
	docker rm gofound
	docker rmi gofound:v1
	docker build -t gofound:v1 -f ./Dockerfile .
	docker run --name gofound -p 4567:4567 -p 5678:5678 -d gofound:v1

.PHONY: all
all: init api gateway build