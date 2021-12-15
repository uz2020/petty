package game

import (
	"context"
	"fmt"
	"time"

	pb "github.com/uz2020/petty/pb/games/xq"
	"github.com/uz2020/petty/xq/db"
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

	o := db.TbMakeFriend{
		FromUserId: player.user.UserId,
		ToUserId:   user.UserId,
		CreatedAt:  time.Now(),
	}

	res := gs.dbConn.Create(&o)
	return out, res.Error
}
