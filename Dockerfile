FROM golang:1.18-alpine

ENV CGO_ENABLED=0
RUN mkdir -p /go/src/github.com/advalistar/go-shopify
WORKDIR /go/src/github.com/advalistar/go-shopify
