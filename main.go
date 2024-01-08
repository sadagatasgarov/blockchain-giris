package main

import (
	"context"
	"log"

	"time"

	"gitlab.com/sadagatasgarov/bchain/node"
	"gitlab.com/sadagatasgarov/bchain/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	node := node.NewNode()
	
	go func() {
		for {
			time.Sleep(2 * time.Second)
			makeTransaction()
		}
	}()
	log.Fatal(node.Start(":3000"))
	
}

func makeTransaction() {
	listenAddr := ":3000"
	client, err := grpc.Dial(listenAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	c := proto.NewNodeClient(client)

	version := &proto.Version{
		Version:    "blocker-0.1",
		Height:     1,
		ListenAddr: ":4000",
	}

	_, err = c.Handshake(context.TODO(), version)
	if err != nil {
		log.Fatal(err)
	}

}
