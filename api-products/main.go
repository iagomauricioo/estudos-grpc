package main

import (
	"context"
	"log"
	"net"
	"products/src/pb/products"

	"google.golang.org/grpc"
)

type server struct {
	products.ProductServiceServer
}

func (s *server) Create(ctx context.Context, product *products.Product) (*products.Product, error) {
	return &products.Product{}, nil
}

func (s *server) FindAll(ctx context.Context, product *products.ProductList) (*products.ProductList, error) {
	return &products.ProductList{}, nil
}

func (s *server) FindById(ctx context.Context, product *products.Product) (*products.Product, error) {
	return &products.Product{
		Id:       int32(1),
		Name:     "Product 1",
		Price:    100,
		Quantity: 10,
	}, nil
}

func main() {
	log.Println("Starting gRPC server")
	srv := server{}
	listener, err := net.Listen("tcp", ":9090")

	if err != nil {
		log.Fatalln("error on create listener. error:", err)
	}

	s := grpc.NewServer()

	products.RegisterProductServiceServer(s, &srv)

	log.Println("server is running on port 9090")

	if err := s.Serve(listener); err != nil {
		log.Fatalln("error on serve. error: ", err)
	}
}
