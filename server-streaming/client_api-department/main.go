package main

import (
	"client_api-department/src/pb/department"
	"context"
	"io"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to create gRPC client: %v", err)
	}

	if len(os.Args) < 2 {
		log.Fatalf("usage: go run main.go <departmentId>")
	}

	depID, err := strconv.Atoi(os.Args[1])

	ListPersons(conn, int32(depID))

	defer conn.Close()
}

func ListPersons(conn *grpc.ClientConn, departmentID int32) {
	personClient := department.NewDepartmentServiceClient(conn)
	stream, err := personClient.ListPerson(context.Background(), &department.ListPersonRequest{DepartmentId: departmentID})

	if err != nil {
		log.Fatalln("failed to list person: ", err)
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListPerson(_) = _, %v", stream, err)
		}
		log.Println(response)
	}
}
