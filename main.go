package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"gitlab.com/sadagatasgarov/bchain/node"
	"gitlab.com/sadagatasgarov/bchain/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	node := node.NewNode()
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	proto.RegisterNodeServer(grpcServer, node)
	fmt.Println("node running on port: ", ":3000")

	go func() {
		for {
			time.Sleep(2 * time.Second)
			makeTransaction()
		}
	}()

	grpcServer.Serve(ln)
}

func makeTransaction() {
	client, err := grpc.Dial(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	c := proto.NewNodeClient(client)

	tx := &proto.Transaction{
		Version: 1,

	}

	_, err = c.HandleTransaction(context.TODO(), tx)
	if err != nil {
		log.Fatal(err)
	}

}
