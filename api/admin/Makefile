
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:

	go build -o admin-web main.go plugin.go

.PHONY: run
run:

	go build -o admin-web main.go  && ./admin-web --registry=consul --registry_address=127.0.0.1:8500

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t admin-web:latest
