package main

import (
	"context"
	"log"
	"net"
	"products/src/config"
	"products/src/pb/products"
	"products/src/repository"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	products.ProductServiceServer
	productRepo *repository.ProductRepository
}

func (s *server) Create(ctx context.Context, product *products.Product) (*products.Product, error) {
	newProduct, err := s.productRepo.Create(*product)
	if err != nil {
		return nil, err
	}
	return &newProduct, nil
}

func (s *server) FindAll(ctx context.Context, empty *emptypb.Empty) (*products.ProductList, error) {
	productList, err := s.productRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return &productList, nil
}

func (s *server) FindById(ctx context.Context, id *products.ProductId) (*products.Product, error) {
	product, err := s.productRepo.FindByID(id.Id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, status.Errorf(codes.NotFound, "product with id %d not found", id.Id)
	}

	return product, nil
}

func (s *server) Update(ctx context.Context, product *products.Product) (*products.Product, error) {
	updatedProduct, err := s.productRepo.Update(*product)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func (s *server) Delete(ctx context.Context, id *products.ProductId) (*emptypb.Empty, error) {
	err := s.productRepo.Delete(id.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func main() {
	log.Println("Starting gRPC server")

	dbConfig := config.NewDatabaseConfig()
	db, err := dbConfig.Connect()
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}

	productRepo := repository.NewProductRepository(db)

	srv := server{productRepo: productRepo}

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
