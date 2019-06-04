# vi: ft=make

GOPATH:=$(shell go env GOPATH)
VERSION ?= $(shell git describe --tags --abbrev=0)
REVISION ?= $(shell git describe --always)
BUILD_DATE ?= $(shell date '+%FT%T%z')
LDFLAGSPATH := github.com/reviewsys/backend/app/interface/persistence/memory
LDFLAGS := -ldflags "-X ${LDFLAGSPATH}.Version=${VERSION} -X ${LDFLAGSPATH}.Revision=$(REVISION) -X ${LDFLAGSPATH}.BuildDate=$(BUILD_DATE)"

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

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
	go build $(LDFLAGS) -o bin/backend main.go
    
test:
	@go get -u github.com/rakyll/gotest
	gotest -p 1 -v ./...

fmt:
	@gofmt -l -w $(SRC)

simplify:
	@gofmt -s -l -w $(SRC)
