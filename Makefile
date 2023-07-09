proto:
	protoc --go_out=plugins=grpc:. jwt-protobuf/protos/*.proto
