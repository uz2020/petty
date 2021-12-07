package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	pb "t/pet"

	"go.etcd.io/etcd/client/v3"

	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

const (
	service = "foo/bar/my-service"
)

type profileServer struct {
	pb.UnimplementedProfileServer
}

func (*profileServer) GetProfile(ctx context.Context, request *pb.ProfileRequest) (*pb.ProfileReply, error) {
	log.Println("get profile")
	reply := &pb.ProfileReply{}
	switch request.PetId {
	case "abc":
		reply.Name = "a dog"
	case "1":
		reply.Name = "a cat"
	default:
		reply.Name = "a snake"
	}
	return reply, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Println("no port")
		return
	}

	port := ":" + os.Args[1]

	cli, err := clientv3.NewFromURL("http://localhost:2379")
	if err != nil {
		log.Fatalf("failed to connect client: %v", err)
	}

	em, err := endpoints.NewManager(cli, service)
	if err != nil {
		log.Fatalf("failed to create manager: %v", err)
	}

	err = em.AddEndpoint(cli.Ctx(), service+"/"+port, endpoints.Endpoint{Addr: port})
	if err != nil {
		log.Fatalf("add endpoint failed: %v", err)
	}

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProfileServer(s, &profileServer{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
