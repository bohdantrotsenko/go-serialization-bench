proto-gen:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. --proto_path=. -I=. proto.proto 
