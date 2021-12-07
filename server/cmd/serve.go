package cmd

import (
	"fmt"
	"log"
	"net"

	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc"

	pb "github.com/uz2020/petty/pb/games/xq"

	"github.com/uz2020/petty/server/game"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	SERVICE  = "pet/games/xq"
	ETCD_URL = "http://localhost:2379"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start app",
	Run: func(cmd *cobra.Command, args []string) {
		if viper.Get("LISTEN_ADDR") == nil {
			panic("LISTEN_ADDR env key not found")
			return
		}

		addr := fmt.Sprintf("%v", viper.Get("LISTEN_ADDR"))
		initApp(addr)
	},
}

func registerService(addr string) {
	cli, err := clientv3.NewFromURL(ETCD_URL)
	if err != nil {
		log.Fatalf("failed to connect client: %v", err)
	}
	em, err := endpoints.NewManager(cli, SERVICE)
	if err != nil {
		log.Fatalf("failed to create manager: %v", err)
	}
	err = em.AddEndpoint(cli.Ctx(), fmt.Sprintf("%s/%s", SERVICE, addr), endpoints.Endpoint{Addr: addr})
	if err != nil {
		log.Fatalf("add endpoint failed: %v", err)
	}

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGameServer(s, &game.GameServer{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initApp(addr string) {
	fmt.Println("Run app on", addr)
	registerService(addr)
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
