package main

import (
	"context"
	"log"
	"net"
	"products/src/pb/products"
	"products/src/repository"

	"google.golang.org/grpc"
)

type server struct {
	products.ProductServiceServer
	productRepo *repository.ProductRepository
}

func (s *server) Create(ctx context.Context, product *products.Product) (*products.Product, error) {
	newProduct, err := s.productRepo.Create(*product)
	if err != nil {
		return product, err
	}
	return &newProduct, nil
}

func (s *server) FindAll(ctx context.Context, product *products.ProductList) (*products.ProductList, error) {
	productList, err := s.productRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return &productList, nil
}

func main() {
	log.Println("Starting gRPC server")
	srv := server{productRepo: &repository.ProductRepository{}}
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
