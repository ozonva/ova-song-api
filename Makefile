.PHONY: build
build: generate .build

.PHONY: all
all: dependencies build

.PHONY: run
run:
	go run ./cmd/ova-song-api

.PHONY: test
test: mock
	go test -v ./...

PHONY: mock
mock:
	go generate ./...

PHONY: generate
generate:
	protoc --proto_path=api --go_out=./pkg/ova-song-api --go_opt=paths=source_relative   \
			--go-grpc_out=./pkg/ova-song-api --go-grpc_opt=paths=source_relative   \
			ova-song-api.proto

PHONY: .build
.build:
ifeq ($(OS), Windows_NT)
	go build -o bin/ova-song-api.exe cmd/ova-song-api/main.go
else
	go build -o bin/ova-song-api cmd/ova-song-api/main.go
endif

.PHONY: dependencies
dependencies:
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go mod download

.PHONY: tidy
tidy:
	go mod tidy
