generate_grpc_code:
	protoc \
	--go-out=pb \
	--go-opt=paths=source_relative \
	--go-grpc_out=pb \
	--go-grpc_opt=paths=source_relative \
	./proto/service.proto