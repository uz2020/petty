all: pb server client

pb:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pb/pet/pet.proto pb/games/xq/xq.proto

server: server/main.go
	$(MAKE) -C server all


client: client/main.go
	$(MAKE) -C client all
