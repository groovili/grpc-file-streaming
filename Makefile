gen-proto:
	protoc ./proto/streaming.proto --go_out=plugins=grpc:. ./proto/*.proto

build:
	go build -o ./.bin/server ./server/server.go
	go build -o ./.bin/client ./client/client.go

build-all: gen-proto build