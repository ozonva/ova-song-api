.PHONY: build
build:	generate .build

.PHONY: all
all: dependencies build

.PHONY: run
run:
	go run ./cmd/ova-song-api

.PHONY: run-lint
run-lint: lint

.PHONY: lint
lint:
	 golangci-lint run

.PHONY: tests
tests: mocks
	go test -v ./...

.PHONY: test-coverage
test-coverage: mocks
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

PHONY: mocks
mocks:
	go generate ./...

PHONY: generate
generate:
	protoc --proto_path=api --go_out=./pkg/ova-song-api --go_opt=paths=source_relative   \
			--go-grpc_out=./pkg/ova-song-api --go-grpc_opt=paths=source_relative   \
			ova-song-api.proto

	protoc --proto_path=api --go_out=./pkg/health-probe --go_opt=paths=source_relative   \
			--go-grpc_out=./pkg/health-probe --go-grpc_opt=paths=source_relative   \
			health-probe.proto

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

.PHONY: migrations
migrations: .migrations-deps .do-migrations

.PHONY: .migrations-deps
.migrations-deps:
		go get -u github.com/pressly/goose/v3/cmd/goose

.PHONY: .do-migrations
.do-migrations:
		goose -dir ./migrations postgres \
					"user=ova_song_db_user password=ova_song_db_user dbname=ova_song_db sslmode=disable" up

.PHONY: run-grpcui
run-grpcui:
	grpcui -plaintext -proto "./api/ova-song-api.proto" localhost:50051

.PHONY: run-grpcui-health
run-grpcui-health:
	grpcui -plaintext -proto "./api/health-probe.proto" localhost:50051

