package main

import (
	"api-shopping_cart/src/pb/shoppingcart"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	shoppingcart.ShoppingCartServiceServer
}

func (s *server) AddItem(srv shoppingcart.ShoppingCartService_AddItemServer) error {
	var quantityItems int32 = 0
	var priceTotal float64 = 0.0
	for {
		newItem, err := srv.Recv()

		//end of file
		if err == io.EOF {
			return srv.Send(&shoppingcart.ShoppingCartTotal{
				QuantityItems: quantityItems,
				PriceTotal:    priceTotal,
			})
		}

		if err != nil {
			return err
		}

		quantityItems += newItem.GetQuantity()
		priceTotal += float64(newItem.GetPriceUnit() * float64(newItem.GetQuantity()))
		if err := srv.Send(&shoppingcart.ShoppingCartTotal{
			QuantityItems: quantityItems,
			PriceTotal:    priceTotal,
		}); err != nil {
			return fmt.Errorf("error on send: %v", err)
		}
	}
}

func main() {
	log.Println("running gRPC server")

	lis, err := net.Listen("tcp", ":9090")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	shoppingcart.RegisterShoppingCartServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
