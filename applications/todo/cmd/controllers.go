package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaiobrito/repository-blogpost/data"
)

type Response[T any] struct {
	Data []T `json:"data"`
}

func (app *App) GetTodos(ctx *gin.Context) {
	var resp Response[*data.Todo]
	resp.Data = []*data.Todo{}
	for _, todo := range app.Todos {
		resp.Data = append(resp.Data, todo)
	}

	ctx.JSON(http.StatusOK, resp)
}

func (app *App) GetTodoById(ctx *gin.Context) {
	id := ctx.Param("id")

	existingTodo := app.Todos[id]

	if existingTodo == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": "Not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, existingTodo)
}

func (app *App) EditTodos(ctx *gin.Context) {
	id := ctx.Param("id")

	existingTodo := app.Todos[id]

	if existingTodo == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": "Not found",
		})
		return
	}

	var todo data.Todo
	err := ctx.BindJSON(&todo)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	todo.ID = id
	app.Todos[id] = &todo
	ctx.JSON(http.StatusOK, todo)
}

func (app *App) CreateTodos(ctx *gin.Context) {
	var todo data.Todo
	err := ctx.BindJSON(&todo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	app.Todos[todo.ID] = &todo
	ctx.JSON(http.StatusOK, todo)
}
