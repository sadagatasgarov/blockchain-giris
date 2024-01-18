package node

import (
	"context"
	"fmt"
	"net"
	"sync"

	"gitlab.com/sadagatasgarov/bchain/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/peer"
)

type Node struct {
	version    string
	listenAddr string
	logger     *zap.SugaredLogger

	peerlock sync.RWMutex
	peers    map[proto.NodeClient]*proto.Version

	proto.UnimplementedNodeServer
}

func NewNode() *Node {
	loggerConfig := zap.NewDevelopmentConfig()
	loggerConfig.Development = true
	loggerConfig.EncoderConfig.TimeKey = ""

	logger, _ := loggerConfig.Build()
	return &Node{
		peers:   make(map[proto.NodeClient]*proto.Version),
		version: "blocker-0.1",
		logger:  logger.Sugar(),
	}
}

func (n *Node) Start(listenAddr string, bootstrapNodes []string) error {
	n.listenAddr = listenAddr

	var (
		opts       = []grpc.ServerOption{}
		grpcServer = grpc.NewServer(opts...)
	)

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	proto.RegisterNodeServer(grpcServer, n)

	n.logger.Infow("node running on", "port", n.listenAddr)

	// bootstrap the network with alist of alreafdy known nodes
	// in the network
	if len(bootstrapNodes) > 0 {
		go n.bootstrapNetwork(bootstrapNodes)
	}

	return grpcServer.Serve(ln)
}

func (n *Node) Handshake(ctx context.Context, v *proto.Version) (*proto.Version, error) {
	c, err := makeNodeClient(v.ListenAddr)
	if err != nil {
		return nil, err
	}

	// do logic here before we accept the incomming connection as valid
	n.addPeer(c, v)

	//fmt.Printf("received verion from %s %+v\n", v, p.Addr)
	return n.getVersion(), nil
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)

	fmt.Println("received tx from: ", peer)
	return &proto.Ack{}, nil
}

func (n *Node) addPeer(c proto.NodeClient, v *proto.Version) {
	n.peerlock.Lock()
	defer n.peerlock.Unlock()

	// Handle the logic whwere we decided we except or  drop
	// the incomming node connections

	n.peers[c] = v

	// connect to all peers in the received peerlist of peers
	//// for _, addr := range v.PeerList {
	//// 	if addr != n.listenAddr {
	//// 		c, v, err := n.dialRemoteNode(addr)
	//// 		if err != nil {
	////
	//// 		}
	// 		//fmt.Printf("[%s] need to connect with %s \n", n.listenAddr, addr)
	//// 	}

	if len(v.PeerList) > 0 {
		go n.bootstrapNetwork(v.PeerList)
	}

	n.logger.Debugw(
		"new peer succesfully connected",
		"we", n.listenAddr,
		"remote", v.ListenAddr,
		"height", v.Height,
	)

}

func (n *Node) deletePeer(c proto.NodeClient) {
	n.peerlock.Lock()
	defer n.peerlock.Unlock()
	delete(n.peers, c)
}

func (n *Node) bootstrapNetwork(addrs []string) error {
	for _, addr := range addrs {
		// if addr == n.listenAddr {
		// 	continue
		// }
		if !n.canConnectWith(addr){
			continue
		}
		n.logger.Debugw(
			"dialing remote nodeboot",
			"we", n.listenAddr,
			"remote", addr,
		)
		c, v, err := n.dialRemoteNode(addr)
		if err != nil {
			return err
		}
		n.addPeer(c, v)
	}
	return nil
}

func (n *Node) dialRemoteNode(addr string) (proto.NodeClient, *proto.Version, error) {
	c, err := makeNodeClient(addr)
	if err != nil {
		return nil, nil, err
	}

	v, err := c.Handshake(context.Background(), n.getVersion())
	if err != nil {
		return nil, nil, err
	}

	return c, v, nil
}

func (n *Node) getVersion() *proto.Version {
	return &proto.Version{
		Version:    "blocker-0.1",
		Height:     0,
		ListenAddr: n.listenAddr,
		PeerList:   n.getPeerList(),
	}
}

func (n *Node) canConnectWith(addr string) bool {
	if n.listenAddr == addr {
		return false
	}

	connectedPeers := n.getPeerList()
	for _, connectAddr := range connectedPeers {
		if addr == connectAddr {
			return false
		}
	}

	return true
}

func (n *Node) getPeerList() []string {
	n.peerlock.RLock()
	defer n.peerlock.RUnlock()
	peers := []string{}
	for _, version := range n.peers {
		peers = append(peers, version.ListenAddr)
	}
	return peers
}

func makeNodeClient(listenAddr string) (proto.NodeClient, error) {
	c, err := grpc.Dial(listenAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return proto.NewNodeClient(c), err
}
