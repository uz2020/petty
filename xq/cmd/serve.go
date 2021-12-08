package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/uz2020/petty/xq/game"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start app",
	Run: func(cmd *cobra.Command, args []string) {
		if viper.Get("LISTEN_ADDR") == nil {
			panic("LISTEN_ADDR env key not found")
		}
		addr := fmt.Sprintf("%v", viper.Get("LISTEN_ADDR"))

		initApp(addr)
	},
}

func initApp(addr string) {
	fmt.Println("Run app on", addr)
	ctxWithCancel, ctxCancel := context.WithCancel(context.Background())

	gs := game.NewGameServer(ctxWithCancel)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Stop App ...")
	ctxCancel()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := gs.Shutdown(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Stop App done.")
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
