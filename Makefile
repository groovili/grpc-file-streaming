gen-streaming:
	protoc ./proto/streaming.proto --go_out=plugins=grpc:. ./proto/*.proto
gen: gen-streaming