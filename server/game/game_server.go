package game

import (
	"context"

	pb "github.com/uz2020/petty/pb/games/xq"
)

type GameServer struct {
	pb.UnimplementedGameServer
}

func (*GameServer) GetTables(ctx context.Context, request *pb.TablesRequest) (*pb.TablesReply, error) {
	return nil, nil
}
