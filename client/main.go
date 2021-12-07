package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/uz2020/petty/pb/pet"
	"google.golang.org/grpc"

	"go.etcd.io/etcd/client/v3"
	resolver "go.etcd.io/etcd/client/v3/naming/resolver"
)

const (
	defaultPetId = "abc"
)

func main() {
	cli, err := clientv3.NewFromURL("http://localhost:2379")
	if err != nil {
		log.Fatalf("new client failed: %v", err)
	}

	etcdResolver, err := resolver.NewBuilder(cli)
	if err != nil {
		log.Fatalf("new builder failed: %v", err)
	}

	conn, err := grpc.Dial("etcd:///foo/bar/my-service", grpc.WithResolvers(etcdResolver), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial failed: %v", err)
	}

	defer conn.Close()
	c := pb.NewProfileClient(conn)

	// Contact the server and print out its response.
	petId := defaultPetId
	if len(os.Args) > 1 {
		petId = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := c.GetProfile(ctx, &pb.ProfileRequest{PetId: petId})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetName())

	r, err = c.GetProfile(ctx, &pb.ProfileRequest{PetId: ""})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetName())

	r, err = c.GetProfile(ctx, &pb.ProfileRequest{PetId: "1"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetName())

}
