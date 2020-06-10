GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker

proto:
	protoc --proto_path=../../protos --go_out=pb/ --micro_out=pb/  ../../protos/im/*.proto
	protoc --proto_path=../../protos --go_out=../../api/pb/ --micro_out=../../api/pb/  ../../protos/im/*.proto

build: proto
	GOOS=linux GOARCH=amd64 go build -o im_services main.go plugin.go

test:
	go test -v ./... -cover

docker:
	docker build . -t im_services:latest