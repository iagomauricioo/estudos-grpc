package main

import (
	"client_api-calc/src/pb/calc"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to create gRPC client: %v", err)
	}

	Calc(conn)

	defer conn.Close()
}

func Calc(conn *grpc.ClientConn) {
	calcClient := calc.NewCalcServiceClient(conn)
	stream, err := calcClient.Calc(context.Background())

	if err != nil {
		log.Fatalln("failed to calc: ", err)
	}

	nums := []int32{1, 2, 3}

	for _, v := range nums {
		if err := stream.Send(&calc.Input{Value: v}); err != nil {
			log.Fatalln("failed to send: ", err)
		}
	}
	response, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalln("failed to receive response: ", err)
	}
	fmt.Println("response: \n", response)
}
