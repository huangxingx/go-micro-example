
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/adminUser/adminUser.proto

.PHONY: build
build: proto

	go build -o adminUser-srv main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t adminUser-srv:latest
