package main

import (
	"client/src/pb/products"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to create gRPC client: %v", err)
	}

	findAllProducts(conn)

	defer conn.Close()
}

func findAllProducts(conn *grpc.ClientConn) {
	productClient := products.NewProductServiceClient(conn)
	productList, err := productClient.FindAll(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalln("failed to find all products: ", err)
	}

	fmt.Printf("products: %+v\n", productList)
}
