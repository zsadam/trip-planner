FROM golang:1.21-alpine AS base
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

FROM base AS dev
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
EXPOSE 2345

ENTRYPOINT ["air"]

FROM base AS builder
WORKDIR /app

COPY . /app
RUN go mod tidy && go mod vendor && go mod verify

RUN go build -o trip-planner -a .

FROM alpine:latest

COPY --from=builder /app/trip-planner /usr/local/bin/trip-planner
EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/trip-planner"]
