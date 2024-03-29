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
	makeNode(":3000", []string{})
	time.Sleep(1 * time.Second)
	makeNode(":4000", []string{":3000"})
	time.Sleep(4 * time.Second)
	makeNode(":5000", []string{":4000"})

	makeNode(":6000", []string{":3000"})
	// go func() {
	// 	for {
	// 		time.Sleep(4 * time.Second)
	// 		//makeTransaction()
	// 	}
	// }()

	select {}
}

func makeNode(listenAddr string, bootstrapNodes []string) *node.Node {
	n := node.NewNode()
	go n.Start(listenAddr, bootstrapNodes)
	return n
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
