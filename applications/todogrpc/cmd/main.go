package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kaiobrito/repository-blogpost/applications/todogrpc/proto"
	"github.com/kaiobrito/repository-blogpost/data"
	"github.com/kaiobrito/repository-blogpost/data/repository"
	"google.golang.org/grpc"
)

var (
	PORT    = 50051

)


func main(){
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	initialData := []data.Todo{}
	proto.RegisterTodoServiceServer(s, &server{
		Repo: repository.CreateMemoryRepository(initialData),
	})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}