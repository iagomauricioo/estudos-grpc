package main

import (
	"client-streaming/src/pb/calc"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	calc.CalcServiceServer
}

func (s *server) Calc(stream calc.CalcService_CalcServer) error {
	for {
		input, err := stream.Recv()

		if err != nil {
			return err
		}

		fmt.Printf("input: %+v\n", input)
	}
}

func main() {
	log.Println("starting gRPC server")

	lis, err := net.Listen("tcp", ":9090")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calc.RegisterCalcServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("gRPC server started port 9090")
}
