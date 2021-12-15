package client

import (
	"github.com/manifoldco/promptui"
)

func makeFriend(cli *Client, argv []string) {
	pf("make friend")
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
