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
func login(username string, password string) (string, error) {
	payload := apiLogin{
		Email:    username,
		Password: password,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	res, err := sendRequest("user/login", http.MethodPost, bytes.NewBuffer(data), nil)
	if err != nil {
		return "", err
	}

	var token apiLoginResponse
	err = json.Unmarshal(res, &token)

	return token.Token, err
}

func (r TodoAPIRepository) getHeaders() map[string]string {
	return map[string]string{
		"Authorization": "Bearer " + r.Token,
	}
}

func (r TodoAPIRepository) GetAll(context.Context) ([]*data.Todo, error) {
	res, err := sendRequest("task", http.MethodGet, nil, r.getHeaders())
	if err != nil {
		return nil, err
	}

	var todos apiResponse[apiTodo]
	err = json.Unmarshal(res, &todos)

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

func (r TodoAPIRepository) GetById(context.Context, string) (*data.Todo, error) {
	return nil, nil
}

func (r TodoAPIRepository) Create(context.Context, data.Todo) error {
	return nil
}

func (r TodoAPIRepository) Save(context.Context, data.Todo) error {
	return nil
}
