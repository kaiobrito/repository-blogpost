package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaiobrito/repository-blogpost/data"
)

type Response[T any] struct {
	Data []T `json:"data"`
}

func (app *App) GetTodos(ctx *gin.Context) {
	log.Println("GetTodos")
	todos, err := app.Repo.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	resp := Response[*data.Todo]{
		Data: todos,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (app *App) GetTodoById(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Println("GetTodoById: " + id)

	existingTodo, err := app.Repo.GetById(ctx.Request.Context(), id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, existingTodo)
}

func (app *App) EditTodos(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Println("EditTodos: " + id)

	_, err := app.Repo.GetById(ctx.Request.Context(), id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	var todo data.Todo
	err = ctx.BindJSON(&todo)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	todo.ID = id
	app.Repo.Save(ctx.Request.Context(), todo)
	ctx.JSON(http.StatusOK, todo)
}

func (app *App) CreateTodos(ctx *gin.Context) {
	log.Println("CreateTodos")
	var todo data.Todo
	err := ctx.BindJSON(&todo)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	app.Repo.Create(ctx.Request.Context(), todo)
	ctx.JSON(http.StatusOK, todo)
}
