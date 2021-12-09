package game

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	pb "github.com/uz2020/petty/pb/games/xq"
	"github.com/uz2020/petty/xq/config"

	"github.com/uz2020/petty/xq/db"
	"github.com/uz2020/petty/xq/utils"
	"gorm.io/gorm"

	"github.com/go-redis/redis/v8"
)

type GameServer struct {
	pb.UnimplementedGameServer
	conf      config.Conf
	ctx       context.Context
	dbConn    *gorm.DB
	redisConn *redis.Client
}

func (gs *GameServer) init(ctx context.Context) {
	gs.ctx = ctx
	gs.conf.Init()
	conf := &gs.conf
	dbConn, err := db.InitDb(conf.MysqlUser, conf.MysqlPasswd, conf.MysqlAddr, conf.MysqlDb)
	if err != nil {
		panic(err)
	}
	gs.dbConn = dbConn
	gs.redisConn = db.InitRedis(conf.RedisAddr)
	registerService(conf.ListenAddr, conf.Service, conf.EtcdUrl, gs)
}

func (gs *GameServer) auth(ctx context.Context, user *db.TbUser) error {
	var userId, token string
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return grpc.Errorf(codes.Unauthenticated, "missing credentials 1")
	}

	val, ok := md["token"]
	if !ok {
		return grpc.Errorf(codes.Unauthenticated, "missing credentials 2")
	}

	token = val[0]
	key := fmt.Sprintf("user_sid_%s", token)
	userId, err := gs.redisConn.Get(ctx, key).Result()
	if err != nil {
		return grpc.Errorf(codes.Unauthenticated, "invalid token")
	}

	result := gs.dbConn.First(user, "user_id = ?", userId)
	err = result.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return grpc.Errorf(codes.NotFound, "record not found")
		}
		return grpc.Errorf(codes.Internal, "db error")
	}

	return nil
}

func (gs *GameServer) GetTables(ctx context.Context, in *pb.TablesRequest) (*pb.TablesReply, error) {
	user := db.TbUser{}
	out := &pb.TablesReply{}

	if err := gs.auth(ctx, &user); err != nil {
		return nil, err
	}

	log.Printf("user: %v", user)
	return out, nil
}

func (gs *GameServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	out := &pb.RegisterResponse{}
	user := db.TbUser{}
	result := gs.dbConn.First(&user, "username = ?", in.Username)
	err := result.Error
	if err == nil {
		return nil, errors.New("username already registered, choose another one")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	user.Username = in.Username
	user.Password = in.Passwd
	user.CreatedAt = time.Now()
	user.UserId = utils.GenUserId()
	result = gs.dbConn.Create(&user)
	err = result.Error
	return out, err
}

func (gs *GameServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	out := &pb.LoginResponse{}
	user := db.TbUser{}
	result := gs.dbConn.First(&user, "username = ? and password = ?", in.Username, in.Passwd)
	err := result.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("username or password incorrect")
		}
		return nil, err
	}
	token, err := utils.NewRandomString()

	if err != nil {
		return nil, err
	}
	out.Token = *token

	// 设置redis cookie
	err = gs.redisConn.Set(ctx, fmt.Sprintf("user_sid_%s", out.Token), user.UserId, time.Hour*24*7).Err()
	if err != nil {
		return nil, err
	}
	return out, nil
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
