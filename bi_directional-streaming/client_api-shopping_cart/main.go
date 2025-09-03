package main

import (
	"context"
	"io"
	"log"
	"src/src/pb/shoppingcart"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AddItem(conn *grpc.ClientConn) {
	cartClient := shoppingcart.NewShoppingCartServiceClient(conn)
	stream, err := cartClient.AddItem(context.Background())

	if err != nil {
		log.Fatalf("%v.AddItem(_) = _, %v", cartClient, err)
	}

	waitch := make(chan struct{})

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(waitch)
			}

			if err != nil {
				log.Fatalf("%v.AddItem(_) = _, %v", stream, err)
			}

			log.Println(res)
		}
	}()

	items := []shoppingcart.AddProduct{
		{ProductId: 1, Quantity: 1, PriceUnit: 5.0},
		{ProductId: 2, Quantity: 10, PriceUnit: 10.0},
		{ProductId: 3, Quantity: 1, PriceUnit: 100.0},
	}

	for _, v := range items {
		if err := stream.Send(&v); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, v, err)
		}
	}

	stream.CloseSend()

	<-waitch
}

func main() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to create gRPC client: %v", err)
	}

	AddItem(conn)

	defer conn.Close()
}
