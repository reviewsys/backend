# vi: ft=make

GOPATH:=$(shell go env GOPATH)

.PHONY: proto test

proto:
	go get -u github.com/golang/protobuf/protoc-gen-go
	protoc -I . \
		-I vendor \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway \
		--go_out=plugins=grpc:. \
		--gorm_out=engine=postgres:. \
		app/delivery/grpc/*/*.proto

build: proto
	go build -o build/backend main.go
    
test:
	@go get -u github.com/rakyll/gotest
	gotest -p 1 -v ./...
