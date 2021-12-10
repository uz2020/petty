package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/uz2020/petty/xq/client"
)

var clientCmd = &cobra.Command{
	Use:   "cli",
	Short: "Start client",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		cli := client.NewClient(ctx)
		cli.Run()
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
