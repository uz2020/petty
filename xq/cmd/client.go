package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/uz2020/petty/xq/client"
)

var clientCmd = &cobra.Command{
	Use:   "cli",
	Short: "Start client",
	Run: func(cmd *cobra.Command, args []string) {
		ctxWithCancel, ctxCancel := context.WithCancel(context.Background())
		cli := client.NewClient(ctxWithCancel)

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		fmt.Println("Stop Client ...")
		ctxCancel()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := cli.Shutdown(ctx); err != nil {
			panic(err)
		}
		fmt.Println("Stop Client done.")
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
