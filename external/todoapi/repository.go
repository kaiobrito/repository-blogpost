package todoapi

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/kaiobrito/repository-blogpost/data"
	"github.com/kaiobrito/repository-blogpost/data/repository"
)

type TodoAPIRepository struct {
	repository.IRepository[data.Todo]

	Token string
}

const (
	BASE_URL = "https://api-nodejs-todolist.herokuapp.com/"
)

func CreateTODOAPIRepository(username string, password string) repository.IRepository[data.Todo] {
	token, err := login(username, password)

	if err != nil {
		panic(err)
	}

	return TodoAPIRepository{
		Token: token,
	}
}

func CreateTODOAPIRepositoryWithToken(token string) repository.IRepository[data.Todo] {
	return TodoAPIRepository{
		Token: token,
	}
}

func login(username string, password string) (string, error) {
	payload := apiLogin{
		Email:    username,
		Password: password,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	var token apiLoginResponse
	err = request("user/login", http.MethodPost, bytes.NewBuffer(data), nil, &token)
	if err != nil {
		return "", err
	}

	return token.Token, err
}

func (r TodoAPIRepository) getHeaders() map[string]string {
	return map[string]string{
		"Authorization": "Bearer " + r.Token,
	}
}

func (r TodoAPIRepository) GetAll(context.Context) ([]*data.Todo, error) {
	var todos apiResponse[[]apiTodo]
	err := request("task", http.MethodGet, nil, r.getHeaders(), &todos)
	if err != nil {
		return nil, err
	}

	var results []*data.Todo
	for _, todo := range todos.Data {
		results = append(results, &data.Todo{
			ID:   todo.ID,
			Name: todo.Name,
			Done: todo.Done,
		})
	}

	return results, err
}

func (r TodoAPIRepository) GetById(_ context.Context, id string) (*data.Todo, error) {
	var res apiResponse[apiTodo]
	err := request("task/"+id, http.MethodGet, nil, r.getHeaders(), &res)

	if err != nil {
		return nil, err
	}

	return &data.Todo{
		ID:   res.Data.ID,
		Name: res.Data.Name,
		Done: res.Data.Done,
	}, err
}

func (r TodoAPIRepository) Create(_ context.Context, todo data.Todo) error {
	payload := apiTodo{
		ID:   "",
		Name: todo.Name,
		Done: todo.Done,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	var res apiTodo
	err = request("task/", http.MethodPost, bytes.NewBuffer(body), r.getHeaders(), &res)
	todo.ID = res.ID
	todo.Name = res.Name
	todo.Done = res.Done

	return err
}

func (r TodoAPIRepository) Save(_ context.Context, todo data.Todo) error {
	payload := apiTodo{
		Name: todo.Name,
		Done: todo.Done,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	err = request[apiTodo]("task/"+todo.ID, http.MethodPut, bytes.NewBuffer(body), r.getHeaders(), nil)
	return err
}
