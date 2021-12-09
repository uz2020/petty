all: pb xq/xq

pb: pb/games/xq/xq.proto
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pb/pet/pet.proto pb/games/xq/xq.proto

xq/xq: xq/game/game_server.go xq/client/client.go xq/main.go
	$(MAKE) -C xq
