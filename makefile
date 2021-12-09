all: pb xq/xq

pb: pb/games/xq/xq.proto
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pb/pet/pet.proto pb/games/xq/xq.proto

xq/xq:
	$(MAKE) -C xq
