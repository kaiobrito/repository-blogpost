package repository

import (
	"context"
	"io"

	"github.com/kaiobrito/repository-blogpost/applications/todogrpc/proto"
	"github.com/kaiobrito/repository-blogpost/data"
	"github.com/kaiobrito/repository-blogpost/data/repository"
	"google.golang.org/grpc"
)

type TodoGRPCService struct {
	client proto.TodoServiceClient
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

func CreateTodoGRPCService(uri string) repository.IRepository[data.Todo] {
	conn, err := grpc.Dial(uri, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewTodoServiceClient(conn)

	return &TodoGRPCService{
		client: client,
	}
}

func (s *TodoGRPCService) GetAll() ([]*data.Todo, error) {
	service, err := s.client.GetAll(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	
	results := []*data.Todo{}
	for {
		todo, err := service.Recv()
		if err == io.EOF {
			return results, nil
		}
		if err != nil {
			return nil, err
		  }
		results = append(results, toDataTodo(todo))
	}
}

func (s TodoGRPCService) GetById(id string) (*data.Todo, error) {
	todo, err := s.client.GetById(context.TODO(), &proto.TodoFilters{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	
	return toDataTodo(todo), nil
}

func (s TodoGRPCService) Create(todo data.Todo) error {
	_, err := s.client.Create(context.TODO(), toProtoTodo(todo))
	return err
}

func (s TodoGRPCService) Save(todo data.Todo) error {
	_, err := s.client.Update(context.TODO(), toProtoTodo(todo))
	return err
}
