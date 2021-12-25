package client

import (
	"context"
	"fmt"
	"log"
	"time"

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
	ss     pb.Game_MyStatusClient
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
	name    string
	prompts []Prompter
	f       func(*Client, []string)
}

var actionPrompts = []ActionPrompt{
	{
		name: "register",
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
		f: register,
	},
	{
		name: "login",
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
		f: login,
	},
	{
		name: "logout",
		f:    logout,
	},
	{
		name: "create table",
		prompts: []Prompter{
			Prompt{
				promptui.Prompt{
					Label: "Name",
				},
			},
		},
		f: createTable,
	},
	{
		name: "join table",
		prompts: []Prompter{
			Prompt{
				promptui.Prompt{
					Label: "Table Id",
				},
			},
		},
		f: joinTable,
	},
	{
		name: "leave table",
		f:    actionHandlerStub,
	},
	{
		name: "start game",
		f:    actionHandlerStub,
	},
	{
		name: "move",
		f:    actionHandlerStub,
	},
	{
		name: "get tables",
		f:    getTables,
	},
	{
		name: "status",
		f:    statusStream,
	},
	{
		name: "my profile",
		f:    myProfile,
	},
	{
		name: "get player",
		prompts: []Prompter{
			Prompt{
				promptui.Prompt{
					Label: "Player User Id",
				},
			},
		},
		f: getPlayer,
	},
}

func NewClient(ctx context.Context) *Client {
	cli := &Client{ctx: ctx}
	cli.conf.Init()
	return cli
}

func actionHandlerStub(cli *Client, argv []string) {}

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

func getTables(cli *Client, argv []string) {
	tbs, err := cli.gc.GetTables(cli.ctx, &pb.TablesRequest{})
	if err != nil {
		pl("get tables err", err)
		return
	}

	for _, tb := range tbs.Tables {
		pf("name: %s, owner: %s, id: %s", tb.Name, tb.Owner, tb.TableId)
	}
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

func getPlayer(cli *Client, argv []string) {
	userId := argv[0]
	resp, err := cli.gc.GetPlayer(cli.ctx, &pb.GetPlayerRequest{UserId: userId})
	if err != nil {
		pl("get player err", err)
		return
	}

	pf("player user id: %s, username: %s", resp.Player.User.UserId, resp.Player.User.Username)
}

func (cli *Client) handleCmd(act Action) {
	cmd := act.cmd
	argv := act.argv

	for i, ap := range actionPrompts {
		if i == cmd {
			ap.f(cli, argv)
			break
		}
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
	configPath := "cli-config"
	viper.AddConfigPath("./")
	viper.SetConfigName(configPath)
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		pf("read in config failed: %v", err)
		err := viper.SafeWriteConfig()
		if err != nil {
			pf("write config failed %s", err)
			return
		}
		pf("config file created")
	}
	defer func() {
		err := viper.WriteConfig()
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

	actionTypes := []string{}
	for _, ap := range actionPrompts {
		actionTypes = append(actionTypes, ap.name)
	}

	prompt := promptui.Select{
		Label:    "Select Action",
		Items:    actionTypes,
		HideHelp: true,
	}

	continuePrompt := promptui.Prompt{
		Label: "continue? (Y/N [default: Y])",
	}

	for {
		token := viper.GetString("token")
		if token != "" && cli.player == nil {
			pf("token loaded %v, getting profile...", token)
			resp, err := cli.gc.GetMyProfile(cli.ctx, &pb.GetMyProfileRequest{})
			if err != nil {
				pf("get profile failed %v", err)
				time.Sleep(time.Second)
				continue
			}
			cli.player = resp.Player

			myProfile(cli, []string{})
			statusStream(cli, []string{})
		}

		argv := []string{}
		i, result, err := prompt.Run()

		if err != nil {
			pf("Prompt failed %v", err)
			return
		}
		cmd := i
		pf("You chose %q\n", result)
		for _, p := range actionPrompts[i].prompts {
			_, result, err := p.Run()
			if err != nil {
				pf("Prompt failed %v", err)
				return
			}
			argv = append(argv, result)
		}

		pl("\n\n")

		cli.handleCmd(Action{
			cmd:  cmd,
			argv: argv,
		})

		pl("\n\n")

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
