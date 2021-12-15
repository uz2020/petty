package client

import (
	"github.com/manifoldco/promptui"
	pb "github.com/uz2020/petty/pb/games/xq"
)

func makeFriend(cli *Client, argv []string) {
	userId := argv[0]
	_, err := cli.gc.MakeFriend(cli.ctx, &pb.MakeFriendRequest{UserId: userId})

	if err != nil {
		pl("err", err)
		return
	}

	pf("make friend success")
}

func init() {
	actionPrompts = append(actionPrompts, ActionPrompt{
		name: "make friend",
		prompts: []Prompter{
			Prompt{
				promptui.Prompt{
					Label: "Player User Id",
				},
			},
		},
		f: makeFriend,
	})
}
