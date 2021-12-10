package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	pb "github.com/uz2020/petty/pb/games/xq"
	"github.com/uz2020/petty/xq/config"
	"google.golang.org/grpc"

	ecv3 "go.etcd.io/etcd/client/v3"
	resolver "go.etcd.io/etcd/client/v3/naming/resolver"

	"github.com/manifoldco/promptui"
)

func pl(a ...interface{}) (n int, err error) {
	fmt.Printf("\t")
	return fmt.Println(a...)
}

func pr(format string, a ...interface{}) (n int, err error) {
	fmt.Printf("\t")
	return fmt.Printf(format, a...)
}

type Client struct {
	ctx   context.Context
	conf  config.Conf
	ecli  *ecv3.Client
	gc    pb.GameClient
	creds cred
}

type Prompter interface {
	Run() (int, string, error)
}

type Prompt struct {
	promptui.Prompt
}

type Select struct {
	promptui.Select
}

func (p Prompt) Run() (int, string, error) {
	result, err := p.Prompt.Run()
	return 0, result, err
}

func (p Select) Run() (int, string, error) {
	return p.Select.Run()
}

type Action struct {
	cmd  int
	argv []string
}

type ActionPrompt struct {
	prompts []Prompter
}

const (
	ActionTypeRegister = iota
	ActionTypeLogin
	ActionTypeCreateTable
	ActionTypeJoinTable
	ActionTypeLeaveTable
	ActionTypeStartGame
	ActionTypeMove
	ActionTypeGetTables
	ActionTypeStatus
)

var actionTypes = []string{
	ActionTypeRegister:    "register",
	ActionTypeLogin:       "login",
	ActionTypeCreateTable: "create table",
	ActionTypeJoinTable:   "join table",
	ActionTypeLeaveTable:  "leave table",
	ActionTypeStartGame:   "start game",
	ActionTypeMove:        "move",
	ActionTypeGetTables:   "get tables",
	ActionTypeStatus:      "status",
}

var ActionPrompts = []ActionPrompt{
	{
		prompts: []Prompter{
			Prompt{
				promptui.Prompt{
					Label: "Username",
				},
			},
			Prompt{
				promptui.Prompt{
					Label: "Password",
				},
			},
		},
	},
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

	resp, err := cli.gc.Login(cli.ctx, &pb.LoginRequest{
		Username: name,
		Passwd:   passwd,
	})

	if err != nil {
		log.Println("login err", err)
		return
	}

	token := resp.Token
	cli.creds.token = token
	log.Println("login success", token)
}

func register(cli *Client, argv []string) {
	name := argv[0]
	passwd := argv[1]

	_, err := cli.gc.Register(cli.ctx, &pb.RegisterRequest{
		Username: name,
		Passwd:   passwd,
	})

	if err != nil {
		pl("register err", err)
		return
	}

	pl("register success")
}

func createTable(cli *Client, argv []string) {
	if len(argv) < 2 {
		return
	}
	name := argv[1]

	_, err := cli.gc.CreateTable(cli.ctx, &pb.CreateTableRequest{
		Name: name,
	})

	if err != nil {
		log.Println("create table err", err)
		return
	}

	log.Println("create table success")
}

func (cli *Client) handleCmd(act Action) {
	cmd := act.cmd
	argv := act.argv

	switch cmd {
	case ActionTypeRegister:
		register(cli, argv)
	case ActionTypeCreateTable:
		go createTable(cli, argv)
	case ActionTypeJoinTable:
	case ActionTypeLeaveTable:
	case ActionTypeStartGame:
	case ActionTypeMove:
	case ActionTypeLogin:
		go login(cli, argv)
	case ActionTypeGetTables:
		reply, err := cli.gc.GetTables(cli.ctx, &pb.TablesRequest{})
		if err != nil {
			fmt.Println("[error]", err)
			return
		}

		fmt.Println(reply)
	case ActionTypeStatus:
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

type cred struct {
	token string
}

func (c *cred) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"token": c.token,
	}, nil
}

func (*cred) RequireTransportSecurity() bool {
	return false
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

	conn, err := grpc.Dial(service, grpc.WithResolvers(etcdResolver), grpc.WithInsecure(), grpc.WithPerRPCCredentials(&cli.creds))
	if err != nil {
		log.Fatalf("dial failed: %v", err)
	}
	fmt.Printf("----------- dial success %v --------------\n", service)

	defer conn.Close()
	cli.gc = pb.NewGameClient(conn)

	prompt := promptui.Select{
		Label: "Select Action",
		Items: actionTypes,
	}

	continuePrompt := promptui.Prompt{
		Label: "continue? (Y/N [default: Y])",
	}

	for {
		argv := []string{}
		i, result, err := prompt.Run()

		if err != nil {
			pr("Prompt failed %v\n", err)
			return
		}
		cmd := i
		pr("You chose %q\n", result)
		for _, p := range ActionPrompts[i].prompts {
			_, result, err := p.Run()
			if err != nil {
				pr("Prompt failed %v\n", err)
				return
			}
			argv = append(argv, result)
		}
		cli.handleCmd(Action{
			cmd:  cmd,
			argv: argv,
		})

		result, err = continuePrompt.Run()
		if err != nil {
			pr("Prompt failed %v\n", err)
			return
		}
		if result != "" && result != "Y" && result != "y" {
			fmt.Println("exiting...")
			os.Exit(0)
		}
	}
}

func (*Client) Shutdown(ctx context.Context) error {
	return nil
}
