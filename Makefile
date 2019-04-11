# vi: ft=make

GOPATH:=$(shell go env GOPATH)

.PHONY: proto test

proto:
	go get github.com/golang/protobuf/protoc-gen-go
	protoc -I . app/delivery/grpc/*/*.proto --go_out=plugins=grpc:.

build: proto
	go build -o build/backend main.go
    
test:
	@go get github.com/rakyll/gotest
	gotest -p 1 -v ./...
