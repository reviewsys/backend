# build the server binary
FROM golang:1.12.4-alpine AS builder
LABEL stage=server-intermediate
ENV GO111MODULE=off
RUN apk add git \
	    curl \
	    make \
	    protobuf \
	    protobuf-dev

WORKDIR /go/src/github.com/reviewsys/backend
COPY . .

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure
RUN dep ensure -add github.com/golang/protobuf/protoc-gen-go
RUN dep ensure -add github.com/infobloxopen/protoc-gen-gorm
RUN cp -r vendor/* ${GOPATH}/src/

RUN go build
RUN go install

RUN cd $GOPATH/src/github.com/infobloxopen/protoc-gen-gorm \
    && make vendor \
    && make install

WORKDIR /go/src/github.com/reviewsys/backend
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/backend ./main.go

# copy the server binary from builder stage; run the server binary
FROM alpine:latest AS runner
WORKDIR /bin
COPY --from=builder /go/src/github.com/reviewsys/backend/bin/backend .


ENTRYPOINT ["backend"]
