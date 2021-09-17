# works when protoc is on path (execute from command line)
.PHONY: compile
compile: ## Compile the proto file.
	protoc --go_opt=paths=source_relative -I pkg/proto/ pkg/proto/credit.proto --go_out=plugins=grpc:pkg/proto/gen
 
.PHONY: server
server: ## Build and run server.
	go build -race -ldflags "-s -w" -o bin/server server/main.go
	bin/server
 
.PHONY: client
client: ## Build and run client.
	go build -race -ldflags "-s -w" -o bin/client client/main.go
	bin/client
