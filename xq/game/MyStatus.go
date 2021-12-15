package game

import (
	"fmt"
	"time"

	pb "github.com/uz2020/petty/pb/games/xq"
	"github.com/uz2020/petty/xq/utils"
)

// 建立流推送
func (gs *GameServer) MyStatus(r *pb.MyStatusRequest, srv pb.Game_MyStatusServer) error {
	ctx := srv.Context()
	player := &Player{}

	isGuest := false
	if err := gs.auth(ctx, player); err != nil {
		isGuest = true
	}

	userId := utils.GenGuestUserId()
	username := userId
	if !isGuest {
		userId = player.user.UserId
		username = player.user.Username
	}
	gs.playerSrvs[userId] = srv

	out := pb.GetMyProfileResponse{
		Player: &pb.Player{
			User: &pb.User{
				UserId:   userId,
				Username: username,
			},
		},
	}

	for {
		err := srv.SendMsg(&out)
		if err != nil {
			fmt.Println("send msg err", err)
			break
		}
		fmt.Println("send msg")
		time.Sleep(time.Second)
	}

	return nil
}
