.PHONY: build
build:
	go build -v -ldflags="-X 'ova_route_api/build.User=$(shell id -u -n)' -X 'ova_route_api/build.Time=$(shell date)'" ./cmd/ova-route-api

.PHONY: test
test:
	mockgen -source=internal/repository/repository.go -destination=internal/repository/mocks/repository_mock.go -package mocks
	go test -v -race -timeout 30s ./...

.PHONY: lint
lint:
	gofumpt -l -w .
	golangci-lint run --fix	
	
.DEFAULT_GOAL := build
