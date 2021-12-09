package game

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc"

	pb "github.com/uz2020/petty/pb/games/xq"
	"github.com/uz2020/petty/xq/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

type GameServer struct {
	pb.UnimplementedGameServer
	conf config.Conf
	ctx  context.Context
	db   *gorm.DB
}

func (gs *GameServer) init(ctx context.Context) {
	gs.ctx = ctx
	gs.conf.Init()
	conf := &gs.conf
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=3s",
		conf.MysqlUser, conf.MysqlPasswd, conf.MysqlAddr, conf.MysqlAddr, conf.MysqlDb)

	defaultLogger := logger.Default.LogMode(logger.Silent)
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   defaultLogger,
		SkipDefaultTransaction:                   true,
	}

	if db, err := gorm.Open(mysql.Open(dsn), gormConfig); err == nil {
		err = db.Use(dbresolver.Register(dbresolver.Config{}).
			SetConnMaxLifetime(1 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(1000))
		if err != nil {
			panic(err)
		}
		gs.db = db
	} else {
		panic(err)
	}

	registerService(conf.ListenAddr, conf.Service, conf.EtcdUrl, gs)
}

func (*GameServer) GetTables(ctx context.Context, request *pb.TablesRequest) (*pb.TablesReply, error) {
	r := &pb.TablesReply{
		TableIds: []int64{33, 44, 55},
	}
	return r, nil
}

func (*GameServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return nil, nil
}

func (*GameServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}

func (*GameServer) MyStatus(r *pb.MyStatusRequest, srv pb.Game_MyStatusServer) error {
	ctx := srv.Context()

	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		now := time.Now().Unix()
		srv.Send(&pb.MyStatusResponse{Time: now})
		time.Sleep(time.Second)
	}
	return nil
}

func (*GameServer) JoinTable(ctx context.Context, in *pb.JoinTableRequest) (*pb.JoinTableResponse, error) {
	return nil, nil
}

func (*GameServer) LeaveTable(ctx context.Context, in *pb.LeaveTableRequest) (*pb.LeaveTableResponse, error) {
	return nil, nil
}

func (*GameServer) StartGame(ctx context.Context, in *pb.StartGameRequest) (*pb.StartGameResponse, error) {
	return nil, nil
}

func (*GameServer) Move(ctx context.Context, in *pb.MoveRequest) (*pb.MoveResponse, error) {
	return nil, nil
}

func (*GameServer) Shutdown(ctx context.Context) error {
	return nil
}

func registerService(addr, service, etcdUrl string, gs *GameServer) {
	// connect etcd
	cli, err := clientv3.NewFromURL(etcdUrl)
	if err != nil {
		log.Fatalf("failed to connect client: %v", err)
	}

	// create endpoint
	em, err := endpoints.NewManager(cli, service)
	if err != nil {
		log.Fatalf("failed to create manager: %v", err)
	}

	// add endpoint
	err = em.AddEndpoint(cli.Ctx(), fmt.Sprintf("%s/%s", service, addr), endpoints.Endpoint{Addr: addr})
	if err != nil {
		log.Fatalf("add endpoint failed: %v", err)
	}

	// serve
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGameServer(s, gs)

	go func() {
		if err := s.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}

func NewGameServer(ctx context.Context) *GameServer {
	gs := &GameServer{}
	gs.init(ctx)

	return gs
}
