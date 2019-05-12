# vi: ft=make

GOPATH:=$(shell go env GOPATH)

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
	go build -o build/backend main.go
    
test:
	@go get -u github.com/rakyll/gotest
	gotest -p 1 -v ./...
