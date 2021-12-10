package client

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/spf13/viper"
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

func pf(format string, a ...interface{}) (n int, err error) {
	fmt.Printf("\t")
	return fmt.Printf(format+"\n", a...)
}

type Client struct {
	ctx    context.Context
	conf   config.Conf
	ecli   *ecv3.Client
	gc     pb.GameClient
	creds  cred
	player *pb.Player
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
	ActionTypeLogout
	ActionTypeCreateTable
	ActionTypeJoinTable
	ActionTypeLeaveTable
	ActionTypeStartGame
	ActionTypeMove
	ActionTypeGetTables
	ActionTypeStatus
	ActionTypeMyProfile
)

var actionTypes = []string{
	ActionTypeRegister:    "register",
	ActionTypeLogin:       "login",
	ActionTypeLogout:      "logout",
	ActionTypeCreateTable: "create table",
	ActionTypeJoinTable:   "join table",
	ActionTypeLeaveTable:  "leave table",
	ActionTypeStartGame:   "start game",
	ActionTypeMove:        "move",
	ActionTypeGetTables:   "get tables",
	ActionTypeStatus:      "status",
	ActionTypeMyProfile:   "my profile",
}

var ActionPrompts = []ActionPrompt{
	// register
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
	// login
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
	// logout
	{},
	// create table
	{
		prompts: []Prompter{
			Prompt{
				promptui.Prompt{
					Label: "Name",
				},
			},
		},
	},
	// join table
	{
		prompts: []Prompter{
			Prompt{
				promptui.Prompt{
					Label: "Table Id",
				},
			},
		},
	},
	// leave table
	{},
	// start game
	{},
	// move
	{},
	// get tables
	{},
	// status
	{},
	// my profile
	{},
}

func NewClient(ctx context.Context) *Client {
	cli := &Client{ctx: ctx}
	cli.conf.Init()
	return cli
}

func login(cli *Client, argv []string) {
	name := argv[0]
	passwd := argv[1]

	resp, err := cli.gc.Login(cli.ctx, &pb.LoginRequest{
		Username: name,
		Passwd:   passwd,
	})

	if err != nil {
		pl("login err", err)
		return
	}

	token := resp.Token
	viper.Set("token", token)
	pl("login success", token)
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
	name := argv[0]

	resp, err := cli.gc.CreateTable(cli.ctx, &pb.CreateTableRequest{
		Name: name,
	})

	if err != nil {
		pl("create table err", err)
		return
	}

	pf("create table success, table id %s", resp.TableId)
}

func logout(cli *Client, argv []string) {
	_, err := cli.gc.Logout(cli.ctx, &pb.LogoutRequest{})
	if err != nil {
		pl("logout err", err)
		return
	}
	pl("logout success")
	viper.Set("token", "")
	cli.player = nil
}

func joinTable(cli *Client, argv []string) {
	tableId := argv[0]

	_, err := cli.gc.JoinTable(cli.ctx, &pb.JoinTableRequest{
		TableId: tableId,
	})

	if err != nil {
		pl("join table err", err)
		return
	}

	pf("join table success")
}

func myProfile(cli *Client, argv []string) {
	if cli.player == nil {
		pl("********* please login first ********")
		return
	}
	pl("------------- me -----------")
	pl("username:", cli.player.User.Username)
	pl("user_id:", cli.player.User.UserId)
	pl("----------------------------")
}

func (cli *Client) handleCmd(act Action) {
	cmd := act.cmd
	argv := act.argv

	switch cmd {
	case ActionTypeMyProfile:
		myProfile(cli, argv)
	case ActionTypeRegister:
		register(cli, argv)
	case ActionTypeCreateTable:
		createTable(cli, argv)
	case ActionTypeJoinTable:
		joinTable(cli, argv)
	case ActionTypeLeaveTable:
	case ActionTypeStartGame:
	case ActionTypeMove:
	case ActionTypeLogin:
		login(cli, argv)
	case ActionTypeLogout:
		logout(cli, argv)
	case ActionTypeGetTables:
		reply, err := cli.gc.GetTables(cli.ctx, &pb.TablesRequest{})
		if err != nil {
			fmt.Println("[error]", err)
			return
		}

		for _, tb := range reply.Tables {
			pf("table id: %s\tname: %s\towner:\t%s", tb.TableId, tb.Name, tb.Owner.Username)
		}

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

type cred struct{}

func (c *cred) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"token": viper.GetString("token"),
	}, nil
}

func (*cred) RequireTransportSecurity() bool {
	return false
}

func (cli *Client) Run() {
	configPath := "cli-config.yaml"
	viper.SetConfigName(configPath)
	viper.AddConfigPath("./")
	viper.ReadInConfig()
	defer func() {
		err := viper.WriteConfigAs(configPath)
		if err != nil {
			pf("write config failed %s", err)
		}
	}()

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
		Label:    "Select Action",
		Items:    actionTypes,
		HideHelp: true,
	}

	continuePrompt := promptui.Prompt{
		Label: "continue? (Y/N [default: Y])",
	}

	for {
		if viper.GetString("token") != "" && cli.player == nil {
			resp, err := cli.gc.GetMyProfile(cli.ctx, &pb.GetMyProfileRequest{})
			if err != nil {
				pf("get profile failed %v", err)
			}
			cli.player = resp.Player

			myProfile(cli, []string{})
		}

		argv := []string{}
		i, result, err := prompt.Run()

		if err != nil {
			pf("Prompt failed %v", err)
			return
		}
		cmd := i
		pf("You chose %q\n", result)
		for _, p := range ActionPrompts[i].prompts {
			_, result, err := p.Run()
			if err != nil {
				pf("Prompt failed %v", err)
				return
			}
			argv = append(argv, result)
		}
		cli.handleCmd(Action{
			cmd:  cmd,
			argv: argv,
		})

		pl()
		pl()
		result, err = continuePrompt.Run()
		if err != nil {
			pf("Prompt failed %v", err)
			return
		}
		if result != "" && result != "Y" && result != "y" {
			pl("exited...")
			break
		}
	}
}

func (*Client) Shutdown(ctx context.Context) error {
	return nil
}
