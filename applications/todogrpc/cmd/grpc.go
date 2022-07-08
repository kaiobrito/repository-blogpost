package main

import (
	"context"

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
		ID: todo.Id,
		Name: todo.Name,
		Done: todo.Done,
	}
}

func toProtoTodo(todo data.Todo) *proto.Todo {
	return &proto.Todo{
		Id: todo.ID,
		Name: todo.Name,
		Done: todo.Done,
	}
}

func (s server) GetAll(_ *emptypb.Empty, server proto.TodoService_GetAllServer) error {
	todos, err := s.Repo.GetAll()
	if err != nil {
		return err
	}
	for _, todo := range todos {
		server.Send(toProtoTodo(*todo))
	}

	return err
}

func (s server) GetById(c context.Context, f *proto.TodoFilters) (*proto.Todo, error) {
	todo, err := s.Repo.GetById(f.Id)
	if err != nil {
		return nil, err
	}

	return toProtoTodo(*todo), nil
}

func (s server) Create(_ context.Context, todo *proto.Todo) (*proto.Todo, error) {
	err := s.Repo.Create(*toDataTodo(todo))
	return todo, err
}

func (s server) Update(_ context.Context, todo *proto.Todo) (*proto.Todo, error) {
	err := s.Repo.Save(*toDataTodo(todo))
	return todo, err
}