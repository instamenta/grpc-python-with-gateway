OUTPUT_DIR_PY := ./
OUTPUT_DIR_GO := ./proxy/chat
OUTPUT_DIR_OPENAPI := ./proxy/third_party/OpenAPI/chat/v1

run:
	python ./__main__.py


all: generate-python generate-grpc-gateway

generate-python:
	@python -m grpc_tools.protoc \
    	-I . \
    	--pyi_out=$(OUTPUT_DIR_PY) \
    	--python_out=$(OUTPUT_DIR_PY) \
    	--grpc_python_out=$(OUTPUT_DIR_PY) \
    	chat.proto

generate-grpc-gateway:
	@protoc -I . \
        -I $(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.19.0 \
   		--go_out=$(OUTPUT_DIR_GO) \
   		--go_opt=paths=source_relative \
   		--go-grpc_out=$(OUTPUT_DIR_GO) \
   		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=$(OUTPUT_DIR_GO) \
		--grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=logtostderr=true:$(OUTPUT_DIR_OPENAPI) \
   		chat.proto
