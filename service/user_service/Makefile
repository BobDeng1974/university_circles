GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker

proto:
	protoc --proto_path=../../protos --go_out=pb/ --micro_out=pb/  ../../protos/user/*.proto
	protoc --proto_path=../../protos --go_out=../../api/pb/ --micro_out=../../api/pb/  ../../protos/user/*.proto

build: proto
	GOOS=linux GOARCH=amd64 go build -o user_services main.go plugin.go

test:
	go test -v ./... -cover

docker:
	docker build . -t user_services:latest