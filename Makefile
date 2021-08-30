.PHONY: build
build:
	go build -v -ldflags="-X 'ova_route_api/build.User=$(shell id -u -n)' -X 'ova_route_api/build.Time=$(shell date)'" ./cmd/ova-route-api

.PHONY: deps
deps:
	go get -u github.com/onsi/ginkgo
	go get -u github.com/onsi/gomega
	go get -u github.com/golang/mock
	go get -u github.com/rs/zerolog/log
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/rs/zerolog/log
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

.PHONY: generate
generate:
	protoc --proto_path=. \
	--go_out=pkg/api --go_opt=paths=import \
	--go-grpc_out=pkg/api --go-grpc_opt=paths=import \
	api/api.proto

.PHONY: test
test:
	mockgen -source=internal/repository/repository.go -destination=internal/repository/mocks/repository_mock.go -package mocks
	go test -v -race -timeout 30s ./...

.PHONY: lint
lint:
	gofumpt -l -w .
	golangci-lint run --fix	
	
.DEFAULT_GOAL := build
