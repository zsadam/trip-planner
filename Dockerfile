FROM golang:1.21

WORKDIR /app

COPY . /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin latest

RUN go mod tidy && go mod vendor && go mod verify
