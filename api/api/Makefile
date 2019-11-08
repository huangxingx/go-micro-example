
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:

	go build -o app-web main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t api-web:latest


.PHONY: run
run:

	go build -o api-web main.go && ./api-web --registry=consul --registry_address=127.0.0.1:8500