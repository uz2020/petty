package game

import (
	"context"
	"fmt"

	pb "github.com/uz2020/petty/pb/games/xq"
)

func (gs *GameServer) MakeFriend(ctx context.Context, in *pb.MakeFriendRequest) (*pb.MakeFriendResponse, error) {
	out := &pb.MakeFriendResponse{}
	player := &Player{}

	if err := gs.auth(ctx, player); err != nil {
		return nil, err
	}

	userId := in.UserId
	user, err := GetUser(gs.dbConn, userId)
	if err != nil {
		return nil, err
	}

	fmt.Println("user id", user.UserId)

	return out, nil
}
