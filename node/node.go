package node

import (
	"context"
	"fmt"
	"net"

	"gitlab.com/sadagatasgarov/bchain/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type Node struct {
	version string
	// peers map[net.Addr]*grpc.ClientConn
	proto.UnimplementedNodeServer
}

func NewNode() *Node {
	return &Node{
		version: "blocker-0.1",
	}
}

func (n *Node) Start(listenAddr string) error {
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	proto.RegisterNodeServer(grpcServer, n)
	return grpcServer.Serve(ln)
}

func (n *Node) Handshake(ctx context.Context, v *proto.Version) (*proto.Version, error) {
	outVersion := &proto.Version{
		Version: n.version,
		Height:  100,
	}

	p, _ := peer.FromContext(ctx)
	fmt.Printf("received verion from %s %+v\n", v, p.Addr)
	return outVersion, nil
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)

	fmt.Println("received tx from: ", peer)
	return &proto.Ack{}, nil
}
