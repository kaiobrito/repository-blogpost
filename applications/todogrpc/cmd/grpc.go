package main

import (
	"context"
	"log"

	"github.com/kaiobrito/repository-blogpost/applications/todogrpc/proto"
	"github.com/kaiobrito/repository-blogpost/data"
	"github.com/kaiobrito/repository-blogpost/data/repository"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	proto.UnimplementedTodoServiceServer

	Repo repository.IRepository[data.Todo]
}

func toDataTodo(todo *proto.Todo) *data.Todo {
	return &data.Todo{
		ID:   todo.Id,
		Name: todo.Name,
		Done: todo.Done,
	}
}

func toProtoTodo(todo data.Todo) *proto.Todo {
	return &proto.Todo{
		Id:   todo.ID,
		Name: todo.Name,
		Done: todo.Done,
	}
}

func (s server) GetAll(_ *emptypb.Empty, server proto.TodoService_GetAllServer) error {
	log.Println("GetAll")
	todos, err := s.Repo.GetAll(server.Context())
	if err != nil {
		return err
	}
	for _, todo := range todos {
		if err := server.Send(toProtoTodo(*todo)); err != nil {
			return err
		}
	}
	log.Println("GetAll: Done")

	return err
}

func (s server) GetById(c context.Context, f *proto.TodoFilters) (*proto.Todo, error) {
	log.Println("GetById: " + f.Id)
	todo, err := s.Repo.GetById(c, f.Id)
	if err != nil {
		return nil, err
	}

	return toProtoTodo(*todo), nil
}

func (s server) Create(ctx context.Context, todo *proto.Todo) (*proto.Todo, error) {
	log.Println("Create")
	err := s.Repo.Create(ctx, *toDataTodo(todo))
	return todo, err
}

func (s server) Update(ctx context.Context, todo *proto.Todo) (*proto.Todo, error) {
	log.Println("Update: " + todo.Id)
	err := s.Repo.Save(ctx, *toDataTodo(todo))
	return todo, err
}
