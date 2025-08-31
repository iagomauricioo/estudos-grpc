package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"server-streaming/src/pb/department"
	"strconv"
	"strings"

	"google.golang.org/grpc"
)

type server struct {
	department.DepartmentServiceServer
}

func (s *server) ListPerson(req *department.ListPersonRequest, srv department.DepartmentService_ListPersonServer) error {
	log.Println("ListPerson")
	file, err := os.Open("./data.csv")
	if err != nil {
		return fmt.Errorf("open file err:%v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		id, _ := strconv.Atoi(data[0])
		name := data[1]
		email := data[2]
		income, _ := strconv.Atoi(data[3])
		departmentId, _ := strconv.Atoi(data[4])

		if int32(departmentId) == req.GetDepartmentId() {
			if err := srv.Send(&department.ListPersonResponse{
				Id:           int32(id),
				Name:         name,
				Email:        email,
				Income:       int32(income),
				DepartmentId: int32(departmentId),
			}); err != nil {
				return fmt.Errorf("error on send: %v", err)
			}
		}
	}
	return nil
}

func main() {
	fmt.Println("starting gRPC server")
	lis, err := net.Listen("tcp", ":9090")

	if err != nil {
		log.Fatalln("failed to listen:", err)
	}

	s := grpc.NewServer()

	department.RegisterDepartmentServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
