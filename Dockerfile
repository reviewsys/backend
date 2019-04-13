FROM golang:alpine AS build-env
WORKDIR /usr/local/go/src/github.com/reviewsys/backend
COPY . /usr/local/go/src/github.com/reviewsys/backend
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN go build -o build/backend .


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /usr/local/go/src/github.com/reviewsys/backend/build/backend /bin/backend
CMD ["backend"]
