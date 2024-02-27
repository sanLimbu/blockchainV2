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
	client, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := proto.NewNodeClient(client)

	version := &proto.Version{
		Version: "blockchain-1.0",
		Height:  100,
	}

	_, err = c.Handshake(context.TODO(), version)
	if err != nil {
		log.Fatal(err)
	}
}
