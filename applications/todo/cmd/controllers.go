package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaiobrito/repository-blogpost/data"
)

func (app *App) GetTodos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, app.Todos)
}

func (app *App) GetTodoById(ctx *gin.Context) {
	id := ctx.Param("id")

	var existingTodo data.Todo

	for _, todo := range app.Todos {
		if id == todo.ID {
			existingTodo = todo
			break
		}
	}

	if existingTodo.ID == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": "Not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, existingTodo)
}

func (app *App) EditTodos(ctx *gin.Context) {
	id := ctx.Param("id")

	var existingTodo data.Todo

	for _, todo := range app.Todos {
		if id == todo.ID {
			existingTodo = todo
			break
		}
	}

	if existingTodo.ID == "" {
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
	app.Todos = append(app.Todos, todo)
	ctx.JSON(http.StatusOK, app.Todos)
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
	app.Todos = append(app.Todos, todo)
	ctx.JSON(http.StatusOK, app.Todos)
}
