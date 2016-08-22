CWD=$(shell pwd)
VENDORGOPATH := $(CWD)/vendor:$(CWD)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep
	if test -d src/github.com/the-pen/go-pen-gateway; then rm -rf src/github.com/the-pen/go-pen-gateway; fi
	mkdir -p src/github.com/the-pen/go-pen-gateway
	cp pen-gateway.go src/github.com/the-pen/go-pen-gateway/

rmdeps:
	if test -d src; then rm -rf src; fi

build:	rmdeps deps bin

deps:
	@GOPATH=$(shell pwd) go get -u "github.com/davecgh/go-spew/spew"
	@GOPATH=$(shell pwd) go get -u "github.com/tarm/serial"
	@GOPATH=$(shell pwd) go get -u "github.com/mikepb/go-serial"

fmt:
	go fmt cmd/*.go
	go fmt *.go

bin: 	self fmt
	@GOPATH=$(GOPATH) go build -o bin/list-ports cmd/list-ports.go
	@GOPATH=$(GOPATH) go build -o bin/ack cmd/ack.go
	@GOPATH=$(GOPATH) go build -o bin/gateway-id cmd/gateway-id.go
