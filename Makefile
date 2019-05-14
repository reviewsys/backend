# vi: ft=make

GOPATH:=$(shell go env GOPATH)
VERSION ?= $(shell git describe --tags --abbrev=0)
REVISION ?= $(shell git describe --always)
BUILD_DATE ?= $(shell date +'%Y-%m-%dT%H:%M:%SZ')
LDFLAGS := -ldflags "-X main.version=${VERSION} -X main.revision=$(REVISION) -X main.buildDate=$(BUILD_DATE)"

.PHONY: proto test

proto:
	go get -u github.com/golang/protobuf/protoc-gen-go
	protoc -I . \
		-I vendor \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway \
		--go_out=plugins=grpc:${GOPATH}/src \
		--gorm_out=engine=postgres:${GOPATH}/src \
		app/interface/rpc/v1.0/protocol/*.proto

client:
	protoc -I . \
		-I vendor \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway \
		--go_out=plugins=grpc:${GOPATH}/src \
		--gorm_out=engine=postgres:${GOPATH}/src \
		--cobra_out=plugins=client:${GOPATH}/src \
		app/interface/rpc/v1.0/protocol/*.proto


build: proto
	# go build $(LDFLAGS) -o bin/backend main.go
	go build -o bin/backend main.go
    
test:
	@go get -u github.com/rakyll/gotest
	gotest -p 1 -v ./...
