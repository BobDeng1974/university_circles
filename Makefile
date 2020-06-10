
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker


#proto:
#	find shopproto -name "*.proto" | xargs -t -I{} protoc -I.:${GOPATH}/src --gofast_out=plugins=micro:. {}
#	proto --proto_path=${GOPATH}/src:. --micro_out=. --gofast_out=. proto/*.proto

build:
	#cd service/home-service && make build
	#cd service/info-service && make build
	cd service/home_service && make build
	cd service/user_service && make build
	cd api && make build

test:
	go test -v ./... -cover

docker:
    docker build . -t home_services:latest
	docker build . -t user_services:latest
	docker build . -t api:latest

