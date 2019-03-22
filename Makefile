generate:
	@go generate ./...
.PHONY: generate

install:
	@go get -v github.com/myitcv/gobin
	@gobin github.com/golang/protobuf/protoc-gen-go@v1.2.0
.PHONY: install

build:
	./build.sh
.PHONY: build

test:
	@go test ./...
.PHONY: test

test.cov:
	@go test ./... -coverprofile=coverage.txt -covermode=atomic
.PHONY: test.cov
