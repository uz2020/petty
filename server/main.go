package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "t/pet"

	"go.etcd.io/etcd/client/v3"

	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

const (
	port    = ":50051"
	service = "foo/bar/my-service"
)

type profileServer struct {
	pb.UnimplementedProfileServer
}

func (*profileServer) GetProfile(ctx context.Context, request *pb.ProfileRequest) (*pb.ProfileReply, error) {
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
	log.Println("Serving gRPC on 0.0.0.0" + port)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
