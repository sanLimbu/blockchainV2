package main

import (
	"context"
	"log"
	"time"

	"github.com/sanLimbu/blockchain/node"
	"github.com/sanLimbu/blockchain/proto"
	"google.golang.org/grpc"
)

func main() {

	makeNode(":3000", []string{})
	time.Sleep(time.Second)
	makeNode(":4000", []string{":3000"})
	time.Sleep(3 * time.Second)
	makeNode(":5000", []string{":4000"})
	select {}
}

func makeNode(listenAddr string, bootstrapNodes []string) *node.Node {
	n := node.NewNode()
	go n.Start(listenAddr)
	if len(bootstrapNodes) > 0 {
		if err := n.BootstrapNetwork(bootstrapNodes); err != nil {
			log.Fatal(err)
		}
	}
	return n
}

func makeTransaction() {
	client, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := proto.NewNodeClient(client)

	version := &proto.Version{
		Version:    "blockchain-1.0",
		Height:     100,
		ListenAddr: ":4000",
	}

	_, err = c.Handshake(context.TODO(), version)
	if err != nil {
		log.Fatal(err)
	}
}
