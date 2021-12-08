all: pb xq client

pb: pb/games/xq/xq.proto
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pb/pet/pet.proto pb/games/xq/xq.proto

xq: xq/main.go
	$(MAKE) -C server all


client: client/main.go
	$(MAKE) -C client all
