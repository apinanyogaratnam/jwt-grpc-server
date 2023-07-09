proto:
	protoc --go_out=plugins=grpc:jwt jwt-protobuf/protos/*.proto

start:
	go run main.go

update-submodules:
	git submodule update --remote --merge
