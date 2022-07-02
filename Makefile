.PHONY: build install

build:
	@go build -o gss

install:
	@go install

build-in-ctn: # Build the server in a container
	@podman run -v "${PWD}":/go/src/gss -w /go/src/gss docker.io/golang:1.18 go build
