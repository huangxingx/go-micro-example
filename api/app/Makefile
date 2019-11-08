
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:

	go build -o app-web main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t app-web:latest


.PHONY: run
run:

	go build -o app-web main.go && ./app-web --registry=consul --registry_address=127.0.0.1:8500