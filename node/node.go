package node

import (
	"context"
	"fmt"
	"net"
	"sync"

	"gitlab.com/sadagatasgarov/bchain/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/peer"
)

type Node struct {
	version  string
	peerlock sync.RWMutex
	peers    map[proto.NodeClient]*proto.Version
	proto.UnimplementedNodeServer
}

func NewNode() *Node {
	return &Node{
		peers:   make(map[proto.NodeClient]*proto.Version),
		version: "blocker-0.1",
	}
}

func (n *Node) addPeer(c proto.NodeClient, v *proto.Version) {
	n.peerlock.Lock()
	defer n.peerlock.Unlock()
	fmt.Printf("new peer connected: (%s) - height:%d\n", v.ListenAddr, v.Height)
	n.peers[c] = v
}

func (n *Node) deletePeer(c proto.NodeClient) {
	n.peerlock.Lock()
	defer n.peerlock.Unlock()
	delete(n.peers, c)
}

func (n *Node) Start(listenAddr string) error {
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	ln, err := net.Listen("tcp", listenAddr)
	fmt.Println("node running on port: ", listenAddr)

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
	//p, _ := peer.FromContext(ctx)

	c, err := makeNodeClient(v.ListenAddr)
	if err != nil {
		return nil, err
	}
	n.addPeer(c, v)

	//fmt.Printf("received verion from %s %+v\n", v, p.Addr)
	return outVersion, nil
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)

	fmt.Println("received tx from: ", peer)
	return &proto.Ack{}, nil
}

func makeNodeClient(listenAddr string) (proto.NodeClient, error) {
	c, err := grpc.Dial(listenAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return proto.NewNodeClient(c), err
}
