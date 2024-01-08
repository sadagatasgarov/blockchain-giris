package main

import (
	"context"
	"fmt"
	"log"

	"time"

	"gitlab.com/sadagatasgarov/bchain/node"
	"gitlab.com/sadagatasgarov/bchain/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	node := node.NewNode()
	fmt.Println("node running on port: ", ":3000")
	go func() {
		for {
			time.Sleep(1 * time.Second)
			makeTransaction()
		}
	}()

	log.Fatal(node.Start(":3000"))
}

func makeTransaction() {
	client, err := grpc.Dial(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	c := proto.NewNodeClient(client)

	version := &proto.Version{
		Version: "blocker-0.1",
		Height:  1,
	}

	_, err = c.Handshake(context.TODO(), version)
	if err != nil {
		log.Fatal(err)
	}

}
