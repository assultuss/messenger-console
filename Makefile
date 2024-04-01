LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate:
	make generate-user-api

generate-user-api:
	mkdir -p pkg/
	protoc --proto_path=/path/to/your/proto/package/ \
	--go_out=pkg/ --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=/path/to/your/proto/package \
	--go-grpc_out=pkg/ --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=/path/to/your/proto/package \
	/path/to/your/proto/package//user.proto