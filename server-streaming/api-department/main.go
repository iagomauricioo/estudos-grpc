package main

import (
	"fmt"
	"server-streaming/src/pb/department"
)

type server struct {
	department.DepartmentServiceServer
}

func (s *server) ListPerson(req *department.ListPersonRequest, srv department.DepartmentService_ListPersonServer) error {
	return nil
}

func main() {
	fmt.Println("starting gRPC server")
}
