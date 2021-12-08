package client

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	pb "github.com/uz2020/petty/pb/games/xq"
	"github.com/uz2020/petty/xq/config"
	"google.golang.org/grpc"

	ecv3 "go.etcd.io/etcd/client/v3"
	resolver "go.etcd.io/etcd/client/v3/naming/resolver"
)

type Client struct {
	ctx  context.Context
	conf config.Conf
	ecli *ecv3.Client
	gc   pb.GameClient
}

func NewClient(ctx context.Context) *Client {
	cli := &Client{ctx: ctx}
	cli.conf.Init()
	go cli.Run()
	return cli
}

func login(cli *Client, argv []string) {
	if len(argv) < 3 {
		return
	}
	name := argv[1]
	passwd := argv[2]

	stream, err := cli.gc.Login(cli.ctx, &pb.LoginRequest{
		Username: name,
		Passwd:   passwd,
	})

	if err != nil {
		log.Fatalln("login err", err)
		return
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Printf("stream err %v", err)
				return
			}
			fmt.Println(resp.Time)
		}
	}()

	<-done
	log.Println("finished")
}

func (cli *Client) handleCmd(line string) {
	argv := strings.Fields(line)

	if len(argv) == 0 {
		return
	}

	cmd := argv[0]

	switch cmd {
	case "login":
		go login(cli, argv)
	case "tables":
		reply, err := cli.gc.GetTables(cli.ctx, &pb.TablesRequest{})
		if err != nil {
			fmt.Println("[error]", err)
			return
		}

		fmt.Println(reply)
	case "status":
		stream, err := cli.gc.MyStatus(cli.ctx, &pb.MyStatusRequest{})
		if err != nil {
			log.Fatalf("[error] %v", err)
			return
		}

		ctx := stream.Context()
		done := make(chan bool)

		go func() {
			for {
				resp, err := stream.Recv()
				if err == io.EOF {
					close(done)
					return
				}
				if err != nil {
					log.Fatalf("can not receive %v", err)
				}
				ts := resp.Time
				log.Printf("new timestamp %d received", ts)
			}
		}()

		go func() {
			<-ctx.Done()
			if err := ctx.Err(); err != nil {
				log.Println(err)
			}
			close(done)
		}()

		<-done
	}
}

func (cli *Client) Run() {
	ecli, err := ecv3.NewFromURL(cli.conf.EtcdUrl)
	if err != nil {
		log.Fatalf("new client failed: %v", err)
	}

	etcdResolver, err := resolver.NewBuilder(ecli)
	if err != nil {
		log.Fatalf("new builder failed: %v", err)
	}

	cli.ecli = ecli

	service := "etcd:///" + cli.conf.Service
	conn, err := grpc.Dial(service, grpc.WithResolvers(etcdResolver), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial failed: %v", err)
	}
	fmt.Printf("dial success %v\n", service)

	defer conn.Close()
	cli.gc = pb.NewGameClient(conn)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("xq client")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.ReplaceAll(text, "\n", "")

		cli.handleCmd(text)
	}
}

func (*Client) Shutdown(ctx context.Context) error {
	return nil
}
