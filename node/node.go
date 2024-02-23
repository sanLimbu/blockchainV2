package node

import (
	"context"
	"fmt"

	"github.com/sanLimbu/blockchain/proto"
	"google.golang.org/grpc/peer"
)

type Node struct {
	proto.UnimplementedNodeServer
	//peer map[net.Addr]
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.None, error) {

	peer, _ := peer.FromContext(ctx)
	fmt.Println("received tx from :", peer)

	return &proto.None{}, nil

}
